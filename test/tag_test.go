package test

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

type TestTagS struct {
	ID   uint   `notNull:"true"`
	Name string `notNull:"true"`
	Age  int
}

func TestTag(t *testing.T) {
	testTagS := TestTagS{
		ID:   0,
		Name: "John",
		Age:  0,
	}

	if err := validate(testTagS); err != nil {
		fmt.Println("Validation error:", err)
		return
	}

	fmt.Println("Validation passed!")
}

func validate(obj interface{}) error {
	value := reflect.ValueOf(obj)
	typ := value.Type()

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		tag := typ.Field(i).Tag.Get("notNull")
		if parseBool, err2 := strconv.ParseBool(tag); err2 == nil {
			if parseBool {
				// 0值
				zero := reflect.Zero(field.Type())
				if reflect.DeepEqual(field.Interface(), zero.Interface()) {
					fieldName := typ.Field(i).Name
					return fmt.Errorf("field %s cannot be empty", strings.ToLower(fieldName))
				}
			}
		}

	}

	return nil
}
