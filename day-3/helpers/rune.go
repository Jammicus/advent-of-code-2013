package helpers

import (
	"fmt"
	"strconv"
)

func ExtractWordValueFromCoords(matrix []rune) int {

	str := ""
	for _, v := range matrix {
		fmt.Println(string(v))
		str = fmt.Sprintf("%v%v", str, string(v))
	}

	fmt.Println(str)

	val, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println(err)
		return -100000
	}

	fmt.Println("Found: ", val)

	return val
}
