package raindrops

import "strconv"

func fact3(input int) bool {
	if input%3 == 0 {
		return true
	}
	return false
}

func fact5(input int) bool {
	if input%5 == 0 {
		return true
	}
	return false
}

func fact7(input int) bool {
	if input%7 == 0 {
		return true
	}
	return false
}

func Convert(input int) string {

	var str string

	if fact3(input) {
		str = "Pling"
	}
	if fact5(input) {
		str += "Plang"
	}
	if fact7(input) {
		str += "Plong"
	}
	if str == "" {
		str = strconv.Itoa(input)
	}
	return str
}
