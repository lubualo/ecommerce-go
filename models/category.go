package models

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}
