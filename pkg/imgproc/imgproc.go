package imgproc

import (
	"bufio"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"strconv"
	"strings"
)

// returns the average pixel value as RGB
func Fromat(inpFile *os.File, outFile *os.File) error {
	inputImg, _, err := image.Decode(inpFile)
	if err != nil {
		return throwDecodeError()
	}

	width := inputImg.Bounds().Dx()
	height := inputImg.Bounds().Dy()
	outputImg := image.NewNRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			outputImg.Set(x, y, inputImg.At(x, y))
		}
	}

	writeImage(outFile, outputImg)
	return nil
}

// Crops the image starting at 0,0
func CropImg(inpFile *os.File, outFile *os.File, newWidth int, newHeight int) error {
	inputImg, _, err := image.Decode(inpFile)
	if err != nil {
		return throwDecodeError()
	}
	// Resize the image
	outputImg := image.NewNRGBA(image.Rect(0, 0, newWidth, newHeight))

	for x := 0; x < newWidth; x++ {
		for y := 0; y < newHeight; y++ {
			outputImg.Set(x, y, inputImg.At(x, y))
		}
	}

	writeImage(outFile, outputImg)
	return nil
}

// Inspects the format and size of image
func Inspect(f *os.File) (formatString string, x int, y int, e error) {

	config, format, err := image.DecodeConfig(f)
	if err != nil {
		return "", 0, 0, throwDecodeError()
	}

	f.Seek(0, 0)
	return format, config.Width, config.Height, nil
}

// Inspects the format and size of image
func ResizeNearestNeighbor(inpFile *os.File, outFile *os.File, newWidth int, newHeight int) error {
	inputImg, _, err := image.Decode(inpFile)
	if err != nil {
		return throwDecodeError()
	}

	width := inputImg.Bounds().Dx()
	height := inputImg.Bounds().Dy()

	// Calculate scaling factors
	scaleX := float64(width) / float64(newWidth)
	scaleY := float64(height) / float64(newHeight)

	// Create a new blank image with the new dimensions
	outputImg := image.NewNRGBA(image.Rect(0, 0, newWidth, newHeight))

	for x := 0; x < newWidth; x++ {
		for y := 0; y < newHeight; y++ {
			// Calculate the corresponding position in the original image
			origX := int(float64(x) * scaleX)
			origY := int(float64(y) * scaleY)

			// Get the color of the nearest pixel
			color := inputImg.At(origX, origY)
			outputImg.Set(x, y, color)
		}
	}

	writeImage(outFile, outputImg)
	return nil
}

// Inspects the format and size of image
func KernelBlur(inpFile *os.File, outFile *os.File, matFile *os.File) error {
	inputImg, _, err := image.Decode(inpFile)
	if err != nil {
		return throwDecodeError()
	}

	width := inputImg.Bounds().Dx()
	height := inputImg.Bounds().Dy()

	kernel, err := readKerenel(matFile)
	if err != nil {
		return err
	}

	outputImg := image.NewNRGBA(image.Rect(0, 0, width, height))
	for x := (len(kernel) - 1) / 2; x < width-((len(kernel)-1)/2); x++ { // Bounds are temp!!!
		for y := (len(kernel) - 1) / 2; y < height-((len(kernel)-1)/2); y++ { // Bounds are temp!!!
			color := getKernelColor(x, y, kernel, inputImg)
			outputImg.Set(x, y, color)
		}
	}

	writeImage(outFile, outputImg)
	return nil
}

// === Helpers ===

func throwDecodeError() error {
	return errors.New("error in decoding image")
}

// Writes the image to file based on file extension
func writeImage(outFile *os.File, outputImg *image.NRGBA) {
	if getFileExtension(outFile.Name(), 3) == "png" {

		png.Encode(outFile, outputImg)

	} else if getFileExtension(outFile.Name(), 3) == "jpg" ||
		getFileExtension(outFile.Name(), 4) == "jpeg" {

		jpeg.Encode(outFile, outputImg, nil)
	}
}

// returns the file extension excluding the dot (e.g., "png")
func getFileExtension(fileName string, extensionNameLength int) string {
	if len(fileName) < extensionNameLength+1 {
		return ""
	}

	// Check dot exists where expected
	if fileName[len(fileName)-(extensionNameLength+1):len(fileName)-extensionNameLength] != "." {
		return ""
	}

	// Get the file extension from the end of the fileName
	ext := fileName[len(fileName)-extensionNameLength:]
	return ext
}

// Returns the kernel read from the file
//
// raises error if file not read
// raises error if not a square kernel
// raises error if kernel of even length
func readKerenel(inpFile *os.File) ([][]int, error) {
	var kernel [][]int
	scanner := bufio.NewScanner(inpFile)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)

		// Convert values to integers
		var row []int
		for _, value := range values {
			num, err := strconv.Atoi(value)
			if err != nil {
				return kernel, err
			}
			row = append(row, num)
		}
		if len(row)%2 != 1 {
			return kernel, errors.New("kernel of even length not allowed")
		}
		kernel = append(kernel, row)
	}

	if len(kernel)%2 != 1 {
		return kernel, errors.New("kernel of even length not allowed")
	}

	for i := 0; i < len(kernel); i++ {
		for j := 0; j < len(kernel[i]); j++ {
			fmt.Printf(" %d", kernel[i][j])
		}
		fmt.Printf("\n")
	}

	return kernel, scanner.Err()
}

// Returns the pixel value at given position for a given kernel and image
func getKernelColor(x int, y int, kernel [][]int, img image.Image) color.Color {
	var totalRed uint64 = 0
	var totalGreen uint64 = 0
	var totalBlue uint64 = 0
	var totalAlpha uint64 = 0
	kernelHalfSize := (len(kernel) - 1) / 2
	for i := -kernelHalfSize; i <= kernelHalfSize; i++ {
		for j := -kernelHalfSize; j <= kernelHalfSize; j++ {
			red, green, blue, alpha := img.At(x+i, y+j).RGBA()
			totalRed += uint64(red * uint32(kernel[i+kernelHalfSize][j+kernelHalfSize]))
			totalGreen += uint64(green * uint32(kernel[i+kernelHalfSize][j+kernelHalfSize]))
			totalBlue += uint64(blue * uint32(kernel[i+kernelHalfSize][j+kernelHalfSize]))
			totalAlpha += uint64(alpha * uint32(kernel[i+kernelHalfSize][j+kernelHalfSize]))
		}
	}
	normalizingFactor := uint64(getKernelSum(kernel))
	red := uint16(totalRed / normalizingFactor)
	green := uint16(totalGreen / normalizingFactor)
	blue := uint16(totalBlue / normalizingFactor)
	alpha := uint16(totalAlpha / normalizingFactor)

	return color.RGBA64{red, green, blue, alpha}
}

// returns the sum of all values in the kernel
func getKernelSum(kernel [][]int) int {
	var total int = 0
	for i := 0; i < len(kernel); i++ {
		for j := 0; j < len(kernel[i]); j++ {
			total += kernel[i][j]
		}
	}
	return total
}
