package internal

type Variable struct {
	Flag        string `json:"flag"`
	Type        string `json:"type"`
	Default     any    `json:"default"`
	Description string `json:"Description"`
}

type TemplateMetadata struct {
	Name        string              `json:"name"`
	DisplayName string              `json:"displayName"`
	Description string              `json:"Description"`
	Language    string              `json:"language"`
	Framework   string              `json:"framework"`
	Variables   map[string]Variable `json:"variables"`
	Flags       map[string]struct {
		Description string `json:"description"`
	} `json:"flags"`
}
