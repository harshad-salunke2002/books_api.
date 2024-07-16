package models

type Books struct {
	Id     int    `json:"id" orm:"auto;pk"`
	Name   string `json:"name"`
	Pages  int    `json:"pages"`
	Writer string `json:"writer"`
}
