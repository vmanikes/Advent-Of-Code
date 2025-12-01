package helpers

import (
	"log"
	"os"
)

func GetFile(filename string) (*os.File, error) {
	inputFile, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		log.Println("unable to open file: ", err)
		return nil, err
	}

	return inputFile, nil
}
