package imgproc

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"

	"os"
)

// returns the average pixel value as RGB
func CropImg(inpFile *os.File, outFile *os.File, newWidth int, newHeight int) error {
	inputImg, _, err := image.Decode(inpFile)
	if err != nil {
		return throwDecodeError()
	}
	// Resize the image
	outputImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

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
	outputImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

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

// === Helpers ===

func throwDecodeError() error {
	return errors.New("error in decoding image")
}

// Writes the image to file based on file extension
func writeImage(outFile *os.File, outputImg *image.RGBA) {
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
