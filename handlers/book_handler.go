package handlers

import (
	"net/http"
	"strconv"

	"github.com/404th/go_grpc_project/models"
	"github.com/404th/go_grpc_project/storage"
	"github.com/gin-gonic/gin"
)

func GetBookList(c *gin.Context) {
	resp := storage.GetBookListSg()
	c.JSON(http.StatusOK, gin.H{
		"message": "GetBookList",
		"body":    resp,
	})
}

func GetBookByID(c *gin.Context) {
	str_id := c.Param("id")
	id, err := strconv.Atoi(str_id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"body": "invalid param",
		})
		return
	}

	bk, err := storage.GetBookByIDSg(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error while getting book by id",
			"body":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "GetBookByID",
		"body":    bk,
	})
}

func CreateBook(c *gin.Context) {
	var data models.Book

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"body": "Bad data",
		})
		return
	}

	if err := storage.CreateBookSg(data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error while creating book",
			"body":    err,
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Book created successfully",
		"body":    nil,
	})
}

func UpdateBook(c *gin.Context) {
	var upd_book models.Book

	if err := c.ShouldBindJSON(&upd_book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "not updated",
			"body":    err.Error(),
		})
		return
	}

	str_id := c.Param("id")
	id, err := strconv.Atoi(str_id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "could not update book",
			"body":    "invalid param while updating",
		})
		return
	}

	if err := storage.UpdateBookSg(upd_book, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "book not updated",
			"body":    err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "book updated successfully",
		"body":    nil,
	})
}

func DeleteBook(c *gin.Context) {
	str_id := c.Param("id")
	id, err := strconv.Atoi(str_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid param while deleting",
			"body":    err.Error(),
		})
		return
	}

	if err := storage.DeleteBookSg(id); err != nil {
		c.JSON(http.StatusNotModified, gin.H{
			"message": "not deleted",
			"body":    err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "book deleted",
		"body":    nil,
	})
}
