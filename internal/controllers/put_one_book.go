package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gityokie/gocrud/internal/api_helpers"
	"github.com/gityokie/gocrud/internal/models"
	"github.com/gityokie/gocrud/internal/store"
)

type PutOneBook struct {
	Store store.Store
}

func (h PutOneBook) Handle(c *gin.Context) {

	var book models.Book

	id := c.Params.ByName("id")
	err := h.Store.GetOneBook(&book, id)
	if err != nil {
		api_helpers.RespondJSON(c, 404, book)
	}
	c.BindJSON(&book)

	err = h.Store.PutOneBook(&book, id)

	if err != nil {
		api_helpers.RespondJSON(c, 404, book)
	} else {
		api_helpers.RespondJSON(c, 200, book)
	}
}
