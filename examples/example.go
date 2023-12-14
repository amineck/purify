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
	form := userForm{
		Name:  " john doe123 ",
		Email: " John@EXAMPLE.com ",
	}
	purify.SanitizeStruct(&form,
		purify.Field(&form.Name, purify.TrimSpace, purify.ToAlpha, purify.ToTitleCase),
		purify.Field(&form.Email, purify.TrimSpace, purify.ToEmail),
	)

	fmt.Printf("%+v\n", form)
}

// Output:
// {Name:John Doe Email:John@examplecom}
