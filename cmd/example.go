/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/vctaragao/code-to-image/internal"
	"github.com/vctaragao/code-to-image/internal/helper"
)

func newExampleCmd(codeToImage *internal.CodeToImage) *cobra.Command {
	return &cobra.Command{
		Use:   "example",
		Short: "Show an example of a chosen template",
		Long:  `Loger description`,
		Run:   runExampleCmd(codeToImage),
	}
}
func init() {
	exampleCmd := newExampleCmd(NewCodeToImage())
	rootCmd.AddCommand(exampleCmd)

	exampleCmd.Flags().StringVarP(&templateId, "template", "t", "default", "the template to generate the example for")
}

func runExampleCmd(codeToImage *internal.CodeToImage) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if err := codeToImage.Create(templateId, "example.html", "example.json"); err != nil {
			helper.LogError("creating template example", err)
			return
		}

		fileServer := http.FileServer(http.Dir("draft"))

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fileServer.ServeHTTP(w, r)
		})

		log.Println("See example file on http://localhost:8080/example.html")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}
}
