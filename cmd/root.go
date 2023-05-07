package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/vctaragao/code-to-image/internal"
)

var rootCmd = &cobra.Command{
	Use:   "code-to-image",
	Short: "code-to-image is a library to populate and convert html templates into image files",
	Long:  `Long description`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func NewCodeToImage() *internal.CodeToImage {
	return internal.NewCodeToImage()
}
