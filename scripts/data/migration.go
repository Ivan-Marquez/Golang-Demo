package data

import (
	"log"

	"github.com/ivan-marquez/golang-demo/pkg/storage/pq"
)

// Migrate func creates and populates db table
// with data retrieved from Data.gov dataset
func Migrate() error {
	storage, err := pq.NewStorage()
	if err != nil {
		return err
	}

	log.Println("creating table…")
	storage.Migrate()
	log.Println("table successfully created.")

	log.Println("requesting data to be inserted…")
	res, err := MakeRequest()
	if err != nil {
		return err
	}
	log.Println("data retrieved successfully.")

	log.Println("parsing retrieved data…")
	// Process Json data
	rr, err := ParseResponse(res.Body())
	if err != nil {
		return err
	}

	log.Println("inserting data…")
	err = storage.Populate(rr)
	if err != nil {
		return err
	}
	log.Println("data inserted successfully.")

	return nil
}
