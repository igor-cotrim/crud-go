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

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")

	models.DeleteProduct(productId)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")

	product := models.EditProduct(productId)

	templates.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		amount := r.FormValue("amount")

		convertedIdToInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do ID para int:", err)
		}

		convertedPriceToFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		convertedAmoutToInt, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		models.UpdateProduct(convertedIdToInt, name, description, convertedPriceToFloat, convertedAmoutToInt)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
