package libsecrets

type Importer interface {
	Parse(data string) (map[string]interface{}, error)
}
