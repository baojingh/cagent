package utils

// is the string empty
func IsEmpty(data string) bool {
	if (data == "") || (len(data) == 0) {
		return true
	}
	return false
}
