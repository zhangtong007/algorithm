package leetcode166

import (
	"strconv"
)

// 两个数相除，要么整除，要么无限循环小数
func fractionToDecimal(numerator int, denominator int) string {
	// 1. 判断能否整除，若能整除，直接返回
	if numerator%denominator == 0 {
		return strconv.Itoa(numerator / denominator)
	}
	s := []byte{}
	// 2. 判断是否异号，决定是否添加负号
	if (numerator > 0) != (denominator > 0) {
		s = append(s, '-')
	}
	// 3. 处理整数部分
	// 置为正数
	numerator = abs(numerator)
	denominator = abs(denominator)
	integerPart := numerator / denominator
	s = append(s, strconv.Itoa(integerPart)...)
	s = append(s, '.')
	// 4. 处理小数部分
	// 存储除数的循环索引部分
	indexMap := map[int]int{}
	remainder := numerator % denominator
	for remainder != 0 && indexMap[remainder] == 0 {
		indexMap[remainder] = len(s)
		remainder *= 10
		s = append(s, '0'+byte(remainder/denominator))
		remainder %= denominator
	}
	// 有循环节
	if remainder > 0 {
		indexInsert := indexMap[remainder]
		s = append(s[:indexInsert], append([]byte{'('}, s[indexInsert:]...)...)
		s = append(s, ')')
	}
	return string(s)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
