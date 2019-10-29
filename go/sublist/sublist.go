package sublist

type Relation string

func checkSubList(listOne []int, listTwo []int) Relation {
	var rel Relation
	rel = "sublist"
	for i := 0; i < len(listOne); i++ {
		if listOne[i] == listTwo[0] {
			for j := 0; j < len(listTwo); j++ {
				if i + j < len(listOne) {
					if listOne[i + j] == listTwo[j] {
						rel = "sublist"
					} else {
						rel = "unequal"
					}
				}
			}
		}
	}
	return rel
}

func checkSuperList(listOne []int, listTwo []int) Relation {
	var rel Relation
	rel = "superlist"
	for i := 0; i < len(listTwo); i++ {
		if listTwo[i] == listOne[0] {
			for j := 0; j < len(listOne); j++ {
				if i + j < len(listTwo) {
					if listTwo[i + j] == listOne[j] {
						rel = "superlist"
					} else {
						rel = "unequal"
					}
				} else if j < 2 {
					rel = "unequal"
				}
			}
		}
	}
	return rel
}

func Sublist(listOne []int, listTwo []int) Relation {
	var rel Relation
	if len(listOne) == 0 && len(listTwo) == 0 {
		rel = "equal"
	} else if len(listOne) < len(listTwo) {
		rel = checkSubList(listOne, listTwo)
	} else if len(listOne) > len(listTwo) {
		rel = checkSuperList(listOne, listTwo)
	} else if len(listOne) == len(listTwo) {
		for i := 0; i < len(listOne); i++ {
			if listOne[i] != listTwo[i] {
				rel = "unequal"
				break
			}
			rel = "equal"
		}
	}
	return rel
}
