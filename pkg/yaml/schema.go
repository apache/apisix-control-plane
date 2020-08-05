package yaml

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
)

// YamlSchema define schema rule
func YamlSchema() string {
	return `
{
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
		},
		"hosts": {
			"type": "array",
			"items": {
				"type": "string"
			}
		},
		"gateways": {
			"type": "array",
			"items": {
				"type": "string"
			}
		},
		"http": {
			"type": "array",
			"items": {
				"type": "object",
				"properties": {
					"route": {
						"type": "array",
						"items": {
							"type": "object",
							"properties": {
								"destination": {
									"type": "object",
									"properties": {
										"port": {
											"type": "number"
										},
										"host": {
											"type": "string"
										},
										"subset": {
											"type": "string"
										},
										"weight": {
											"type": "number"
										}
									}
								}
							}
						}
					},
					"label": {
						"type": "object"
					},
					"match": {
						"type": "array",
						"items": {
							"type": "object",
							"properties": {
								"headers": {
									"type": "object"
								}
							}
						}
					}
				}
			}
		},
		"host": {
			"type": "string"
		},
		"subsets": {
			"type": "array",
			"items": {
				"type": "object",
				"properties": {
					"name": {
						"type": "string"
					},
					"ips": {
						"type": "array",
						"items": {
							"type": "string"
						}
					},
					"selector": {
						"type": "object",
						"properties": {
							"labels": {
								"type": "object"
							}
						}
					}
				}
			}
		},
		"selector": {
			"type": "object",
			"properties": {
				"labels": {
					"type": "object"
				}
			}
		},
		"sort": {
			"type": "array",
			"items": {
				"type": "object",
				"properties": {
					"name": {
						"type": "string"
					},
					"conf": {
						"type": "object"
					}
				}
			}
		}
	},
	"additionalProperties": true
}
`
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
