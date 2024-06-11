package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Item struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Terms       string  `json:"terms"`
	Sku         string  `json:"sku"`
	Price       float32 `json:"price"`
	Images      string  `json:"images"`
	Section     string  `json:"section"`
}

type Query struct {
	Data []*Item `json:"data"`
}

func main() {
	//connect to a database
	connStr := "user=root port=5432 dbname=root password=secret sslmode=disable"
	conn, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("Unable to connect: %v\n", err)
	}

	defer conn.Close()

	log.Print("Connected to database!")

	// Handlers

	http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		data, err := getAllRows(conn)
		if err != nil {
			fmt.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(data)

	})

	http.ListenAndServe(":1000", nil)

}

func getAllRows(conn *sql.DB) ([]byte, error) {
	rows, err := conn.Query("select * from zara limit 9")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	var name, description, terms, section, sku string
	var price float32
	var images string

	var data []Item

	for rows.Next() {

		err := rows.Scan(&sku, &name, &description, &price, &images, &terms, &section)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		fmt.Println("Record is", sku, name, description, price, images, terms, section)

		row := Item{Sku: sku, Name: name, Description: description, Price: price, Images: images, Terms: terms, Section: section}

		data = append(data, row)

	}

	if err = rows.Err(); err != nil {
		log.Fatal("Error scanning rows", err)
		return nil, err
	}

	fmt.Println("-----------------------------------")

	return json.Marshal(data)

}
