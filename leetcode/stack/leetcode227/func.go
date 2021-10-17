package leetcode227

import (
	"strconv"
)

/*
基本计算器 II
给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
整数除法仅保留整数部分。
输入：s = "3+2*2"
输出：7
*/

/*
基本思路：中序表达式，存在表达式优先级
维护一个操作数栈 和数据栈
将中序表达式转化为后缀表达式方式计算，比如示例



操作数级别高的，需要置于栈顶
操作数同级的，将操作数取出，进行计算，将计算值压栈
操作数级别
暂不考虑
* /
+ -
*/

/*
1 + 2 - 3 * 4 / 5
---> 转逆序表达式 12+34*5/-
opts: +
tokens: 1 2

-号入栈，发现-号优先级<= 栈顶，栈顶出栈，-号入栈
opts: -
tockens: 1 2 +

* 号优先级高于栈顶，直接入栈
opts: + *
tockens: 1 2 + 3 4

除号优先级<=栈顶元素*，栈顶出栈入tockens，/号入栈
*/

/*
1 + 2 - 3 * 4 - (5 + 6 * 7)
opts: +
vals: 1 2

opts: -
vals 1 2 + 3

opts: - -
vals 1 2 + 3 4 *

opts: - - (
vals: 1 2 + 3 4 *

opts: - - + (
vals: 1 2 + 3 4 * 5 6

opts: - - + *(
vals:1 2 + 3 4 * 5 6 7

opts: - - + *
vals: 1 2 + 3 4 * 5 6 7

弹栈
vals: 1 2 + 3 4 * 5 6 7 * + - -
1+2 - 3*4 - (5 + 6*7)
*/

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
	for i, n := 0, len(s); i < n; i++ {
		// 数字， 加到bytes里
		if s[i] >= '0' && s[i] <= '9' {
			bytes = append(bytes, s[i])
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
		// 如果栈顶大于当前元素
		curOpt := s[i : i+1]
		for !opts.isEmpty() && getPriority(opts.peek()) >= getPriority(curOpt) {
			tockens.push(opts.pop())
		}
		opts.push(curOpt)
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
