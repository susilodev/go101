package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// returning float from file
func GetFloatFromFile(file string) (float64, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return 0.0, errors.New("can't ReadFile")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 0.0, errors.New("failed parse float64")
	}

	return value, nil
}

// write float value to spesifict file
func WriteFloatToFile(value float64, fileLocation string) (float64 error) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileLocation, []byte(valueText), 0644)
	return
}
