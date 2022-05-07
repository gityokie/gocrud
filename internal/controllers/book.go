package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gityokie/gocrud/internal/api_helpers"
	"github.com/gityokie/gocrud/internal/models"
)

func ListBook(c *gin.Context) {
	var book []models.Book
	err := models.GetAllBook(&book)
	if err != nil {
		api_helpers.RespondJSON(c, 404, book)
	} else {
		api_helpers.RespondJSON(c, 200, book)
	}
}

func AddNewBook(c *gin.Context) {
	var book models.Book
	c.BindJSON(&book)
	err := models.AddNewBook(&book)
	if err != nil {
		api_helpers.RespondJSON(c, 404, book)
	} else {
		api_helpers.RespondJSON(c, 200, book)
	}
}

func GetOneBook(c *gin.Context) {
	id := c.Params.ByName("id")
	var book models.Book
	err := models.GetOneBook(&book, id)
	if err != nil {
		api_helpers.RespondJSON(c, 404, book)
	} else {
		api_helpers.RespondJSON(c, 200, book)
	}
}

func PutOneBook(c *gin.Context) {
	var book models.Book
	id := c.Params.ByName("id")
	err := models.GetOneBook(&book, id)
	if err != nil {
		api_helpers.RespondJSON(c, 404, book)
	}
	c.BindJSON(&book)
	err = models.PutOneBook(&book, id)
	if err != nil {
		api_helpers.RespondJSON(c, 404, book)
	} else {
		api_helpers.RespondJSON(c, 200, book)
	}
}

func DeleteBook(c *gin.Context) {
	var book models.Book
	id := c.Params.ByName("id")
	err := models.DeleteBook(&book, id)
	if err != nil {
		api_helpers.RespondJSON(c, 404, book)
	} else {
		api_helpers.RespondJSON(c, 200, book)
	}
}
