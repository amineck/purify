# purify [![Godoc](https://godoc.org/github.com/amineck/purify?status.svg)](https://godoc.org/github.com/amineck/purify)

An idiomatic Go package for sanitizing user input, inspired by [ozzo-validation](https://github.com/go-ozzo/ozzo-validation).

## Why?

Most sanitization packages use stuct tags to define sanitization rules. This is a great approach, but:
- It's not always possible to add struct tags to your models. (ex: when using protobuf types)
- Struct tags can be error-prone and hard to read.

This package allows you to define sanitization rules in a more declarative way.

## Installation

```bash
go get github.com/amineck/purify
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
    purify.SanitizeStruct(&form,
        purify.Field(&form.Name, purify.TrimSpace, purify.ToAlpha, purify.ToTitleCase),
        purify.Field(&form.Email, purify.TrimSpace, purify.ToEmail),
    )
    fmt.Printf("%+v\n", form)
}

// Output:
// {Name:John Doe Email:John@example.com}
```

## Functions

* TrimSpace: Trim leading and trailing spaces. `"  hello world  "` => `"hello world"`
* ToCamelCase: Convert to camel case. `"hello world"` => `"helloWorld"`
* ToKebabCase: Convert to kebab case. `"hello world"` => `"hello-world"`
* ToSnakeCase: Convert to snake case. `"hello world"` => `"hello_world"`
* ToTitleCase: Convert to title case. `"hello world"` => `"Hello World"`
* LTrimSpace: Trim leading spaces. `"  hello world  "` => `"hello world  "`
* RTrimSpace: Trim trailing spaces. `"  hello world  "` => `"  hello world"`
* ToEmail: Lowercase email domain. `"John@EXAMPLE.COM"` => `"John@example.com"`
* ToAlphaNumeric: Remove non-alphanumeric characters. `"hello world123"` => `"helloworld123"`
* ToAlpha: Remove non-alphanumeric characters and spaces. `"hello world123"` => `"helloworld"`
* ToNumeric: Remove non-numeric characters. `"hello world123"` => `"123"`
* StripHTML: Remove HTML tags. `"hello <b>world</b>"` => `"hello world"`
* ToName: Makes a string safe to use in a file name. `"hello world"` => `"hello-world"`
* ToPath: Makes a string safe to use as an url path. `"hello world"` => `"hello-world"`
* StripAccents: Remove accents. `"héllö wórld"` => `"hello world"`