package libsecrets

import (
	"gopkg.in/yaml.v2"
)

// ExporterYAML formats data as yaml
type ExporterYAML struct {
	data *map[string]string
}

// NewExporterYAML instantiate a yaml formater
func NewExporterYAML(data *map[string]string) *ExporterYAML {
	return &ExporterYAML{
		data: data,
	}
}

// String exports the scope data in yaml format
func (f *ExporterYAML) String() string {
	data, err := yaml.Marshal(f.data)

	if err != nil {
		panic("Could not marshal Scope data into yaml")
	}

	return string(data)
}
