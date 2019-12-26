package config

type typeFile string

const (
	TypeYaml    typeFile = "yaml"
	TypeJson    typeFile = "json"
	TypeUnknown typeFile = "unknown"
)

var implementedTypes = map[string]typeFile{".yaml": TypeYaml, ".yml": TypeYaml, ".json": TypeJson}

func detectType(ext string) typeFile {
	t, found := implementedTypes[ext]
	if !found {
		return TypeUnknown
	}
	return t
}
