package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gityokie/gocrud/internal/api_helpers"
	"github.com/gityokie/gocrud/internal/models"
	"github.com/gityokie/gocrud/internal/store"
)

type GetOneBook struct {
	Store store.Store
}

func (h GetOneBook) Handle(c *gin.Context) {

	id := c.Params.ByName("id")

	var book models.Book

	err := h.Store.GetOneBook(&book, id)

	if err != nil {
		api_helpers.RespondJSON(c, 404, book)
	} else {
		api_helpers.RespondJSON(c, 200, book)
	}
}
