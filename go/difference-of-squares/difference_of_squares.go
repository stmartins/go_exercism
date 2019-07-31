package diffsquares

func Difference(n int) int {

	sqOfSum := SquareOfSum(n)
	sumOfSq := SumOfSquares(n)

	dif := sqOfSum - sumOfSq
	return dif
}

func SquareOfSum(n int) int {

	var sqOfSum int

	for i := 1; i <= n; i++ {
		sqOfSum += i
	}
	sqOfSum = sqOfSum * sqOfSum
	return sqOfSum
}

func SumOfSquares(n int) int {

	var sumOfSq int

	for i := 1; i <= n; i++ {
		sumOfSq += i * i
	}
	return sumOfSq
}
