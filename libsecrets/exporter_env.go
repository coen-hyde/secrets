package libsecrets

import "strings"

// ExporterEnv formats data for environment variable
type ExporterEnv struct {
	data *map[string]string
}

// NewExporterHuman instantiate a json formater
func NewExporterEnv(data *map[string]string) *ExporterEnv {
	return &ExporterEnv{
		data: data,
	}
}

// String exports the data
func (f *ExporterEnv) String() string {
	output := ""

	for key, value := range *f.data {
		escapedValue := strings.Replace(value, "\"", "\\\"", -1)
		output += "export " + key + "=\"" + escapedValue + "\"\n"
	}

	return string(output)
}
