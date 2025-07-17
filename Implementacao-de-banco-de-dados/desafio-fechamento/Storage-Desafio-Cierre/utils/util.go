package utils

import (
	"fmt"
	"log"
	"os"
)

func OpenFile(fileName string) (*os.File, error) {
	file, err := os.Open(fmt.Sprintf("../docs/db/json/%s", fileName))
	if err != nil {
		log.Fatal(err)
	}

	return file, nil
}
