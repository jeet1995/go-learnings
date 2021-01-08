package main

import (
	"fmt"
	"reflect"
)

type AnimalWithTag struct {
	Name   string // `required max:"100"`
	Origin string
}

func main() {
	t := reflect.TypeOf(AnimalWithTag{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
