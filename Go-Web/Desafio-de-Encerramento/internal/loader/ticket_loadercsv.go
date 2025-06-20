package loader

import (
	"app/internal"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// NewLoaderTicketCSV creates a new ticket loader from a CSV file
func NewLoaderTicketCSV(filePath string) (*LoaderTicketCSV, error) {
	return &LoaderTicketCSV{
		filePath: filePath,
	}, nil
}

// LoaderTicketCSV represents a ticket loader from a CSV file
type LoaderTicketCSV struct {
	filePath string
}

// Load loads the tickets from the CSV file
func (t *LoaderTicketCSV) Load() (map[int]internal.TicketAttributes, error) {
	// open the file
	f, err := os.Open(t.filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer f.Close()

	// read the file
	r := csv.NewReader(f)

	// read the records
	ti := make(map[int]internal.TicketAttributes)
	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			err = fmt.Errorf("error reading record: %v", err)
			return nil, err
		}

		// serialize the record
		id := record[0]
		idInt, _ := strconv.Atoi(id)

		priceint, _ := strconv.ParseFloat(record[5], 64)
		ticket := internal.TicketAttributes{
			Name:    record[1],
			Email:   record[2],
			Country: record[3],
			Hour:    record[4],
			Price:   priceint,
		}

		// add the ticket to the map
		ti[idInt] = ticket
	}

	return ti, nil
}
