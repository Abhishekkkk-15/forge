package internal

type Variable struct {
	Type        string `json:"type"`
	Flag        string `json:"flag,omitempty"`
	Default     any    `json:"default,omitempty"`
	Source      string `json:"source,omitempty"`
	Description string `json:"description,omitempty"`
}

type Flag struct {
	Type        string `json:"type"`
	Default     any    `json:"default"`
	Description string `json:"description"`
}

type TemplateMetadata struct {
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Variables   map[string]Variable `json:"variables"`
	Flags       map[string]Flag     `json:"flags"`
}
