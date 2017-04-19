package main

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	Username string `validate:"required,printascii,max=16"`
	Name     string
	Email    string `validate:"required,email"`
	Age      int    `validate:"min=0"`
}

func main() {
	v := validator.New()

	fmt.Println("---- valid values")
	u := User{
		Username: "valid user",
		Email:    "valid@example.com",
		Age:      20,
	}
	if err := v.Struct(u); err == nil {
		fmt.Println("Valid!")
	} else {
		fmt.Println("Invalid!", err)
	}

	fmt.Println("---- invalid values")
	u = User{
		Username: "invalid user",
		Email:    "invalid",
		Age:      20,
	}
	if err := v.Struct(u); err == nil {
		fmt.Println("Valid!")
	} else {
		fmt.Println("Invalid!", err)
	}
}
