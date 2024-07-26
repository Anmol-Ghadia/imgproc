package imgproc

import (
	"fmt"
	"os"

	"github.com/Anmol-Ghadia/imgproc/pkg/imgproc"
	"github.com/spf13/cobra"
)

var formatCmd = &cobra.Command{
	Use: "format",
	// Aliases: []string{"rev"},
	Short: "crops the image and saves to a new location, input output width height",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		inFilePath := args[0]
		inFile, err := os.Open(inFilePath)
		if err != nil {
			fmt.Printf("Error opening input file: %v\n", err)
			return
		}
		defer inFile.Close()

		outFile, err := os.Create(args[1])
		if err != nil {
			fmt.Println("Error writing output file:", err)
			return
		}
		defer outFile.Close()

		if imgproc.Fromat(inFile, outFile) != nil {
			fmt.Printf("Error resizing\n")
			return
		}

		fmt.Printf("Image saved as %s\n", args[1])
	},
}

func init() {
	rootCmd.AddCommand(formatCmd)
}
