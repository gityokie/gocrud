package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gityokie/gocrud/internal/api_helpers"
	"github.com/gityokie/gocrud/internal/models"
	"github.com/gityokie/gocrud/internal/store"
)

type AddNewBook struct {
	Store store.Store
}

func (h AddNewBook) Handle(c *gin.Context) {
	var book models.Book

	c.BindJSON(&book)

	err := h.Store.AddNewBook(&book)
	if err != nil {
		api_helpers.RespondJSON(c, 404, book)
	} else {
		api_helpers.RespondJSON(c, 200, book)
	}
}
