package leetcode224

import (
	"strconv"
)

type Stack struct {
	vals []string
	top  int
}

func (s *Stack) push(val string) {
	s.top++
	s.vals[s.top] = val
}

func (s *Stack) pop() string {
	val := s.vals[s.top]
	s.top--
	return val
}

func (s *Stack) peek() string {
	return s.vals[s.top]
}

func (s *Stack) isEmpty() bool {
	return s.top == -1
}

// 逆波兰表达式求值方法
func evalRPN(tokens []string) int {
	stack := &Stack{
		vals: make([]string, len(tokens)),
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
	v, _ := strconv.Atoi(stack.pop())
	return v
}

func isOpt(val string) bool {
	return val == "+" || val == "-" || val == "*" || val == "/"
}

func (s *Stack) calculate(opt string) {
	var newVal int
	num2, _ := strconv.Atoi(s.pop())
	num1, _ := strconv.Atoi(s.pop())
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
	s.push(strconv.Itoa(newVal))
}

// ----------以上是逆波兰式计算方式---------

func calculate(s string) int {
	s = s + " "
	strs := process(s)
	return evalRPN(strs)
}

// 转化为逆波兰式数组
func process(s string) []string {
	opts := &Stack{
		vals: make([]string, len(s)),
		top:  -1,
	}
	tockens := &Stack{
		vals: make([]string, len(s)),
		top:  -1,
	}
	// 用于接收数字
	bytes := make([]byte, 0)
	needZero := true
	for i, n := 0, len(s); i < n; i++ {
		// 数字， 加到bytes里
		if s[i] >= '0' && s[i] <= '9' {
			bytes = append(bytes, s[i])
			needZero = false
			continue
		} else {
			// 如果num不为空 先将数字入栈，重新创建bytes
			if len(bytes) != 0 {
				tockens.push(string(bytes))
				bytes = make([]byte, 0)
			}
		}
		// 空格不做处理
		if s[i] == ' ' {
			continue
		}
		// 左括号，直接入栈
		if s[i] == '(' {
			opts.push(s[i : i+1])
			needZero = true
		}
		if s[i] == ')' {
			needZero = false
			opt := opts.pop()
			for opt != "(" {
				tockens.push(opt)
			}
		}
		if (s[i] == '+' || s[i] == '-') && needZero {
			tockens.push("0")
		}
		// 如果栈顶大于当前元素
		curOpt := s[i : i+1]
		for !opts.isEmpty() && getPriority(opts.peek()) >= getPriority(curOpt) {
			tockens.push(opts.pop())
		}
		opts.push(curOpt)
		// 初始化
		needZero = true
	}
	for opts.top != -1 {
		tockens.push(opts.pop())
	}
	return tockens.vals[:tockens.top+1]
}

func getPriority(ch string) int {
	if ch == "+" || ch == "-" {
		return 0
	}
	if ch == "*" || ch == "/" {
		return 1
	}
	// 括号
	return 2
}
