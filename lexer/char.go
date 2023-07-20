package lexer

// isIdentChar checks if the provided byte is a valid character for an identifier
func isIdentChar(char byte) bool {
	return char >= 'a' && char <= 'z' ||
		char >= 'A' && char <= 'Z' ||
		char == '_'
}

// isDigit checks if the byte is a numeric character
func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}
