package luhn

import c "strconv"
import s "strings"

func ValidInput(input string) bool {
	for i := 0; i < len(input); i++ {
		if !(input[i] >= '0' && input[i] <= '9') {
			return false
		}
	}
	return true
}

func isModTen(tot int) bool {
	if tot%10 == 0 {
		return true
	}
	return false
}

func Valid(input string) bool {

	str := s.Replace(input, " ", "", -1)
	len := len(str)

	if ValidInput(str) == false || len == 0 || len == 1 {
		return false
	}

	total := 0
	index := 0
	pair := 1

	if len%2 == 1 {
		index = 1
		pair = -1
	}

	for ; index < len; index += 2 {
		numb, _ := c.Atoi(string(str[index]))
		numb *= 2
		if numb > 9 {
			numb -= 9
		}
		total += numb
		num, _ := c.Atoi(string(str[index+pair]))
		total += num
	}
	if pair == -1 {
		num, _ := c.Atoi(string(str[index+pair]))
		total += num
	}
	return isModTen(total)
}
