package leetcode224

func calculate(s string) int {
	stack := [][]int{}
	num := 0
	i := 0
	t := 1
	for i < len(s) {

		if '0' <= s[i] && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')
			i++
		} else if s[i] == '+' {
			t = 1
			stack = append(stack, []int{num, t})
			num = 0
			i++
			for i < len(s) {
				if s[i] == ' ' {
					i++
					continue
				}
				if s[i] < '0' || s[i] > '9' {
					break
				}
				num = num*10 + int(s[i]-'0')
				i++
			}
			num += stack[len(stack)-1][0]
			stack = stack[:len(stack)-1]
		} else if s[i] == '-' {
			t = -1
			stack = append(stack, []int{num, t})
			num = 0
			i++
			for i < len(s) {
				if s[i] < '0' || s[i] > '9' {
					break
				}
				num = num*10 + int(s[i]-'0')
				i++
			}
			num = stack[len(stack)-1][0] - num
			stack = stack[:len(stack)-1]
		} else if s[i] == '(' {
			stack = append(stack, []int{num, t})
			num = 0
			i++
		} else if s[i] == ')' {
			num = stack[len(stack)-1][0] + num*stack[len(stack)-1][1]
			stack = stack[:len(stack)-1]
			i++
		} else {
			i++
		}
	}
	return num
}
