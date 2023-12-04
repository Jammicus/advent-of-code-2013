package helpers

func StringToArray(input string) []rune {
	runes := make([]rune, len(input))

	for j, char := range input {

		runes[j] = char

	}

	return runes
}