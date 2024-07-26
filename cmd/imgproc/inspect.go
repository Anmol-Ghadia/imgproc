package imgproc

import (
	"fmt"
	"os"

	"github.com/Anmol-Ghadia/imgproc/pkg/imgproc"
	"github.com/spf13/cobra"
)

var inspectCmd = &cobra.Command{
	Use: "inspect",
	// Aliases: []string{"insp"},
	Short: "Inspects an image, for format and dimensions",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		filePath := args[0]
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			return
		}
		defer file.Close()

		format, x, y, err := imgproc.Inspect(file)
		if err != nil {
			fmt.Printf("Error decoding image\n")
			return
		}

		fmt.Printf("Read %s image of size %dx%d\n", format, x, y)
	},
}

func init() {
	rootCmd.AddCommand(inspectCmd)
}
