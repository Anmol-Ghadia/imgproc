package imgproc

import (
	"errors"
	"image"

	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"

	"os"
)

// returns the average pixel value as RGB
func CropImg(inpFile *os.File, outFile *os.File, newWidth int, newHeight int) error {
	inputImg, _, err := image.Decode(inpFile)
	if err != nil {
		return errors.New("decode error")
	}
	// Resize the image
	outputImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	for x := 0; x < newWidth; x++ {
		for y := 0; y < newHeight; y++ {
			outputImg.Set(x, y, inputImg.At(x, y))
		}
	}

	jpeg.Encode(outFile, outputImg, nil)
	return nil
}

// Inspects the format and size of image
func Inspect(f *os.File) (formatString string, x int, y int, e error) {

	config, format, err := image.DecodeConfig(f)

	if err != nil {
		return "", 0, 0, errors.New("error decoding image")
	}

	return format, config.Width, config.Height, nil
}
