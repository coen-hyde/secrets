package libsecrets

import "strings"

// FormatterEnv formats data for environment variable
type FormatterEnv struct {
	data *map[string]string
}

// NewFormatterHuman instantiate a json formater
func NewFormatterEnv(data *map[string]string) *FormatterEnv {
	return &FormatterEnv{
		data: data,
	}
}

// String exports the data
func (f *FormatterEnv) String() string {
	output := ""

	for key, value := range *f.data {
		escapedValue := strings.Replace(value, "\"", "\\\"", -1)
		output += "export " + key + "=\"" + escapedValue + "\"\n"
	}

	return string(output)
}
