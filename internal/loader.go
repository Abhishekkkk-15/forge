package internal

import (
	"encoding/json"
	"fmt"

	"github.com/pterm/pterm"
)

func LoadMetadata(tempalteName string) (*TemplateMetadata, error) {
	path := fmt.Sprintf("templates/%s/template.json", tempalteName)
	data, err := TemplateFS.ReadFile(path)
	if err != nil {
		pterm.Error.Printf("template %s metadata not found", tempalteName)
		return nil, fmt.Errorf("template %s metadata not found", tempalteName)
	}
	var meta TemplateMetadata
	if err := json.Unmarshal(data, &meta); err != nil {
		return nil, err
	}
	return &meta, nil

}
