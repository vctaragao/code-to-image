package list_template

import (
	"encoding/json"
	"fmt"

	"github.com/vctaragao/code-to-image/internal/entity"
	"github.com/vctaragao/code-to-image/internal/helper"
)

func Execute() (OutputDto, error) {
	folders, err := helper.GetDirectoryContent("template")
	if err != nil {
		return OutputDto{}, fmt.Errorf("getting directory content: %w", err)
	}

	var out OutputDto
	for _, folder := range folders {
		if folder.Name() == "layout" {
			continue
		}

		template, err := formatTemplate(folder.Name())
		if err != nil {
			helper.LogError("invalid template: ", err)
			continue
		}
		out.Templates = append(out.Templates, *template)
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
