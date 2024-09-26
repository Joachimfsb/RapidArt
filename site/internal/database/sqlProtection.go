package database

import (
	"fmt"
	"strings"
)

func SQLCheck(input string) error {
	valid := false
	for i := 0; i < len(input); i++ {
		if strings.Contains(input, "WHERE") == true || strings.Contains(input, "where") == true ||
			strings.Contains(input, "FROM") == true || strings.Contains(input, "from") == true ||
			strings.Contains(input, "DROP") || strings.Contains(input, "drop") == true || string(input[i]) == "=" ||
			strings.Contains(input, "INSERT") || strings.Contains(input, "insert") ||
			string(input[i]) == "$" || string(input[i]) == "*" || string(input[i]) == "'" || string(input[i]) == "+" ||
			string(input[i]) == ";" || string(input[i]) == "`" {
			valid = false
			break
		}
		valid = true
	}
	if valid == true {
		return nil
	}
	return fmt.Errorf("not valid input")
}
