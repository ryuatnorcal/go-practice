// folder name is package name
package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// to export wrute file name start with uppercase
func WriteFloatToFile(value float64, filename string) {
	valueText := fmt.Sprintf("%.2f", value)
	// filename, byte, permission
	os.WriteFile(filename, []byte(valueText), 0644)
}
func GetFloatFromFile(fileName string, defaultValue float64) (float64, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return defaultValue, errors.New("Failed to find file") // default value
	}
	balanceText := string(file)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return defaultValue, errors.New("Failed to parse file") // default value
	}
	return balance, nil
}
