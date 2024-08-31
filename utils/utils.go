package utils

// a-z A-Z _
func IsLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// we only support int number for now. (12121)
func IsDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
