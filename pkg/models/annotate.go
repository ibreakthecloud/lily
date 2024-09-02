package models

type Annotation struct {
	EntityName  string `json:"entity_name"`
	EntityType  string `json:"entity_type"`
	Type        string `json:"type"`
	Description string `json:"description"`
}
