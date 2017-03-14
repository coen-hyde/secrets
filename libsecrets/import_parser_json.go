package libsecrets

import "encoding/json"

// ParserJSON parse's data from json
type ImportParserJSON struct{}

func NewImportParserJSON() *ImportParserJSON {
	return &ImportParserJSON{}
}

// Parse converts the raw json to a structured data
func (f *ImportParserJSON) Parse(rawJson string) (map[string]string, error) {
	structuredData := map[string]string{}
	err := json.Unmarshal([]byte(rawJson), &structuredData)

	if err != nil {
		return nil, err
	}

	return structuredData, nil
}
