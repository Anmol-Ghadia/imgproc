package imgproc

import (
	"fmt"
	"os"

	"github.com/Anmol-Ghadia/imgproc/pkg/imgproc"
	"github.com/spf13/cobra"
)

var kernelBlurCmd = &cobra.Command{
	Use: "kernelBlur",
	// Aliases: []string{"kblur"},
	Short: "Blurs the image based on the input kernel",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {

		inFile, err := os.Open(args[0])
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

		matFile, err := os.Open(args[2])
		if err != nil {
			fmt.Println("Error reading matrix file:", err)
			return
		}
		defer matFile.Close()

		err = imgproc.KernelBlur(inFile, outFile, matFile)
		if err != nil {
			// !!! TEMP
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("Image saved as %s\n", args[1])
	},
}

func init() {
	rootCmd.AddCommand(kernelBlurCmd)
}
