package parser

import (
	"gopkg.in/yaml.v3"
)

func Parse(data []byte) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := yaml.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
