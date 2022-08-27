package models

type Book struct {
	Name       string `json:"name"`
	ISBN       string `json:"isbn"`
	AuthorName Author `json:"author_name"`
	Year       int64  `json:"year"`
}
