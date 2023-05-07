/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vctaragao/code-to-image/internal"
	"github.com/vctaragao/code-to-image/internal/helper"
)

func newTemplateCmd(codeToImage *internal.CodeToImage) *cobra.Command {
	return &cobra.Command{
		Use:   "template",
		Short: "A list of the avaliable templates",
		Long:  `A longer description`,
		Run:   runTemplateCmd(codeToImage),
	}
}

func init() {
	templateCmd := newTemplateCmd(NewCodeToImage())
	listCmd.AddCommand(templateCmd)
}

func runTemplateCmd(codeToImage *internal.CodeToImage) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		templates, err := codeToImage.ListTemplate()
		if err != nil {
			helper.LogError("listing templates", err)
			return
		}

		fmt.Printf("%-15s %-15s\n", "ID", "LAYOUT")

		for _, template := range templates.Templates {
			fmt.Printf("%-15s %-15s\n", template.Id, template.Config.LayoutId)
		}
	}
}
