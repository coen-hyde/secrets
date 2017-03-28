package libsecrets

import (
	"gopkg.in/yaml.v2"
)

// ParserJSON parse's data from json
type ImporterYAML struct {
	Options ImportOptions
}

func NewImporterYAML(options ImportOptions) *ImporterYAML {
	return &ImporterYAML{
		Options: options,
	}
}

// Parse converts the raw yaml to a structured data
func (f *ImporterYAML) Parse(data string) (map[string]interface{}, error) {
	var structuredData map[string]interface{}
	err := yaml.Unmarshal([]byte(data), &structuredData)

	if err != nil {
		return nil, err
	}

	return structuredData, nil
}
