package utils

import (
	"encoding/json"
	"strings"
)

// is the string empty
func IsEmpty(data string) bool {
	if (data == "") || (len(data) == 0) {
		return true
	}
	return false
}

func Map2Str(tmp map[string]string) string {
	str, err := json.Marshal(tmp)
	if err != nil {
		panic("Failed to parse map to string")
	} else {
		return string(str)
	}
}

func Map2StrWithEqu(tmp map[string]string) string {
	var builder strings.Builder
	for k, v := range tmp {
		builder.WriteString(k)
		builder.WriteString("=")
		builder.WriteString(v)
		builder.WriteString(",")
	}
	res := builder.String()
	res = strings.TrimRight(res, ",")
	return res
}
