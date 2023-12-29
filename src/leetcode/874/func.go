package leetcode874

/*
// Key 使用 string方式，效率较差
import "strconv"

func robotSim(commands []int, obstacles [][]int) int {
	// 用来存储障碍物位置
	obstaclesMap := make(map[string]bool)
	for _, val := range obstacles {
		key := strconv.Itoa(val[0]) + "," + strconv.Itoa(val[1])
		obstaclesMap[key] = true
	}
	ans := 0
	// 0, 1, 2, 3 分别对应 北，东，南，西(上右下左)
	direct := 0
	// 对应四个方向的移动
	moveDirect := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	curPosition := []int{0, 0}
	for _, v := range commands {
		if v == -1 {
			direct = (direct + 1) % 4
			continue
		}
		if v == -2 {
			direct = (direct + 3) % 4
			continue
		}
		for i := 0; i < v; i++ {
			x, y := curPosition[0]+moveDirect[direct][0], curPosition[1]+moveDirect[direct][1]
			key := strconv.Itoa(x) + "," + strconv.Itoa(y)
			if obstaclesMap[key] {
				break
			}
			curPosition[0], curPosition[1] = x, y
			ans = max(ans, x*x+y*y)
		}
	}
	return ans
}
*/

func robotSim(commands []int, obstacles [][]int) int {
	// 用来存储障碍物位置
	obstaclesMap := make(map[int]bool)
	for _, val := range obstacles {
		key := calHash(val[0], val[1])
		obstaclesMap[key] = true
	}
	ans := 0
	// 0, 1, 2, 3 分别对应 北，东，南，西(上右下左)
	direct := 0
	// 对应四个方向的移动
	moveDirect := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	curPosition := []int{0, 0}
	for _, v := range commands {
		if v == -1 {
			direct = (direct + 1) % 4
			continue
		}
		if v == -2 {
			direct = (direct + 3) % 4
			continue
		}
		for i := 0; i < v; i++ {
			x, y := curPosition[0]+moveDirect[direct][0], curPosition[1]+moveDirect[direct][1]
			key := calHash(x, y)
			if obstaclesMap[key] {
				break
			}
			curPosition[0], curPosition[1] = x, y
			ans = max(ans, x*x+y*y)
		}
	}
	return ans
}

func calHash(x, y int) int {
	return (x+30000)*60001 + (y + 30000)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
