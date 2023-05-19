package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/proullon/ramsql/driver"
)

func main() {

	// Connect db
	db, err := sql.Open("ramsql", "goimdb")
	if err != nil {
		log.Fatal("error", err)
		return
	}

	// Create table
	createTb := `
	CREATE TABLE IF NOT EXISTS goimdb (
	id INT AUTO_INCREMENT,
	imdbID TEXT NOT NULL UNIQUE,
	title TEXT NOT NULL,
	year INT NOT NULL,
	rating FLOAT NOT NULL,
	isSuperHero BOOLEAN NOT NULL,
	PRIMARY KEY (id) 
	);
	`
	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("error:", err)
		return
	}

	fmt.Println("table created.")

	// Insert Data
	insert := `
	INSERT INTO goimdb(imdbID,title,year,rating,isSuperHero)
	VALUES (?,?,?,?,?);
	`

	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal("Prepare statement error:", err)
	}

	r, err := stmt.Exec("1234", "Guardian of the Galaxy Vol.3", 2023, 8.4, true)
	if err != nil {
		log.Fatal("Insert error:", err)
	}

	l, err := r.LastInsertId()
	fmt.Println("lastInsertId", l, "error:", err)
	ef, err := r.RowsAffected()
	fmt.Println("RowsAffected", ef, "error:", err)

	// Query all row
	rows, err := db.Query(`
	SELECT id, imdbID, title, year, rating, isSuperHero
	FROM goimdb
	`)
	if err != nil {
		log.Fatal("Query error:", err)
	}

	for rows.Next() {
		var id int
		var imdbID, title string
		var year int
		var rating float32
		var isSuperHero bool

		err := rows.Scan(&id, &imdbID, &title, &year, &rating, &isSuperHero)
		if err != nil {
			log.Fatal("for rows error:", err)
		}
		fmt.Println("row:", id, imdbID, title, year, rating, isSuperHero)
	}

	// Update row in table
	stmt2, err := db.Prepare(`
	UPDATE goimdb
	SET rating=?
	WHERE imdbID=?
	`)

	_, err = stmt2.Exec(9.2, "1234")
	if err != nil {
		log.Fatal("Update error:", err)
	}

	// Query one row in table
	rowx := db.QueryRow(`SELECT id, imdbID, title, year, rating, isSuperHero FROM goimdb WHERE imdbID=?`, "1234")
	var id int
	var imdbID, title string
	var year int
	var rating float32
	var isSuperHero bool
	err = rowx.Scan(&id, &imdbID, &title, &year, &rating, &isSuperHero)
	if err != nil {
		log.Fatal("Scan one rowx error:", err)
	}

	fmt.Println("one rowx:", id, imdbID, title, year, rating, isSuperHero)
}
