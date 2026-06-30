package models

type MetaFile struct {
	Prev bool          `json:"prev"`
	Data []ObjectModel `json:"data"`
}
