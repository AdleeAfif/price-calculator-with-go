package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, errors.New("unable to open the file")
	}

	defer file.Close() //execute once the function return

	scanner := bufio.NewScanner(file)
	var fileLine []string
	for scanner.Scan() {
		fileLine = append(fileLine, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		// file.Close()
		return nil, errors.New("unable to read the file")
	}

	// file.Close()
	return fileLine, nil
}

func (fm FileManager) WriteJSON(data interface{}) error { //or any
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("unable to create the file")
	}

	defer file.Close()

	time.Sleep(3 * time.Second) //simulate long load file processing

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("unable to write the file")
	}

	return nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{inputPath, outputPath}
}
