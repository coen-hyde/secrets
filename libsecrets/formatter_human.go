package libsecrets

// FormatterHuman formats data for human consumption
type FormatterHuman struct {
	data *map[string]string
}

// NewFormatterHuman instantiate a json formater
func NewFormatterHuman(data *map[string]string) *FormatterHuman {
	return &FormatterHuman{
		data: data,
	}
}

// String exports the data for humans
func (f *FormatterHuman) String() string {
	output := ""

	for key, value := range *f.data {
		output += key + ": " + value
	}

	return string(output)
}
