/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vctaragao/code-to-image/internal"
	"github.com/vctaragao/code-to-image/internal/helper"
)

var (
	templateId string
	outputId   string
	content    string
)

func newCreateCmd(codetoImage *internal.CodeToImage) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Crate an image from a html template",
		Long:  `A longer description`,
		Run:   runCreateCmd(codetoImage),
	}
}

func init() {
	createCmd := newCreateCmd(NewCodeToImage())
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&templateId, "template_id", "t", "default", "id of the template to use")
	createCmd.Flags().StringVarP(&outputId, "output_id", "o", "default.html", "id of the html/png output file")
	createCmd.Flags().StringVarP(&content, "content", "c", "", "name of the file that has the content to populate the template")

	createCmd.MarkFlagRequired("content")
}

func runCreateCmd(codeToImage *internal.CodeToImage) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if err := codeToImage.Create(templateId, outputId, content); err != nil {
			helper.LogError("error on creating image from html template", err)
		}
	}
}
