package internal

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Add(numbers string) (result int, err error) {
	delimiter := findDelimiter(numbers)
	numbers = strings.Replace(numbers, "\n", delimiter, -1)
	splitNumbers := strings.Split(numbers, delimiter)
	sum := 0
	for _, i := range splitNumbers {
		number, _ := strconv.Atoi(i)
		if number < 0 {
			msg := fmt.Sprintf("negatives not allowed: %s", numbers)
			err := errors.New(msg)
			return 0, err
		}
		sum += number
	}
	return sum, nil
}

func findDelimiter(text string) (delimiter string) {
	if strings.HasPrefix(text, "//") && len(text) >= 3 {
		return string(text[2])
	}

	return ","
}
