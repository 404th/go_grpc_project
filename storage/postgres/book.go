package postgres

import (
	"errors"
	"log"

	"github.com/404th/go_grpc_project/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type bookRepoImpl struct {
	db *sqlx.DB
}

var BookRepo = bookRepoImpl{}

var dbSchema = `
	CREATE TABLE IF NOT EXISTS author (
		id varchar(150) primary key not null,
		firstname varchar(150) not null,
		secondname varchar(150) not null
	);

	CREATE TABLE IF NOT EXISTS book (
		"id" varchar(150) primary key not null,
		"name" varchar(300) not null,
		"isbn" varchar(100) not null,
		"author_id" varchar(150),
		"year" int not null,
		CONSTRAINT fk_book_author
			FOREIGN KEY(author_id)
				REFERENCES author(id)
	);
`

func init() {
	db, err := sqlx.Connect("postgres", "host=localhost port=2323 user=postgres password=postgres dbname=bookstore sslmode=disable")
	if err != nil {
		log.Fatalf("error while connecting to database: %s", err.Error())
	}

	db.MustExec(dbSchema)

	tx := db.MustBegin()

	tx.MustExec(`INSERT INTO author (
		id, firstname, secondname
	) VALUES (
		$1, $2, $3
	);`, uuid.NewString(), "John", "Doe")

	tx.MustExec(`INSERT INTO author (
		id, firstname, secondname
	) VALUES (
		$1, $2, $3
	);`, uuid.NewString(), "Steve", "Jobs")

	if err := tx.Commit(); err != nil {
		tx.Rollback()
	}
	BookRepo.db = db
}

func (r *bookRepoImpl) GetBookListSg() (resp []models.Book) {
	// for _, v := range r.db {
	// 	resp = append(resp, v)
	// }

	return resp
}

func (r *bookRepoImpl) GetBookByIDSg(id int) (*models.Book, error) {
	var data models.Book
	// var exists bool = false

	// for ind, val := range r.db {
	// 	if ind == id {
	// 		data = val
	// 		exists = true
	// 		break
	// 	}
	// }

	// if !exists {
	// 	return nil, errors.New("book not found")
	// }

	return &data, nil
}

func (r *bookRepoImpl) CreateBookSg(bk models.Book) error {
	new_id := uuid.NewString()

	query := `INSERT INTO book (
		id, name, isbn, author_id, year
	) VALUES (
		$1, $2, $3, $4, $5
	);`

	_, err := r.db.Exec(query, new_id, bk.Name, bk.ISBN, bk.AuthorID, bk.Year)
	if err != nil {
		return err
	}

	return nil
}

func (r *bookRepoImpl) UpdateBookSg(bk models.Book, id int) error {
	// for k := range r.db {
	// 	if k == id {
	// 		r.db[k] = bk
	// 		return nil
	// 	}
	// }

	return errors.New("book not found")
}

func (r *bookRepoImpl) DeleteBookSg(id int) error {
	// for k := range r.db {
	// 	if k == id {
	// 		delete(r.db, k)
	// 		return nil
	// 	}
	// }

	return errors.New("book not found")
}
