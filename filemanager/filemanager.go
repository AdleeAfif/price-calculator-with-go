package filemanager

import (
	"bufio"
	"errors"
	"os"
)

func ReadLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("unable to open the file")
	}

	scanner := bufio.NewScanner(file)
	var fileLine []string
	for scanner.Scan() {
		fileLine = append(fileLine, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("unable to read the file")
	}

	file.Close()
	return fileLine, nil
}
