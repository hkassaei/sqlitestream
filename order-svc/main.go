package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Car represents a car with the following properties:
// model, make, year, mileage, color, and price.
type Car struct {
	Model   string
	Make    string
	Year    int
	Mileage int
	Color   string
	Price   int
}

func main() {

	orderid := 0
	const timeFormat = "2006-01-02 15:04:05" //slightly modified version of RFC3339 format in the time package
	dbPath := "/data/cars.db"
	// Open a connection to the SQLite database.
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the "cars" table
	_, err = db.Exec(`DROP TABLE IF EXISTS cars`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`
		CREATE TABLE cars (
			orderid SERIAL NOT NULL PRIMARY KEY,
			make TEXT NOT NULL,
			year INTEGER NOY NULL,
			mileage INTEGER NOT NULL,
			color TEXT NOY NULL,
			price INTEGER NOT NULL,
			order_time TEXT NOT NULL
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Seed the random number generator.
	rand.Seed(time.Now().UnixNano())

	for {
		orderid++
		// Create a random car.
		car := &Car{
			Make:    []string{"Acura", "Audi", "BWM", "Honda", "Hyundai", "Kia", "Mercedes", "Nissan", "Porsche", "Tesla", "Volvo"}[rand.Intn(11)],
			Year:    rand.Intn(23) + 2000,
			Mileage: rand.Intn(200000-5000) + 5000,
			Color:   []string{"Blue", "Green", "Red", "Black", "White", "Silver", "Navy", "Orange"}[rand.Intn(8)],
			Price:   rand.Intn(100000-5000) + 5000,
		}

		// Insert the random car into the "cars" table.
		orderTime := time.Now().UTC().Format(timeFormat)
		_, err = db.Exec(`
		INSERT INTO cars (orderid, make, year, mileage, color, price, order_time)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`, orderid, car.Make, car.Year, car.Mileage, car.Color, car.Price, orderTime)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("new order: %d %s %d %d %s %d %s\n", orderid, car.Make, car.Year, car.Mileage, car.Color, car.Price, orderTime)
		time.Sleep(10 * time.Second)
	}
}
