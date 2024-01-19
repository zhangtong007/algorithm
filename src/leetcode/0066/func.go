package leetcode66

func plusOne(digits []int) []int {
	n := len(digits)
	bit := 1
	for i := n - 1; i >= 0; i-- {
		digits[i] = digits[i] + bit
		if digits[i] == 10 {
			bit = 1
			digits[i] = 0
			continue
		}
		bit = 0
	}
	if bit == 1 {
		digits = append(digits, 0)
		copy(digits[1:], digits)
		digits[0] = 1
		return digits
	}
	return digits
}
