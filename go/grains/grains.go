package grains

import "errors"

func Square(input int) (uint64, error) {
	if input < 1 || input > 64 {
		return 0, errors.New("error, input have to be between 1 and 64")
	}
	square := uint64(1)
	for i := 1; i < input; i++ {
		square *= 2
	}
	return square, nil
}

func Total() uint64 {
	var total uint64

	for i := 0; i <= 64; i++ {
		square, _ := Square(i)
		total += square
	}
	return total
}
