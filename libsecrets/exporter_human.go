package libsecrets

// ExporterHuman formats data for human consumption
type ExporterHuman struct {
	data *map[string]string
}

// NewExporterHuman instantiate a json formater
func NewExporterHuman(data *map[string]string) *ExporterHuman {
	return &ExporterHuman{
		data: data,
	}
}

// String exports the data for humans
func (f *ExporterHuman) String() string {
	output := ""

	for key, value := range *f.data {
		output += key + ": " + value + "\n"
	}

	return string(output)
}
