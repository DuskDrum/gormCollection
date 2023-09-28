package test

import (
	"fmt"
	"testing"
)

const (
	SunDay int = iota
	Monday
	_
	Wednesday
	Thursday
	Friday
	Saturday
)

func TestIota(t *testing.T) {
	fmt.Println("SunDay： \t", SunDay)
	fmt.Println("Monday： \t", Monday)
	fmt.Println("Wednesday： \t", Wednesday)
	fmt.Println("Thursday： \t", Thursday)
	fmt.Println("Friday： \t", Friday)
	fmt.Println("Saturday： \t", Saturday)
}
