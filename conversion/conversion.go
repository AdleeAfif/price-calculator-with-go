package conversion

import (
	"errors"
	"strconv"
)

func StringToFloat(stringValues []string) ([]float64, error) {
	var floatValues []float64
	for _, sValue := range stringValues {
		floatValue, err := strconv.ParseFloat(sValue, 64)
		if err != nil {
			return nil, errors.New("unable to convert string to float")
		}
		floatValues = append(floatValues, floatValue)
	}
	return floatValues, nil
}
