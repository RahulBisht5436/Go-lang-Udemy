package utility

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func FileBytesToSlice(data []byte) ([]float64, error) {
	// 1. Convert to string and strip outer spacing and brackets
	str := strings.TrimSpace(string(data))
	str = strings.Trim(str, "[]")

	// Handle empty input gracefully
	if str == "" {
		return []float64{}, nil
	}

	// 2. Split the string by commas
	parts := strings.Split(str, "\r\n")
	var result []float64

	// 3. Convert each string segment into an integer
	for _, part := range parts {
		trimmedPart := strings.TrimSpace(part)
		if trimmedPart == "" {
			continue // Skips extra or trailing commas
		}

		num, err := strconv.ParseFloat(trimmedPart, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid number %q: %w", trimmedPart, err)
		}
		if CheckFormat(num, "float64") {
			result = append(result, num)
		}
	}

	return result, nil
}

func CheckFormat(data any, format string) bool {

	switch format {

	case "string":
		_, ok := data.(string)
		return ok

	case "int":
		_, ok := data.(int)
		return ok

	case "float64":
		_, ok := data.(float64)
		return ok

	case "bool":
		_, ok := data.(bool)
		return ok

	default:
		return false
	}
}

func ReadFiles(fileName string) ([]float64, error) {
	fmt.Println(fileName)
	priceInfoRawContent, priceInfoRawErr := os.ReadFile(fileName)
	if priceInfoRawErr != nil {
		return []float64{}, errors.New("Not able to read data")
	}
	pricesInfo, priceInfoErr := FileBytesToSlice(priceInfoRawContent)
	if priceInfoErr != nil {
		return []float64{}, priceInfoErr
	}

	return pricesInfo, nil

}

func WriteJSON(path string, data any) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("unable to create file %q: %w", path, err)
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(data); err != nil {
		return fmt.Errorf("failed to encode data to %q: %w", path, err)
	}
	return nil
}
