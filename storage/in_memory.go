package storage

import (
	"errors"

	"github.com/404th/go_grpc_project/models"
)

// STORE
var BookStore map[int]models.Book

func init() {
	// init store
	BookStore = make(map[int]models.Book)

	// storing extra contents
	BookStore[1] = models.Book{
		Name: "The Graf Monte Cristo",
		ISBN: "2132-d343-3243-4d24-3242",
		AuthorName: models.Author{
			Firstname:  "Samuel",
			Secondname: "Umtiti",
		},
		Year: 2019,
	}

	BookStore[2] = models.Book{
		Name: "Master and Margaret",
		ISBN: "4322-435d-3hv3-f9fd-4545",
		AuthorName: models.Author{
			Firstname:  "Clement",
			Secondname: "Lenglet",
		},
		Year: 1989,
	}
}

func GetBookListSg() (resp []models.Book) {
	for _, v := range BookStore {
		resp = append(resp, v)
	}

	return resp
}

func GetBookByIDSg(id int) (*models.Book, error) {
	var data models.Book
	var exists bool = false

	for ind, val := range BookStore {
		if ind == id {
			data = val
			exists = true
			break
		}
	}

	if !exists {
		return nil, errors.New("book not found")
	}

	return &data, nil
}

func CreateBookSg(bk models.Book) error {
	BookStore[len(BookStore)+1] = bk

	return nil
}

func UpdateBookSg(bk models.Book, id int) error {
	for k := range BookStore {
		if k == id {
			BookStore[k] = bk
			return nil
		}
	}

	return errors.New("book not found")
}

func DeleteBookSg(id int) error {
	for k := range BookStore {
		if k == id {
			delete(BookStore, k)
			return nil
		}
	}

	return errors.New("book not found")
}
