package main

import "github.com/xeipuuv/gojsonschema"

type jsonschema struct {
	schemaLoaders map[string]gojsonschema.JSONLoader
}

const notificationSchema = "notification.json"

func newJSONSchema() *jsonschema {
	j := &jsonschema{}
	j.schemaLoaders["notification"] =
		gojsonschema.NewReferenceLoader("file://" + notificationSchema)
	return j
}

func (js *jsonschema) ValidateJSON(jsonType, json string) bool {
	loader := gojsonschema.NewStringLoader(json)
	schemaLoader := js.schemaLoaders[jsonType]
	result, _ := gojsonschema.Validate(schemaLoader, loader)
	return result
}
