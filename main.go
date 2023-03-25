package main

import (
	"example.com/backend"
)

type Product struct {
	id int
	name string
	inventory int
	price int
}

func main() {
	app := backend.App{}

	app.Port = ":9003"
	app.Initialize()
	app.Run()
}


// func main() {
// 	fmt.Println()
// 	db, err := sql.Open("sqlite3", "./practiceit.db")

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	rows, err := db.Query("SELECT * FROM products")

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		var p Product
// 		rows.Scan(&p.id, &p.name, &p.inventory, &p.price)
// 		fmt.Println("Product: ", p.id, " ", p.name, " ", p.inventory, " ", p.price)
// 	}
// }