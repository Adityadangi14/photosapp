package models

type ObjectModel struct {
	ObjectName string         `json:"objectName"`
	Type       string         `json:"type"`
	CreatedAt  string         `json:"createdAt"`
	AbsPath    string         `json:"absPath"`
	BlurHash   string         `json:"blurHash"`
	OtherInfo  map[string]any `json:"otherInfo"`
}
