package hamming

import err "errors"

func Distance(a, b string) (int, error) {

	dif := 0

	if len(a) != len(b) {
		return 0, err.New("true")
	}
	if a != b {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				dif++
			}
		}
		return dif, nil
	}
	return 0, nil
}
