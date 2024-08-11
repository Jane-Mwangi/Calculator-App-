package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct{
	InputFilePath string
	OutputFilePath string
}

func (fm FileManager)ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

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

func (fm FileManager)WriteResult(data interface{}) error {

	file, err := os.Create(fm.OutputFilePath)

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


func New(inputFilePath , outputFilePath string) FileManager {
	return FileManager{
		InputFilePath: inputFilePath,
		OutputFilePath: outputFilePath,
	}
}