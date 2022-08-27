package inmemory

import (
	"errors"

	"github.com/404th/go_grpc_project/models"
)

type bookRepoImpl struct {
	db map[int]models.Book
}

var BookRepo = bookRepoImpl{}

func init() {
	BookRepo.db = make(map[int]models.Book)
}

func (r *bookRepoImpl) GetBookListSg() (resp []models.Book) {
	for _, v := range r.db {
		resp = append(resp, v)
	}

	return resp
}

func (r *bookRepoImpl) GetBookByIDSg(id int) (*models.Book, error) {
	var data models.Book
	var exists bool = false

	for ind, val := range r.db {
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

func (r *bookRepoImpl) CreateBookSg(bk models.Book) error {
	r.db[len(r.db)+1] = bk

	return nil
}

func (r *bookRepoImpl) UpdateBookSg(bk models.Book, id int) error {
	for k := range r.db {
		if k == id {
			r.db[k] = bk
			return nil
		}
	}

	return errors.New("book not found")
}

func (r *bookRepoImpl) DeleteBookSg(id int) error {
	for k := range r.db {
		if k == id {
			delete(r.db, k)
			return nil
		}
	}

	return errors.New("book not found")
}
