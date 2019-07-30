package isogram

import s "strings"

func isLetter(b1 rune, b2 rune) bool {
	if b1 >= 'A' && b1 <= 'Z' && b2 >= 'A' && b2 <= 'Z' {
		return true
	}
	return false
}

func IsIsogram(input string) bool {

	if input == "" {
		return true
	}

	str := s.ToUpper(input)

	for i1, c1 := range str {
		for i2, c2 := range str {
			if i2 > i1 && isLetter(c1, c2) && c1 == c2 {
				return false
			}
		}
	}
	return true
}
