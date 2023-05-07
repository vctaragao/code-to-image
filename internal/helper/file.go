package helper

import (
	"fmt"
	"os"
)

const TEMPLATES_FOLDER = "template"

func GetFileFromTemplate(templateFolder, file string) ([]byte, error) {
	currentFolder, err := os.Getwd()
	if err != nil {
		fmt.Println("error on getting current folder: ", err)
		return []byte{}, err
	}

	filePath := getTemplatePath(currentFolder, templateFolder, file)

	f, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("error on getting %s/%s: %v", templateFolder, file, err)
		return []byte{}, err
	}

	return f, nil
}

func getTemplatePath(currentFolder, templateFolder, file string) string {
	return fmt.Sprintf("%s\\%s\\%s\\%s", currentFolder, TEMPLATES_FOLDER, templateFolder, file)
}