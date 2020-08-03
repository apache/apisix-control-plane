package yaml

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
)

// YamlSchema define schema rule
func YamlSchema() string {
	return `{
	"type": "object",
	"properties": {
		"kind": {
			"type": "string",
			"enum": ["Gateway", "Rule", "Destination", "Plugin"]
		},
		"name": {
			"type": "string"
		},
		"servers": {
			"type": "array",
			"minItems" : 1,
			"items": {
				"type": "object",
				"properties": {
					"port": {
						"type": "number",
						"minimum": 1
					},
					"name": {
						"type": "string"
					},
					"protocol": {
						"type": "string",
						"enum": ["HTTP", "TCP"]
					},
					"hosts": {
						"type": "array",
						"items": {
							"type": "string"
						}
					}
				}
			}
		}
	},
	"additionalProperties": true
}`
}

func Validate(request string) (bool, error) {
	schemaLoader := gojsonschema.NewStringLoader(YamlSchema())
	schema, err := gojsonschema.NewSchema(schemaLoader)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	requestLoader := gojsonschema.NewStringLoader(request)
	result, err := schema.Validate(requestLoader)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	if result.Valid() {
		return true, nil
	} else {
		return false, fmt.Errorf(result.Errors()[0].String())
	}
}
