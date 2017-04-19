package main

import (
	"log"
	"path/filepath"

	"github.com/xeipuuv/gojsonschema"
)

type jsonschema map[string]gojsonschema.JSONLoader

var schemas = []string{
	"notification",
}

func newJSONSchema() *jsonschema {
	js := &jsonschema{}

	for _, name := range schemas {
		path, err := filepath.Abs(name)
		if err != nil {
			log.Fatalln(err)
		}
		(*js)[name] = gojsonschema.NewReferenceLoader("file://" + path + ".json")
	}

	return js
}

func (js *jsonschema) validateJSON(jsonType, json string) bool {
	loader := gojsonschema.NewStringLoader(json)
	schemaLoader, ok := (*js)[jsonType]
	if !ok {
		return false
	}

	result, err := gojsonschema.Validate(schemaLoader, loader)
	if err != nil {
		return false
	}
	return result.Valid()
}
