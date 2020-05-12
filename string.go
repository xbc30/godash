package godash

import (
	"strings"
)

/*
 *@Method: StringCapitalize
 *@Description: Converts the first character of string to upper case and the remaining to lower case.
 *@Param: 1.input string
 *@Return: string
 */
func StringCapitalize(input string) string {
	var output string
	for index, item := range input {
		if index == 0 {
			output += strings.ToUpper(string(item))
		} else {
			output += strings.ToLower(string(item))
		}
	}
	return output
}

/*
 *@Method: StringEndsWith
 *@Description: Checks if string ends with the given target string.
 *@Param: 1.input string 2.target string
 *@Return: bool
 */
func StringEndsWith(input, suffix string, end int) bool {
	newString := string([]byte(input)[:len(input)-end])
	return strings.HasSuffix(newString, suffix)
}
