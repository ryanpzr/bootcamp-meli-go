package loader

import (
	"app/internal"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func NewLoaderTicketCSV(filePath string) (*LoaderTicketCSV, error) {
	return &LoaderTicketCSV{
		filePath: filePath,
	}, nil
}

type LoaderTicketCSV struct {
	filePath string
}

func (t *LoaderTicketCSV) Load() (map[int]internal.TicketAttributes, error) {
	f, err := os.Open(t.filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer f.Close()

	r := csv.NewReader(f)

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

		ti[idInt] = ticket
	}

	return ti, nil
}
