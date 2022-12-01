package controllers

import (
	"html/template"
	"main/models"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.SearchAllProducts()

	templates.ExecuteTemplate(w, "Index", allProducts)
}
