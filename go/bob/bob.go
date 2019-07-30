package bob

func isUpperCase(c byte) bool {
	if c >= 'A' && c <= 'Z' {
		return true
	}
	return false
}

func isPonctuation(c byte) bool {

	if c == '!' || c == '?' || c == '.' || c == ',' || c == '\'' || c == '-' {
		return true
	}
	return false
}

func isNumber(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func isLetter(c byte) bool {
	if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
		return true
	}
	return false
}

func isSpecialChar(b byte) bool {
	special := "@#$%^&*()_+=[]{}\\\"|;:./<>"
	for i := 0; i < len(special); i++ {
		if b == special[i] {
			return true
		}
	}
	return false
}

func isYelling(str string) bool {
	isup := true
	in := false

	for i := 0; i < len(str); i++ {
		if isLetter(str[i]) && !isSpecialChar(str[i]) && str[i] != ' ' && !isPonctuation(str[i]) && !isNumber(str[i]) {
			in = true
			isup = isUpperCase(str[i])
			if isup == false {
				break
			}
		}
	}
	if isup == true && in == true {
		for i := 0; i < len(str); i++ {
			if !isPonctuation(str[i]) && isNumber(str[i]) {
				isup = true
			}
		}
	} else {
		isup = false
	}
	return isup
}

func isQuestion(str string) bool {
	isquest := false
	if str[len(str)-1] == '?' {
		isquest = true
	} else {
		for i := 0; i < len(str); i++ {
			if str[i] == '?' {
				i++
				isquest = true
				for i < len(str) {
					if str[i] != ' ' {
						isquest = false
					}
					i++
				}
			}
		}
	}
	return isquest
}

func isYellingQuestion(str string) bool {
	if isYelling(str) && isQuestion(str) {
		return true
	}
	return false
}

func sayNothing(str string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' && str[i] != '\t' && str[i] != '\n' && str[i] != '\r' {
			return false
		}
	}
	return true
}

func Hey(remark string) string {
	var answer string
	switch {
	case sayNothing(remark):
		answer = "Fine. Be that way!"
	case isYellingQuestion(remark):
		answer = "Calm down, I know what I'm doing!"
	case isQuestion(remark):
		answer = "Sure."
	case isYelling(remark):
		answer = "Whoa, chill out!"
	default:
		answer = "Whatever."
	}
	return answer
}

// best solution
package bob

import "strings"

var responses = map[string]string{
	"shoutingQuestion": "Calm down, I know what I'm doing!",
	"question":         "Sure.",
	"shouting":         "Whoa, chill out!",
	"silence":          "Fine. Be that way!",
	"default":          "Whatever.",
}

// Hey returns Bobs answer to the user input
func Hey(remark string) string {
	switch remark = strings.TrimSpace(remark); {
	case silent(remark):
		return responses["silence"]
	case yelling(remark):
		if asking(remark) {
			return responses["shoutingQuestion"]
		}
		return responses["shouting"]
	case asking(remark):
		return responses["question"]
	default:
		return responses["default"]
	}
}

func yelling(remark string) bool {
	return strings.ToUpper(remark) == remark && strings.ToLower(remark) != strings.ToUpper(remark)
}

func asking(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func silent(remark string) bool {
	return remark == ""
}
