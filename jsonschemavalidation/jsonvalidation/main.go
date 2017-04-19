package main

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/xeipuuv/gojsonschema"
)

type User struct {
	Username string `validate:"required,printascii,max=16"`
	Profile  string `validate:"required,jsonschema=profile"`
}

const profileSchema = `
{
	"type":"object",
	"properties":{
		"age": {
			"type":"integer",
			"minimum":0
		}
	}
}
`

var schema = map[string]string{"profile": profileSchema}

func main() {
	v := validator.New()
	v.RegisterValidation("jsonschema", isValidJSON)

	fmt.Println("---- valid values")
	u := User{
		Username: "valid user",
		Profile:  `{"age":20}`,
	}
	if err := v.Struct(u); err == nil {
		fmt.Println("Valid!")
	} else {
		fmt.Println("Invalid!", err)
	}

	fmt.Println("---- invalid values")
	u = User{
		Username: "invalid user",
		Profile:  `{"age":-666}`,
	}
	if err := v.Struct(u); err == nil {
		fmt.Println("Valid!")
	} else {
		fmt.Println("Invalid!", err)
	}
}

func isValidJSON(fl validator.FieldLevel) bool {
	json := fl.Field().String()
	field := fl.Param()

	schemaLoader := gojsonschema.NewStringLoader(schema[field])
	jsonLoader := gojsonschema.NewStringLoader(json)

	result, _ := gojsonschema.Validate(schemaLoader, jsonLoader)
	return result.Valid()
}
