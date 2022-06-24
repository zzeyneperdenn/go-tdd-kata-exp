package internal

import (
	"strconv"
	"strings"
)

func Add(numbers string) (result int, err error) {
	numbers = strings.Replace(numbers, "\n", ",", -1)
	splitNumbers := strings.Split(numbers, ",")
	sum := 0
	for _, i := range splitNumbers {
		number, _ := strconv.Atoi(i)
		sum += number
	}
	return sum, nil
}
