package main

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

func main() {
	schemaLoader := gojsonschema.NewStringLoader(`{"type":"string"}`)

	fmt.Println("---- Valid JSON")
	jsonLoader := gojsonschema.NewStringLoader(`"string"`)
	result, _ := gojsonschema.Validate(schemaLoader, jsonLoader)
	if result.Valid() {
		fmt.Println("Valid!")
	} else {
		fmt.Println("Invalid!")
	}

	fmt.Println("---- Invalid JSON")
	jsonLoader = gojsonschema.NewStringLoader(`{}`)
	result, _ = gojsonschema.Validate(schemaLoader, jsonLoader)
	if result.Valid() {
		fmt.Println("Valid!")
	} else {
		fmt.Println("Invalid!")
	}
}
