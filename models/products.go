package models

import "main/db"

type Product struct {
	Name, Description string
	Price             float64
	Id, Amount        int
}

func SearchAllProducts() []Product {
	db := db.ConnectDatabase()

	selectFromAllProducts, err := db.Query("select * from products order by id asc")
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

		p.Id = id
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

func DeleteProduct(id string) {
	db := db.ConnectDatabase()

	deleteProduct, err := db.Prepare("delete from products where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)

	defer db.Close()
}

func EditProduct(id string) Product {
	db := db.ConnectDatabase()

	product, err := db.Query("select * from products where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	productToUpdate := Product{}

	for product.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = product.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}

		productToUpdate.Id = id
		productToUpdate.Name = name
		productToUpdate.Description = description
		productToUpdate.Price = price
		productToUpdate.Amount = amount
	}

	defer db.Close()

	return productToUpdate
}

func UpdateProduct(id int, name, description string, price float64, amount int) {
	db := db.ConnectDatabase()

	insertDataInDatabase, err := db.Prepare("update products set name=$2, description=$3, price=$4, amount=$5 where id=$1")
	if err != nil {
		panic(err.Error())
	}

	insertDataInDatabase.Exec(id, name, description, price, amount)

	defer db.Close()
}
