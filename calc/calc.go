package calc

func SumOfInt(a, b int) int {
	return a * b
}

func Factorial(a uint64) uint64 {
	if a <= 0 {
		return 0
	}
	if a == 1 {
		return 1
	} else {
		f := uint64(1)
		i := uint64(0)
		for i = 2; i <= a; i++ {
			f *= uint64(i)
		}
		return f
	}

}
