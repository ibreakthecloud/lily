package models

type Annotation struct {
	EntityName  string `json:"entity_name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}
