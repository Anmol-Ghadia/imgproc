package test

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"testing"
)

func TestAdd(t *testing.T) {

	got := 1 + 1
	want := 2

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

// Test 1 for format command
// All pixel black, square image
func TestFormat1(t *testing.T) {
	outputFile := "output/format/whiteBox_100x100.png"
	outputImage, err := openPngImage(outputFile)
	if err != nil {
		t.Errorf("Error: %s\n", err.Error())
		return
	}
	expectedFile := "expected/format/whiteBox_100x100.png"
	ExpectedImage, err := openPngImage(expectedFile)
	if err != nil {
		t.Errorf("Error: %s\n", err.Error())
		return
	}

	err = compareImage(outputImage, ExpectedImage)
	if err != nil {
		t.Errorf("Error: %s\n", err.Error())
		return
	}

}

func TestFormat2(t *testing.T) {
	outputFile := "output/format/flower_720x720.png"
	outputImage, err := openPngImage(outputFile)
	if err != nil {
		t.Errorf("Error: %s\n", err.Error())
		return
	}
	expectedFile := "expected/format/flower_720x720.png"
	ExpectedImage, err := openPngImage(expectedFile)
	if err != nil {
		t.Errorf("Error: %s\n", err.Error())
		return
	}

	err = compareImage(outputImage, ExpectedImage)
	if err != nil {
		t.Errorf("Error: %s\n", err.Error())
		return
	}

}

func TestFormat3(t *testing.T) {
	outputFile := "output/format/blackBox_100x100.jpg"
	outputImage, err := openJpegImage(outputFile)
	if err != nil {
		t.Errorf("Error: %s\n", err.Error())
		return
	}
	expectedFile := "expected/format/blackBox_100x100.jpg"
	ExpectedImage, err := openJpegImage(expectedFile)
	if err != nil {
		t.Errorf("Error: %s\n", err.Error())
		return
	}

	err = compareImage(outputImage, ExpectedImage)
	if err != nil {
		t.Errorf("Error: %s\n", err.Error())
		return
	}

}

func TestResize1(t *testing.T) {
	outputFile := "output/resize/text_218x80.png"
	outputImage, err := openPngImage(outputFile)
	if err != nil {
		t.Errorf("Error: %s\n", err.Error())
		return
	}
	expectedFile := "expected/resize/text_218x80.png"
	ExpectedImage, err := openPngImage(expectedFile)
	if err != nil {
		t.Errorf("Error: %s\n", err.Error())
		return
	}

	err = compareImage(outputImage, ExpectedImage)
	if err != nil {
		t.Errorf("Error: %s\n", err.Error())
		return
	}

}

func TestCrop1(t *testing.T) {
	outputFile := "output/crop/flower_450x500.jpg"
	outputImage, err := openJpegImage(outputFile)
	if err != nil {
		t.Errorf("Error: %s\n", err.Error())
		return
	}

	expectedFile := "expected/crop/flower_450x500.jpg"
	ExpectedImage, err := openJpegImage(expectedFile)
	if err != nil {
		t.Errorf("Error: %s\n", err.Error())
		return
	}

	err = compareImage(outputImage, ExpectedImage)
	if err != nil {
		t.Errorf("Error: %s\n", err.Error())
		return
	}

}

func TestCrop2(t *testing.T) {
	outputFile := "output/crop/bird_600x256.png"
	outputImage, err := openPngImage(outputFile)
	if err != nil {
		t.Errorf("Error: %s\n", err.Error())
		return
	}

	expectedFile := "expected/crop/bird_600x256.png"
	ExpectedImage, err := openPngImage(expectedFile)
	if err != nil {
		t.Errorf("Error: %s\n", err.Error())
		return
	}

	err = compareImage(outputImage, ExpectedImage)
	if err != nil {
		t.Errorf("Error: %s\n", err.Error())
		return
	}

}

func TestMulti1(t *testing.T) {
	outputFile := "output/multi/flower_450x500.png"
	outputImage, err := openPngImage(outputFile)
	if err != nil {
		t.Errorf("Error: %s\n", err.Error())
		return
	}

	expectedFile := "expected/multi/flower_450x500.png"
	ExpectedImage, err := openPngImage(expectedFile)
	if err != nil {
		t.Errorf("Error: %s\n", err.Error())
		return
	}

	err = compareImage(outputImage, ExpectedImage)
	if err != nil {
		t.Errorf("Error: %s\n", err.Error())
		return
	}

}

func compareImage(out image.Image, expected image.Image) error {
	if out.Bounds().Dx() != expected.Bounds().Dx() {
		return errors.New("Image width does not match")
	}

	if out.Bounds().Dy() != expected.Bounds().Dy() {
		return errors.New("Image height does not match")
	}

	for x := 0; x < out.Bounds().Dx(); x++ {
		for y := 0; y < out.Bounds().Dy(); y++ {
			if comparePixelRGB(out.At(x, y), expected.At(x, y)) != nil {
				return errors.New("Different pixel values")
			}
		}
	}

	return nil
}

func comparePixelRGB(output color.Color, expected color.Color) error {

	r1, g1, b1, a1 := output.RGBA()
	r2, g2, b2, a2 := expected.RGBA()

	// original * alpha = newValue

	var factor uint32 = (2 ^ 16) - 1
	r1 = (r1 / a1) / factor
	r2 = (r2 / a2) / factor
	g1 = (g1 / a1) / factor
	g2 = (g2 / a2) / factor
	b1 = (b1 / a1) / factor
	b2 = (b2 / a2) / factor
	if r1 != r2 || g1 != g2 || b1 != b2 {
		// fmt.Printf("!!! Alpha: %d==%d\n", a1, a2)
		fmt.Printf("!!! R: %d==%d G: %d==%d B: %d==%d\n", r1, r2, g1, g2, b1, b2)
		return errors.New("") // Placeholder
	}

	return nil
}

func openJpegImage(filePath string) (image.Image, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return image.Black, errors.New("Error in reading")
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		return image.Black, errors.New("Error in decode")
	}
	return img, nil
}

func openPngImage(filePath string) (image.Image, error) {
	file, err := os.Open(filePath)
	if err != nil {
		// ERROR !!!
		return image.Black, errors.New("Error in reading")
	}

	img, err := png.Decode(file)
	if err != nil {
		// ERROR !!!
		return image.Black, errors.New("Error in decode")
	}
	return img, nil
}
