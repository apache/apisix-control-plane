package yaml

import "github.com/ghodss/yaml"

func ToJson(y string) ([]byte, error) {
	return yaml.YAMLToJSON([]byte(y))
}
