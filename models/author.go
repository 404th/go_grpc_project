package models

type Author struct {
	ID         string `json:"id" db:"id"`
	Firstname  string `json:"firstname" db:"firstname"`
	Secondname string `json:"secondname" db:"secondname"`
}
