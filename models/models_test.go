package models

import (
	"encoding/json"
	"testing"
)

func TestParseModel(t *testing.T) {
	jsonString := `
	[
		{
			"path": "/path/to/json/resource",
			"response": {
				"type": "JSON",
				"value": "/path/to/json",
				"status_code": 200
			}
		}
	]
	`

	jsonByte := []byte(jsonString)
	result := make([]Rule, 0)
	err := json.Unmarshal(jsonByte, &result)

	if err != nil {
		t.Errorf("Error parsing, %s", err.Error())
		return
	}

	rule := result[0]
	expectedRule := Rule{
		Path: "/path/to/json/resource",
		Response: Response{
			Type:       "JSON",
			Value:      "/path/to/json",
			StatusCode: 200,
		},
	}

	if rule != expectedRule {
		t.Errorf("Match rule mismatch, expected %v got %v", expectedRule, rule)
	}
}
