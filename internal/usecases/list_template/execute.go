package list_template

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/vctaragao/code-to-image/internal/entity"
	"github.com/vctaragao/code-to-image/internal/helper"
)

func Execute() (OutputDto, error) {
	currentFolder, err := os.Getwd()
	if err != nil {
		return OutputDto{}, fmt.Errorf("getting executable folder: %w", err)
	}

	dir, err := os.Open(currentFolder + "/template")
	if err != nil {
		return OutputDto{}, fmt.Errorf("openning templates folder: %w", err)
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return OutputDto{}, fmt.Errorf("readding templates folder: %w", err)
	}

	var out OutputDto
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() && fileInfo.Name() != "layout" {
			template, err := formatTemplate(fileInfo.Name())
			if err != nil {
				helper.LogError("invalid template: ", err)
				continue
			}
			out.Templates = append(out.Templates, *template)
		}
	}

	return out, nil
}

func formatTemplate(templateId string) (*entity.Template, error) {
	data, err := helper.GetFileFromTemplate(templateId, "config.json")
	if err != nil {
		return &entity.Template{}, fmt.Errorf("getting the configuration of the template: %w", err)
	}

	var conf *entity.Config
	json.Unmarshal(data, &conf)

	return entity.NewTemplate(templateId, conf), nil
}
