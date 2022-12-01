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

func CreateNewProduct(name, description string, price float64, amount int) {
	db := db.ConnectDatabase()

	insertDataInDatabase, err := db.Prepare("insert into products(name, description, price, amount) values($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insertDataInDatabase.Exec(name, description, price, amount)

	defer db.Close()
}
