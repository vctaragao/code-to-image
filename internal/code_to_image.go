package internal

import (
	"github.com/vctaragao/code-to-image/internal/usecases/create"
	"github.com/vctaragao/code-to-image/internal/usecases/list_template"
)

type CodeToImage struct {
}

func NewCodeToImage() *CodeToImage {
	return &CodeToImage{}
}

func (c *CodeToImage) Create(templateId, outputId, textfile string) error {
	dto := &create.InputDto{
		TemplateId: templateId,
		OutputId:   outputId,
		TextFile:   textfile,
	}

	return create.Execute(dto)
}

func (c *CodeToImage) ListTemplate() (list_template.OutputDto, error) {
	return list_template.Execute()
}
