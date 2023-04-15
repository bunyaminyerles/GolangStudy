package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	// connect to database
	db, err := sql.Open("mysql", "root:1234@/testdb")
	if err != nil {
		panic(err)
	}

	// close database connection
	defer db.Close()

	// create user table
	createStatement := `CREATE TABLE IF NOT EXISTS users (
    		id INT AUTO_INCREMENT PRIMARY KEY,
    		name VARCHAR(50) NOT NULL,
    		email VARCHAR(50) NOT NULL
    		)`
	_, err = db.Exec(createStatement)
	if err != nil {
		log.Fatal(err)
	}

	// insert data
	res, err := db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", "alfa", "alfa@beta.com")
	if err != nil {
		log.Fatal(err)
	}

	// get last inserted id
	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Last inserted ID: %d", lastInsertedID)

	// get affected row count
	affectedRowCount, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Affected rows: %d", affectedRowCount)

	// update data
	res, err = db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", "beta", "beta@alfa.com", lastInsertedID)
	if err != nil {
		log.Fatal(err)
	}

	// get users table
	result, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	for result.Next() {
		var id int
		var name string
		var email string
		err = result.Scan(&id, &name, &email)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ID: %d, Name: %s, Email: %s ", id, name, email)
	}
	result.Close()

	// get single user
	var id int
	var name string
	var email string
	err = db.QueryRow("SELECT * FROM users WHERE id = ?", lastInsertedID).Scan(&id, &name, &email)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID: %d, Name: %s, Email: %s ", id, name, email)

	// prepare statement
	stmt, err := db.Prepare("SELECT * FROM users WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// get single user with prepared statement
	err = stmt.QueryRow(lastInsertedID).Scan(&id, &name, &email)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID: %d, Name: %s, Email: %s ", id, name, email)

	// transaction
	// get users table
	log.Println("Before transaction")
	result, err = db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	for result.Next() {
		var id int
		var name string
		var email string
		err = result.Scan(&id, &name, &email)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ID: %d, Name: %s, Email: %s ", id, name, email)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	_, err = tx.Exec("INSERT INTO users (name, email) VALUES (?, ?)", "alfa", "alfa@alfa.com")
	if err != nil {
		log.Fatal(err)
		tx.Rollback()
	}

	// get users table
	log.Println("Middle of transaction")
	result, err = db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	for result.Next() {
		var id int
		var name string
		var email string
		err = result.Scan(&id, &name, &email)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ID: %d, Name: %s, Email: %s ", id, name, email)
	}

	_, err = tx.Exec("INSERT INTO users (name, email) VALUES (?, ?)", "beta", "beta@beta.com")
	if err != nil {
		log.Fatal(err)
		tx.Rollback()
	}

	tx.Commit()

	// get users table
	log.Println("End of transaction ")
	result, err = db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	for result.Next() {
		var id int
		var name string
		var email string
		err = result.Scan(&id, &name, &email)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ID: %d, Name: %s, Email: %s ", id, name, email)
	}

	// delete data
	res, err = db.Exec("DELETE FROM users WHERE id = ?", lastInsertedID)
	if err != nil {
		log.Fatal(err)
	}

	// drop users table
	_, err = db.Exec("DROP TABLE users")
	if err != nil {
		log.Fatal(err)
	}
}
