package main

import (
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/go-playground/validator.v9"
)

type notification struct {
	UUID    string    `validate:"required,uuid4"`
	Content string    `validate:"required"`
	Created time.Time `validate:"required"`
}

func main() {
	v := validator.New()

	fmt.Println("---- valid values")
	n := notification{
		UUID:    uuid.NewV4().String(),
		Content: `{}`,
		Created: time.Now(),
	}
	if err := v.Struct(n); err == nil {
		fmt.Println("Valid!")
	} else {
		fmt.Println("Invalid!", err)
	}

	fmt.Println("---- invalid values")
	n = notification{
		UUID:    "invalid",
		Content: `{}`,
		Created: time.Now(),
	}
	if err := v.Struct(n); err == nil {
		fmt.Println("Valid!")
	} else {
		fmt.Println("Invalid!", err)
	}
}
