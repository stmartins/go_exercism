package scrabble

import "strings"

func Score(input string) int {
	points := 0

	if input == "" {
		return points
	}

	letters := map[int]string{
		1:  "AEIOULNRST",
		2:  "DG",
		3:  "BCMP",
		4:  "FHVWY",
		5:  "K",
		8:  "JX",
		10: "QZ",
	}

	for _, inputLetter := range input {
		for letterPoint, comp := range letters {
			if strings.Contains(comp, strings.ToUpper(string(inputLetter))) {
				points += letterPoint
			}
		}
	}
	return points
}
