package leetcode2834

/*
给你两个正整数：n 和 target 。

如果数组 nums 满足下述条件，则称其为 美丽数组 。

nums.length == n.
nums 由两两互不相同的正整数组成。
在范围 [0, n-1] 内，不存在 两个 不同 下标 i 和 j ，使得 nums[i] + nums[j] == target 。
返回符合条件的美丽数组所可能具备的 最小 和，并对结果进行取模 109 + 7。
*/
/*
贪心：
从1开始递增枚举值i，满足target-i 没有被选取到

易得：
若i >= target, 均可取
若i < target:
	1, 则target-1不可取
	2, 则target-2不可取
so: [1-target) 前一半可取，后一半不可取
*/

const mod = 1e9 + 7

func minimumPossibleSum(n int, target int) int {
	if n == 1 {
		return 1
	}
	halfTarget := target / 2
	// 分一半之后长度>=n 求和1-n即可
	if halfTarget >= n {
		return sum(1, n) % mod
	}
	ans := sum(1, halfTarget)
	n -= halfTarget
	ans += sum(target, target+n-1)
	return ans % mod
}

// 求a-b的和
func sum(a, b int) int {
	return (a + b) * (b - a + 1) / 2
}
