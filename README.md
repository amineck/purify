# purify [![Godoc](https://godoc.org/github.com/amineck/purify?status.svg)](https://godoc.org/github.com/amineck/purify)

An idiomatic Go package for sanitizing user input, inspired by [ozzo-validation](https://github.com/go-ozzo/ozzo-validation).

## Why?

Most sanitization packages use stuct tags to define sanitization rules. This is a great approach, but:
- It's not always possible to add struct tags to your models. (ex: when using protobuf types)
- Struct tags can be error-prone and hard to read.

This package allows you to define sanitization rules in a more declarative way.

## Installation

```bash
go get github.com/amineck/scrub
```

## Usage

```go
type userForm struct {
	Name  string
	Email string
}

func main() {
	form := userForm {
		Name:  " john doe123 ",
		Email: " John@EXAMPLE.com ",
	}
	scrub.SanitizeStruct(&form,
		scrub.Field(&form.Name, scrub.TrimSpace, scrub.ToAlpha, scrub.ToTitleCase),
		scrub.Field(&form.Email, scrub.TrimSpace, scrub.ToEmail),
	)

	fmt.Printf("%+v\n", form)
}

// Output:
// {Name:John Doe Email:John@example.com}
```