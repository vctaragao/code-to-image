package list_draft

import (
	"fmt"
	"path/filepath"

	"github.com/vctaragao/code-to-image/internal/entity"
	"github.com/vctaragao/code-to-image/internal/helper"
)

func Execute() (OutputDto, error) {
	fileInfos, err := helper.GetDirectoryContent("draft")
	if err != nil {
		return OutputDto{}, fmt.Errorf("getting draft directory content: %w", err)
	}

	var out OutputDto
	for _, fileInfo := range fileInfos {
		fileName := fileInfo.Name()
		if filepath.Ext(fileName) == ".html" {
			out.Drafts = append(out.Drafts, *entity.NewDraft(fileName))
		}

	}

	return out, nil
}
