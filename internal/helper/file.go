package helper

import (
	"fmt"
	"log"
	"os"
)

const (
	TEMPLATES_FOLDER = "template"
	LAYOUTS_FOLDER   = "layout"
)

func GetFileFromTemplate(templateFolder, file string) ([]byte, error) {
	filePath, err := GetTemplatePath(templateFolder, file)
	if err != nil {
		return []byte{}, fmt.Errorf("getting template path: %w", err)
	}

	f, err := os.ReadFile(filePath)
	if err != nil {
		return []byte{}, fmt.Errorf("reading %s/%s: %w", templateFolder, file, err)
	}

	return f, nil
}

func GetTemplatePath(templateFolder, file string) (string, error) {
	return fmt.Sprintf("%s\\%s\\%s\\%s", getCurrentFolder(), TEMPLATES_FOLDER, templateFolder, file), nil
}

func GetLayoutPath(layout string) string {
	return fmt.Sprintf("%s\\%s\\%s\\%s", getCurrentFolder(), TEMPLATES_FOLDER, LAYOUTS_FOLDER, layout)
}

func GetDirectoryContent(folder string) ([]os.FileInfo, error) {
	dir, err := os.Open(getCurrentFolder() + "/" + folder)
	if err != nil {
		return []os.FileInfo{}, fmt.Errorf("openning %s folder: %w", folder, err)
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return []os.FileInfo{}, fmt.Errorf("readding %s folder: %w", folder, err)
	}

	return fileInfos, nil
}

func getCurrentFolder() string {
	currentFolder, err := os.Getwd()
	if err != nil {
		LogError("error on getting current folder", err)
		log.Fatal(err)
	}

	return currentFolder
}
