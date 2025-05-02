package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func GetCategories(c *gin.Context) {
	categories := []Category{
		{ID: 1, Name: "Статуэтки", Slug: "figurines"},
		{ID: 2, Name: "Посуда", Slug: "dishes"},
		{ID: 3, Name: "Вещи", Slug: "items"},
		{ID: 4, Name: "Мебель", Slug: "furniture"},
		{ID: 5, Name: "Картины", Slug: "paintings"},
	}
	c.JSON(http.StatusOK, categories)
}
