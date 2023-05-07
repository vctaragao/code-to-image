package create

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"text/template"

	"github.com/vctaragao/code-to-image/internal/entity"
	"github.com/vctaragao/code-to-image/internal/helper"
)

func Execute(dto *InputDto) error {
	cont, err := getContentFromFile(dto.ContentFile)
	if err != nil {
		return fmt.Errorf("getting content file: %w", err)
	}

	if err = buildFromTemplate(dto, cont); err != nil {
		return fmt.Errorf("building from template: %w", err)
	}

	return nil
}

func getContentFromFile(file string) (*entity.Content, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return &entity.Content{}, fmt.Errorf("reading content file: %w", err)
	}

	var cont *entity.Content
	if err := json.Unmarshal(data, &cont); err != nil {
		return &entity.Content{}, fmt.Errorf("decoding content file: %w", err)
	}

	return cont, nil
}

func buildFromTemplate(dto *InputDto, cont *entity.Content) error {
	body, err := fillTemplateBody(dto.TemplateId, cont)
	if err != nil {
		return fmt.Errorf("filling template body: %w", err)
	}

	draft, err := fillLayout(dto.TemplateId, body)
	if err != nil {
		return fmt.Errorf("filling layout: %w", err)
	}

	if err = os.WriteFile("draft/"+dto.OutputId, draft, 0644); err != nil {
		return fmt.Errorf("creating/writing to output file: %w", err)
	}

	return nil
}

func fillTemplateBody(templateFolder string, cont *entity.Content) ([]byte, error) {
	templateBodyPath, err := helper.GetTemplatePath(templateFolder, "body.html")
	if err != nil {
		return []byte{}, fmt.Errorf("getting template body path: %w", err)
	}

	buffer := bytes.NewBuffer([]byte{})
	if err := template.Must(template.New("body.html").ParseFiles(templateBodyPath)).Execute(buffer, cont); err != nil {
		return []byte{}, fmt.Errorf("filling template body: %w", err)
	}

	return buffer.Bytes(), nil
}

func fillLayout(templateFolder string, body []byte) ([]byte, error) {
	conf, err := getTemplateConfig(templateFolder)
	if err != nil {
		return []byte{}, err
	}

	header, err := helper.GetFileFromTemplate(templateFolder, "header.html")
	if err != nil {
		return []byte{}, fmt.Errorf("getting header.html: %w", err)
	}

	style, err := helper.GetFileFromTemplate(templateFolder, "style.css")
	if err != nil {
		return []byte{}, fmt.Errorf("getting style.css: %w", err)
	}

	layoutPath := helper.GetLayoutPath(conf.LayoutId)
	layout := entity.NewLayout(string(body), string(header), string(style))

	buffer := bytes.NewBuffer([]byte{})
	if err = template.Must(template.New(conf.LayoutId).ParseFiles(layoutPath)).Execute(buffer, layout); err != nil {
		return []byte{}, fmt.Errorf("filling layout: %w", err)
	}

	return buffer.Bytes(), nil
}

func getTemplateConfig(templateFolder string) (*entity.Config, error) {
	data, err := helper.GetFileFromTemplate(templateFolder, "config.json")
	if err != nil {
		return &entity.Config{}, fmt.Errorf("getting template config file: %w", err)
	}

	var conf *entity.Config
	if err := json.Unmarshal(data, &conf); err != nil {
		return &entity.Config{}, fmt.Errorf("decoding template config file: %w", err)
	}

	return conf, nil
}
