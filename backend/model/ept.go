package model

type Ept struct {
	Base
	Name       string
	Version    string
	StorageKey string
	FileSize   int64
	Integrity  string
}
