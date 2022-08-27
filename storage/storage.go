package storage

import (
	"github.com/404th/go_grpc_project/models"
	"github.com/404th/go_grpc_project/storage/inmemory"
)

// STORE
var Store = inmemory.BookRepo

func init() {
	// storing extra contents
	Store.CreateBookSg(
		models.Book{
			Name: "The Graf Monte Cristo",
			ISBN: "2132-d343-3243-4d24-3242",
			AuthorName: models.Author{
				Firstname:  "Samuel",
				Secondname: "Umtiti",
			},
			Year: 2019,
		},
	)

	Store.CreateBookSg(
		models.Book{
			Name: "Master and Margaret",
			ISBN: "4322-435d-3hv3-f9fd-4545",
			AuthorName: models.Author{
				Firstname:  "Clement",
				Secondname: "Lenglet",
			},
			Year: 1989,
		},
	)
}
