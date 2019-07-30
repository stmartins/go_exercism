package acronym

import "strings"

func Abbreviate(s string) string {
	s = strings.Replace(s, "-", " ", -1)
	s = strings.Replace(s, "_", " ", -1)
	result := strings.Split(s, " ")
	var str []string
	for _, data := range result {
		if data != "" {
			str = append(str, strings.ToUpper(string(data[0])))
		}
	}
	acro := strings.Join(str, "")
	return acro
}
