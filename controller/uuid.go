package controller

import "regexp"

func isUUIDv4(input string) bool {
	// Expressão regular para UUID v4
	var uuidv4Regex = regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)
	return uuidv4Regex.MatchString(input)
}
