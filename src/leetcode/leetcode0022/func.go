package leetcode22

/*
使用分治计算：
分治思想：不重，不漏

n对括号拼接字符串，将字符串进行拆分，划分为子问题：

S ==> (A)B

将原来的字符串S 划分为字符串A和字符串B 其中A再一对括号内
这样能保证(A)B 已经是不能再分割的整体

    S    ==> (A)B
    n对      k对 n-k-1对  S总计有n对括号，那么 设A为k对括号，则B为 n-k-1对括号
枚举K： 0-n-1
如样例 n=3
            k = 0，则：    () | (())  or  () | ()()  A为空， B为 (())  or ()()
            k = 1, 则:     (()) | ()                A为()，B为()
            k = 2, 则：    ((())) |   or  (()())    A为()() or (())  B为空
*/

var results = map[int][]string{}

func generateParenthesis(n int) []string {
	// 递归边界
	if n == 0 {
		return []string{""}
	}
	// 记忆化
	if ans, exist := results[n]; exist {
		return ans
	}
	ans := []string{}
	// (A)B
	for k := 0; k < n; k++ {
		// 枚举0到n-1
		A := generateParenthesis(k)
		B := generateParenthesis(n - k - 1)
		// 将a和b的组合进行排列重组
		for _, aStr := range A {
			for _, bStr := range B {
				ans = append(ans, "("+aStr+")"+bStr)
			}
		}
	}
	results[n] = append([]string(nil), ans...)
	return ans
}
