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

func TestValidateIP(t *testing.T) {
	ip := ""
	res := ValidateIP(ip)
	fmt.Println(res)

	ip = "1.2"
	res = ValidateIP(ip)
	fmt.Println(res)

	ip = "1.2.3.4"
	res = ValidateIP(ip)
	fmt.Println(res)

	ip = "127.0.0.1"
	res = ValidateIP(ip)
	fmt.Println(res)
}

func TestFormatDateNow(t *testing.T) {
	res := FormatDateNow()
	println(res)
}
