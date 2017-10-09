package wedge

func Div(dividend int, divisor int) (quotient int, remainder int) {
	quotient = dividend/divisor
	remainder = dividend - quotient*divisor
	return
}

func NSplit(amount int, n_parts int) (most_parts int, lucky_part int) {
	if n_parts <= 0 {
		panic("n_parts must be greater than zero")
	}
	most_parts, lucky_part = Div(amount, n_parts)
	lucky_part += most_parts
	return
}