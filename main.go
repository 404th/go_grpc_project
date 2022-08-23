package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Author struct {
	Firstname  string
	Secondname string
}

type Book struct {
	Name       string
	ISBN       string
	AuthorName Author
	Year       int64
}

// STORE
var BookStore map[int]Book

func main() {
	// 1. Declaring GIN project
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())

	// init store
	BookStore = make(map[int]Book)

	// storing extra contents
	BookStore[1] = Book{
		Name: "The Graf Monte Cristo",
		ISBN: "2132-d343-3243-4d24-3242",
		AuthorName: Author{
			Firstname:  "Samuel",
			Secondname: "Umtiti",
		},
		Year: 2019,
	}

	BookStore[2] = Book{
		Name: "Master and Margaret",
		ISBN: "4322-435d-3hv3-f9fd-4545",
		AuthorName: Author{
			Firstname:  "Clement",
			Secondname: "Lenglet",
		},
		Year: 1989,
	}

	// ROUTES
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"body": "pong",
		})
	})

	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"body": BookStore,
		})
	})

	engine.GET("/:id", func(c *gin.Context) {
		str_id := c.Param("id")
		id, err := strconv.Atoi(str_id)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"body": "invalid param",
			})
			return
		}

		var data Book
		var exists bool = false

		for ind, val := range BookStore {
			if ind == id {
				data = val
				exists = true
				break
			}
		}

		if exists {
			c.JSON(http.StatusOK, gin.H{
				"body": data,
			})
			return
		}

		c.JSON(http.StatusNotFound, gin.H{
			"body": "Book not found",
		})
	})

	engine.POST("/", func(c *gin.Context) {
		var data Book

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"body": "Bad data",
			})
			return
		}

		BookStore[len(BookStore)+1] = data
		c.JSON(http.StatusCreated, gin.H{
			"body": fmt.Sprintf("book created under name %d", len(BookStore)),
		})
	})

	// listening port
	engine.Run(":4040")
}
