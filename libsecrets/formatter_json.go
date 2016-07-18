package libsecrets

import (
	"encoding/json"
)

// FormatterJSON formats data as json
type FormatterJSON struct {
	data *map[string]string
}

// NewFormatterJSON instantiate a json formater
func NewFormatterJSON(data *map[string]string) *FormatterJSON {
	return &FormatterJSON{
		data: data,
	}
}

// String exports the scope data in json format
func (f *FormatterJSON) String() string {
	jsonData, err := json.MarshalIndent(f.data, "", "  ")

	if err != nil {
		panic("Could not marshal Scope data into json")
	}

	return string(jsonData)
}
