package imgproc

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:     "imgproc",
	Version: version,
	Short:   "imgproc(image-process) - a simple CLI to process images",
	Long: `imgproc(image-process) is a fast and lightweight image processor
   
Utilities include formatting, compressing, transformations, effects`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
