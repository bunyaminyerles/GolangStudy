package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"
)

var db *sql.DB
var err error

// Product struct
type Product struct {
	id    int
	name  string
	price float64
}

// init function
func init() {

	// create connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// connect to database
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Successfully connected!")

}

// create table
func createTable() sql.Result {
	res, err := db.Exec("CREATE TABLE IF NOT EXISTS products (id SERIAL PRIMARY KEY, name VARCHAR(255), price NUMERIC(10,2))")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully created table!")

	return res
}

// insert data
func insertData(product Product) {
	_, err := db.Exec("INSERT INTO products (name, price) VALUES ($1, $2)", product.name, product.price)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Inserted: %s, %f", product.name, product.price)
}

// get data
func getData() *sql.Rows {
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No rows were returned!")
			return rows
		}
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Product
		err = rows.Scan(&p.id, &p.name, &p.price)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(p)
	}
	defer rows.Close()
	return rows
}

// update data
func updateData(product Product) sql.Result {
	res, err := db.Exec("UPDATE products SET name = $1, price = $2 WHERE id = $3", product.name, product.price, product.id)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Updated: %s, %f", product.name, product.price)

	return res
}

// delete data
func deleteData(id int) {
	_, err := db.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Deleted: %d", id)
}

// get product by id
func getProductById(id int) Product {
	var p Product
	err := db.QueryRow("SELECT * FROM products WHERE id = $1", id).Scan(&p.id, &p.name, &p.price)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No rows were returned!")
			return p
		}
		log.Fatal(err)
	}
	return p
}

// drop table
func dropTable() sql.Result {
	res, err := db.Exec("DROP TABLE products")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Table dropped!")

	return res
}

func main() {

	// close database connection
	defer db.Close()

	// create table
	createTable()

	// insert data
	insertData(Product{name: "Laptop", price: 2000.00})
	insertData(Product{name: "Mouse", price: 20.00})
	insertData(Product{name: "Keyboard", price: 50.00})

	// get data
	getData()

	// update data
	updateData(Product{id: 3, name: "Laptop", price: 3000.00})

	// get product by id
	p1 := getProductById(3)
	log.Println(p1)

	// delete data
	deleteData(3)

	// drop table
	dropTable()

}
