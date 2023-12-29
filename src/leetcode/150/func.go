package leetcode150

import (
	"fmt"
	"strconv"
)

/*
 逆波兰表达式求值
 tokens = ["2","1","+","3","*"]
 9
*/

type Stack struct {
	vals []int
	top  int
}

func evalRPN1(tokens []string) int {
	// stack 遇到操作数opt，则将栈顶两元素出栈，操作之后入栈
	fmt.Printf("%v", tokens)
	stack := &Stack{
		vals: make([]int, len(tokens)),
		top:  -1,
	}
	var isOpt bool
	for _, val := range tokens {
		// 题目保证输入合法, 如果不是操作数，入栈
		_, isOpt = opts[val]
		if isOpt {
			stack.push(val)
			continue
		}
		// 否则，将栈顶两个元素出栈，进行计算 然后将计算后的值入栈
		stack.processOpt(val)
	}
	// 站内最终只剩计算结果
	return stack.vals[stack.top]
}

func (s *Stack) push(val string) {
	v, _ := strconv.Atoi(val)
	s.top++
	s.vals[s.top] = v
}

func (s *Stack) processOpt(opt string) {
	val := opts[opt](s.vals[s.top-1], s.vals[s.top])
	s.top--
	// 新值入栈
	s.vals[s.top] = val
}

// 方法2： 使用switch判断，不借助其他
func evalRPN2(tokens []string) int {
	stack := &Stack{
		vals: make([]int, len(tokens)),
		top:  -1,
	}
	for _, val := range tokens {
		// 不是操作数，则数据压栈
		if !isOpt(val) {
			stack.push(val)
			continue
		}
		// 否则，处理计算
		stack.calculate(val)
	}
	return stack.vals[stack.top]
}

func (s *Stack) calculate(opt string) {
	var (
		newVal     int
		num1, num2 = s.vals[s.top-1], s.vals[s.top]
	)
	s.top--
	switch opt {
	case "+":
		newVal = num1 + num2
	case "-":
		newVal = num1 - num2
	case "*":
		newVal = num1 * num2
	case "/":
		newVal = num1 / num2
	}
	s.vals[s.top] = newVal
}

var opts = map[string]func(a, b int) int{
	"+": func(a, b int) int {
		return a + b
	},
	"-": func(a, b int) int {
		return a - b
	},
	"*": func(a, b int) int {
		return a * b
	},
	"/": func(a, b int) int {
		return a / b
	},
}

// 提交
func evalRPN(tokens []string) int {
	stack := make([]string, 0, len(tokens))
	var (
		num1 int
		num2 int
	)
	for i := range tokens {
		if !isOpt(tokens[i]) {
			stack = append(stack, tokens[i])
			continue
		}
		num2, _ = strconv.Atoi(stack[len(stack)-1])
		stack = stack[:len(stack)-1]
		num1, _ = strconv.Atoi(stack[len(stack)-1])
		stack = stack[:len(stack)-1]
		stack = append(stack, strconv.Itoa(cal(num1, num2, tokens[i])))
	}
	ret, _ := strconv.Atoi(stack[0])
	return ret
}

func cal(num1, num2 int, opt string) int {
	return opts[opt](num1, num2)
}

func isOpt(val string) bool {
	return val == "+" || val == "-" || val == "*" || val == "/"
}
