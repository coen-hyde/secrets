package libsecrets

import "encoding/json"

// ParserJSON parse's data from json
type ImporterJSON struct {
	Options ImportOptions
}

func NewImporterJSON(options ImportOptions) *ImporterJSON {
	return &ImporterJSON{
		Options: options,
	}
}

// Parse converts the raw json to a structured data
func (f *ImporterJSON) Parse(data string) (map[string]interface{}, error) {
	var structuredData map[string]interface{}
	err := json.Unmarshal([]byte(data), &structuredData)

	if err != nil {
		return nil, err
	}

	return structuredData, nil
}
