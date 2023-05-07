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

func newDraftCmd(codeToImage *internal.CodeToImage) *cobra.Command {
	return &cobra.Command{
		Use:   "draft",
		Short: "A list the post drafts",
		Long:  `A longer description`,
		Run:   runDraftCmd(codeToImage),
	}
}

func init() {
	draftCmd := newDraftCmd(NewCodeToImage())
	listCmd.AddCommand(draftCmd)

	draftCmd.Flags().BoolP("serve", "s", false, "serve the post drafts for for better vizualization")
}

func runDraftCmd(codeToImage *internal.CodeToImage) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		drafts, err := codeToImage.ListDraft()
		if err != nil {
			helper.LogError("listing drafts", err)
			return
		}

		fmt.Printf("%-15s\n", "ID")

		for _, draft := range drafts.Drafts {
			fmt.Printf("%-15s\n", draft.Id)
		}
	}
}
