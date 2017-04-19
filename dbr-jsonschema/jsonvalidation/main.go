package main

import (
	"fmt"
	"time"

	"github.com/go-playground/validator"
	uuid "github.com/satori/go.uuid"
	"github.com/xeipuuv/gojsonschema"
)

type notification struct {
	UUID    string    `validate:"required,uuid4"`
	Content string    `validate:"required,jsonschema"`
	Created time.Time `validate:"required"`
}

func main() {
	v := validator.New()
	v.RegisterValidation("jsonschema", isValidJSON)

	fmt.Println("---- valid values")
	n := notification{
		UUID:    uuid.NewV4().String(),
		Content: `"string"`,
		Created: time.Now(),
	}
	if err := v.Struct(n); err == nil {
		fmt.Println("Valid!")
	} else {
		fmt.Println("Invalid!", err)
	}

	fmt.Println("---- invalid values")
	n = notification{
		UUID:    uuid.NewV4().String(),
		Content: `{}`,
		Created: time.Now(),
	}
	if err := v.Struct(n); err == nil {
		fmt.Println("Valid!")
	} else {
		fmt.Println("Invalid!", err)
	}
}

func isValidJSON(fl validator.FieldLevel) bool {
	schemaLoader := gojsonschema.NewStringLoader(`{"type":"string"}`)

	json := fl.Field().String()
	jsonLoader := gojsonschema.NewStringLoader(json)

	result, _ := gojsonschema.Validate(schemaLoader, jsonLoader)
	return result.Valid()
}
