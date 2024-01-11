package main

import (
	"fmt"

	"github.com/amineck/purify"
)

type userForm struct {
	Name  string
	Email string
}

func main() {
	// Struct Example
	form := userForm{
		Name:  " john doe123 ",
		Email: " John@EXAMPLE.com ",
	}
	err := purify.SanitizeStruct(&form,
		purify.Field(&form.Name, purify.TrimSpace, purify.ToAlpha, purify.ToTitleCase),
		purify.Field(&form.Email, purify.TrimSpace, purify.ToEmail),
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Output: %+v\n", form)
	// Output: {Name:John Doe Email:John@examplecom}

	m := map[string]string{
		"name":  " john doe123 ",
		"email": " John@EXAMPLE.com ",
	}
	err = purify.SanitizeMap(m,
		purify.Key("name", purify.TrimSpace, purify.ToAlpha, purify.ToTitleCase),
		purify.Key("email", purify.TrimSpace, purify.ToEmail),
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", form)
	// Output: map[Email:John@examplecom Name:John Doe]
}
