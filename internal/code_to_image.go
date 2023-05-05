package internal

import (
	"github.com/vctaragao/code-to-image/internal/usecases/create"
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

	if err := create.Execute(dto); err != nil {
		return err
	}
	return nil
}
