package utils

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

func RespToMapWithRegex(resp string, re *regexp.Regexp) map[string]interface{} {
	matches := re.FindAllStringSubmatch(resp, -1)

	// Create a map to store the key-value pairs
	data := make(map[string]interface{})
	for _, match := range matches {
		key := strings.TrimSpace(match[1])
		value := strings.TrimSpace(strings.ReplaceAll(match[2], "\"", ""))
		data[key] = value
	}

	return data
}

func RespToMap(resp string) map[string]interface{} {
	// Regular expression to extract key-value pairs
	re := regexp.MustCompile(`([^,=\r\n]+)[,=]([^\r\n]+)`)
	return RespToMapWithRegex(resp, re)
}

func MapToStruct(m map[string]interface{}, s interface{}) error {
	// Convert the map to JSON
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		return fmt.Errorf("failed to marshal map to JSON: %w", err)
	}

	// Unmarshal the JSON into the struct
	err = json.Unmarshal(jsonBytes, s)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON to struct: %w", err)
	}

	return nil
}
