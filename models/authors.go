package models

type Author struct {
	Id    int      `json:"id" orm:"auto"`
	Name  string   `json:"name"`
	Books []*Books `json:"books" orm:"reverse(many)"`
}
