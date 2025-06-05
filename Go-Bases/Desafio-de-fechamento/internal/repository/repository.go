package repository

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type repository struct {
}

type Repository interface {
	GetInfoCsv(string) ([][]string, error)
}

func NewRepository() *repository {
	return &repository{}
}

func (r *repository) GetInfoCsv(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	rd := csv.NewReader(file)

	var allRecords [][]string

	for {
		record, err := rd.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Linha ignorada por erro:", err)
			continue
		}
		allRecords = append(allRecords, record)
	}

	return allRecords, nil
}
