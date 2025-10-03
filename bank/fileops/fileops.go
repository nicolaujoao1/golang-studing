package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}
func GetFloatFromFile(fileName string) (float64, error) {
	data, error := os.ReadFile(fileName)

	if error != nil {
		return 0, errors.New("error reading balance file")
	}
	balanceText := string(data)
	balance, error := strconv.ParseFloat(balanceText, 64)

	if error != nil {
		return 0, errors.New("error parsing balance")
	}
	return balance, nil
}
