package math

func adder(xs ...int) int {
	res := 0
	for _, v := range xs {
		res += v
	}
	return res
}
