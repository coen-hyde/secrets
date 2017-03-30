package libsecrets

import (
	"encoding/json"
)

// ExporterJSON formats data as json
type ExporterJSON struct {
	data *map[string]string
}

// NewExporterJSON instantiate a json formater
func NewExporterJSON(data *map[string]string) *ExporterJSON {
	return &ExporterJSON{
		data: data,
	}
}

// String exports the scope data in json format
func (f *ExporterJSON) String() string {
	jsonData, err := json.MarshalIndent(f.data, "", "  ")

	if err != nil {
		panic("Could not marshal Scope data into json")
	}

	return string(jsonData)
}
