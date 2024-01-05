package utils

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strconv"
	"strings"
	"unicode"
)

/*
Use to convert student score string to float32
If the string is not a number, it will return 0.0
*/
func ProcessFloatScore(value string) float32 {
	floatVal, err := strconv.ParseFloat(value, 32)

	if err != nil {
		return 0.0
	}

	return float32(floatVal)
}

func CheckStudentClassColIsReal(value string) bool {
	if strings.HasPrefix(value, "CT") || strings.HasPrefix(value, "AT") || strings.HasPrefix(value, "DT") {
		return true
	} else {
		return false
	}
}

func ProcessStudentName(name string) string {
	parts := strings.Fields(name)

	/*
		This function uses the strings.Fields function to split the name into parts.
		It then iterates over the parts and checks if two consecutive parts are capitalized.
		If they are, it inserts a space between them and updates the parts slice.
	*/
	for i := 0; i < len(parts)-1; i++ {
		if unicode.IsUpper([]rune(parts[i])[0]) && unicode.IsUpper([]rune(parts[i+1])[0]) {
			parts[i] = parts[i] + " " + parts[i+1]
			parts = append(parts[:i+1], parts[i+2:]...)
		}
	}

	// Init format engine
	caser := cases.Title(language.Vietnamese)

	// Join parts and convert to Vietnamese's name format (capitalize first letter)
	return caser.String(strings.Join(parts, " "))
}

func CleanStudentStringScore(value string) string {
	return strings.ReplaceAll(value, "\u00a0", "")
}
