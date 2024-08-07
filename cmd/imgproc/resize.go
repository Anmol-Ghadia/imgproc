package imgproc

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Anmol-Ghadia/imgproc/pkg/imgproc"
	"github.com/spf13/cobra"
)

var resizeCmd = &cobra.Command{
	Use: "resize",
	// Aliases: []string{"rev"},
	Short: "resizes the image and saves to a new location, input output width height",
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {

		// read width
		width, err := strconv.Atoi(args[2])
		if err != nil {
			// Handle the error if the conversion fails
			fmt.Printf("Width is not a number\n")
			return
		}

		// read height
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

		// Process image
		if imgproc.ResizeNearestNeighbor(inFile, outFile, width, height) != nil {
			fmt.Printf("Error resizing\n")
			return
		}

		fmt.Printf("Image saved as %s\n", args[1])
	},
}

func init() {
	rootCmd.AddCommand(resizeCmd)
}
