package leetcode2201

// 哈希表 暴力破解
// 存储挖掘点，遍历部件，从左上端点至右下端点是否均在dig中，均在，则可以掘出
func digArtifacts(n int, artifacts [][]int, dig [][]int) int {
	exist := make(map[point]bool, len(dig))
	for _, p := range dig {
		exist[point{x: p[0], y: p[1]}] = true
	}
	ans := 0
	// 遍历部件
z:
	for _, arartifact := range artifacts {
		for i := arartifact[0]; i <= arartifact[2]; i++ {
			for j := arartifact[1]; j <= arartifact[3]; j++ {
				if !exist[point{x: i, y: j}] {
					continue z
				}
			}
		}
		ans++
	}
	return ans
}

type point struct {
	x, y int
}
