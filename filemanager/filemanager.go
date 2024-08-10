package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)

	//error handling
	if err != nil {

		return nil, errors.New("Failed to open file.")
	}
	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {

		file.Close()
		return nil, errors.New("Failed to read file.")
	}

	file.Close()
	return lines, nil
}

func WriteJSON(path string, data interface{}) error {

	file, err := os.Create(path)

	if err != nil {
		return errors.New("failed to Create file.")
	}

	encoder := json.NewEncoder(file)
	encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("failed to convert data to JSON")
	}

	file.Close()

	return nil
}
