package parser

import (
	"gopkg.in/yaml.v3"
)

// Протестировать под Фаззинг-тестированием
func Parse(data []byte) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := yaml.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	// if len(data) > 20 { // тестирование отработки ошибки для Фаззинга
	// 	panic("len is bigger")
	// }
	return result, nil
}
