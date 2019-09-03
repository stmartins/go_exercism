package sublist

type Relation string

func Sublist(listOne []int, listTwo []int) Relation {
	var rel Relation
	if len(listOne) == 0 && len(listTwo) == 0 {
		rel = "equal"
	} else if len(listOne) < len(listTwo) {
		rel = "sublist"
	} else if len(listOne) > len(listTwo) {
		rel = "superlist"
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
