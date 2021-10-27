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

// ----------以上是逆波兰式计算方式---------

func calculate(s string) int {
	s = s + " "
	strs := process(s)
	return evalRPN(strs)
}

func process(s string) []string {
	s += ""
	number := make([]byte, 0)
	tokens := make([]string, 0)
	opts := make([]byte, 0)
	needZero := true
	for i := range s {
		if s[i] >= '0' && s[i] <= '9' {
			number = append(number, s[i])
			needZero = false
			continue
		} else {
			if len(number) != 0 {
				tokens = append(tokens, string(number))
				number = make([]byte, 0)
			}
		}
		if s[i] == ' ' {
			continue
		}
		if s[i] == '(' {
			needZero = true
			opts = append(opts, '(')
			continue
		}
		if s[i] == ')' {
			needZero = false
			for opts[len(opts)-1] != '(' {
				tokens = append(tokens, string(opts[len(opts)-1]))
				opts = opts[:len(opts)-1]
			}
		}
		if (s[i] == '+' || s[i] == '-') && needZero {
			tokens = append(tokens, "0")
		}
		curRank := getRank(s[i])
		for (len(opts) > 0) && getRank(opts[len(opts)-1]) >= curRank {
			tokens = append(tokens, string(opts[len(opts)-1]))
			opts = opts[:len(opts)-1]
		}
		opts = append(opts, s[i])
		needZero = true
	}
	for len(opts) != 0 {
		tokens = append(tokens, string(opts[len(opts)-1]))
		opts = opts[:len(opts)-1]
	}
	return tokens
}

func getRank(i byte) int {
	if i == '+' || i == '-' {
		return 1
	}
	if i == '*' || i == '/' {
		return 2
	}
	return 3
}
