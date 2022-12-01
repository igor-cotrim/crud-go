package models

import "main/db"

type Product struct {
	Name, Description string
	Price             float64
	Id, Amount        int
}

func SearchAllProducts() []Product {
	db := db.ConnectDatabase()

	selectFromAllProducts, err := db.Query("select * from products")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectFromAllProducts.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = selectFromAllProducts.Scan(&id, &name, &description, &price, &amount)

		if err != nil {
			panic(err.Error())
		}

		p.Name = name
		p.Description = description
		p.Price = price
		p.Amount = amount

		products = append(products, p)
	}

	defer db.Close()

	return products
}
