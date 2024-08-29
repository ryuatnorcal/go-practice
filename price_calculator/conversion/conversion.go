package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	var floats []float64
	for _, line := range strings {
		floatPrice, err := strconv.ParseFloat(line, 64)
		if err != nil {

			return nil, errors.New("Failed to convert string to float")
		}
		floats = append(floats, floatPrice)
	}

	return floats, nil
}
