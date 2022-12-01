package controllers

import (
	"html/template"
	"log"
	"main/models"
	"net/http"
	"strconv"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.SearchAllProducts()

	templates.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		amount := r.FormValue("amount")

		convertedPriceToFloat, err := strconv.ParseFloat(price, 64)

		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		convertedAmoutToInt, err := strconv.Atoi(amount)

		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		models.CreateNewProduct(name, description, convertedPriceToFloat, convertedAmoutToInt)
	}

	http.Redirect(w, r, "/", 301)
}
