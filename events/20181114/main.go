package main

import (
	"fmt"

	"gopkg.in/go-playground/validator.v8"
)

// Item -
type Item struct {
	Name   string `validate:"required"`
	Number int    `validate:"gt=100"`
}

func main() {

	v := validator.New(&validator.Config{TagName: "validate"})

	items := []Item{
		{Name: "A", Number: 200},
		{Name: "", Number: 200},
		{Name: "C", Number: 99},
	}

	for i := range items {
		if err := v.Struct(&items[i]); err != nil {
			fmt.Println(err)
		}
	}
}
