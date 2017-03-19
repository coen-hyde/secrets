package libsecrets

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

// ParserJSON parse's data from json
type ImporterYAML struct{}

func NewImporterYAML() *ImporterYAML {
	return &ImporterYAML{}
}

// Parse converts the raw yaml to a structured data
func (f *ImporterYAML) Parse(data string) (map[string]interface{}, error) {
	var structuredData map[string]interface{}
	fmt.Println(data)
	err := yaml.Unmarshal([]byte(data), &structuredData)

	if err != nil {
		return nil, err
	}

	return structuredData, nil
}
