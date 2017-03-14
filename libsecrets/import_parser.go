package libsecrets

type ImportParser interface {
	Parse(rawData string) (map[string]string, error)
}
