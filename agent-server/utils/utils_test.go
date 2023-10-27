package utils

import (
	"fmt"
	"testing"
)

func TestMap2Str(t *testing.T) {
	m := make(map[string]string)
	m["hello"] = "world"
	m["k1"] = "v1"

	str := Map2Str(m)
	fmt.Println(str)
	fmt.Println(m)
}

func TestMap2StrWithEqu(t *testing.T) {
	m := make(map[string]string)
	m["hello"] = "world"
	m["k1"] = "v1"

	str := Map2StrWithEqu(m)
	fmt.Println(str)
}
