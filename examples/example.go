package main

import (
	"fmt"

	"github.com/amineck/purify"
)

type userForm struct {
	Name        string
	Email       string
	EmailSHA256 string
}

func main() {
	// Struct Example
	form := userForm{
		Name:        " john doe123 $%@#",
		Email:       " John@EXAMPLE.com ",
		EmailSHA256: "John@EXAMPLE.com ",
	}
	err := purify.SanitizeStruct(&form,
		purify.Field(&form.Name, purify.ToAlphaNumeric, purify.TrimSpace, purify.ToTitleCase),
		purify.Field(&form.Email, purify.TrimSpace, purify.ToEmail),
		purify.Field(&form.EmailSHA256, purify.TrimSpace, purify.ToEmail, purify.When(true, purify.ToSHA256)),
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Output: %+v\n", form)
	// Output: {Name:John Doe123 Email:John@example.com EmailSHA256:20740450eae791b7928c1869d3a0c964e8685544eb0d70c386d6ba825270b12e}

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
	fmt.Printf("Output: %+v\n", m)
	// Output: map[Email:John@examplecom Name:John Doe]
}
