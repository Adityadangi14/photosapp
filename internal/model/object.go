package models

type ObjectModel struct {
	ObjectName string
	Type       string
	CreatedAt  string
	AbsPath    string
	OtherInfo  map[string]any
}
