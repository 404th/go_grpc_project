package models

type Book struct {
	ID       string `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	ISBN     string `json:"isbn" db:"isbn"`
	AuthorID string `json:"author_id" db:"author_id"`
	Year     int64  `json:"year" db:"year"`
}
