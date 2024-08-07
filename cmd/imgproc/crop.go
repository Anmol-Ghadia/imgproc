package imgproc

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Anmol-Ghadia/imgproc/pkg/imgproc"
	"github.com/spf13/cobra"
)

var cropCmd = &cobra.Command{
	Use: "crop",
	// Aliases: []string{"rev"},
	Short: "crops the image and saves to a new location, input output width height",
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {

		// read crop width
		width, err := strconv.Atoi(args[2])
		if err != nil {
			// Handle the error if the conversion fails
			fmt.Printf("Width is not a number\n")
			return
		}

		// read crop height
		height, err := strconv.Atoi(args[3])
		if err != nil {
			// Handle the error if the conversion fails
			fmt.Printf("Height is not a number\n")
			return
		}

		// open input file
		inFilePath := args[0]
		inFile, err := os.Open(inFilePath)
		if err != nil {
			fmt.Printf("Error opening input file: %v\n", err)
			return
		}
		defer inFile.Close()

		// create output file
		outFile, err := os.Create(args[1])
		if err != nil {
			fmt.Println("Error writing output file:", err)
			return
		}
		defer outFile.Close()

		// Check bounds
		_, originalWidth, originalHeight, err := imgproc.Inspect(inFile)
		if err != nil {
			fmt.Printf("Error inspecting input file: %v\n", err)
			return
		}
		if width <= 0 || originalWidth <= width {
			fmt.Printf("Width out of bounds, expected: 0<width<input_image_width, given: 0<%v<%v\n", width, originalWidth)
			return
		}
		if height <= 0 || originalHeight <= height {
			fmt.Printf("Height out of bounds, expected: 0<height<input_image_height, given: 0<%v<%v\n", height, originalHeight)
			return
		}

		// process image
		if imgproc.CropImg(inFile, outFile, width, height) != nil {
			fmt.Printf("Error cropping\n")
			return
		}

		fmt.Printf("Image saved as %s\n", args[1])
	},
}

func init() {
	rootCmd.AddCommand(cropCmd)
}
