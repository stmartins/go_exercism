package triangle

import "math"

type Kind int

const (
    NaT = iota
    Equ
    Iso
    Sca
)

func KindFromSides(a, b, c float64) Kind {

    if (a <= 0 || b <= 0 || c <= 0) {
        return NaT
    } else if (a + b < c || a + c < b || b + c < a) {
        return NaT
    } else if (math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c)) {
        return NaT
    } else if (math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0)) {
        return NaT
    }
    if (a == b && b == c) {
        return Equ
    } else if (a == b || a == c || b == c) {
        return Iso
    } else if (a != b && a != c && b != c) {
        return Sca
    }
	return NaT
}
