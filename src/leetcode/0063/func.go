package leetcode63

// 线性动态规划
// 分析：
//     1. 机器人每次只能向下或者向右移动=>那么每个位置只能来源于上方或者左方
//     2. dp[i][j]表示机器人到达i，j位置时的总路径数，易得到 dp[i][j] = dp[i-1][j] + dp[i][j-1]
//     3. 障碍物不可达：所以if 障碍物, dp[i][j] = 0
//     4. 初始状态：处于0,0位置，有且只有一种方法 dp[0][0]=1
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp := make([][]int, 0, len(obstacleGrid))
	for i := 0; i < len(obstacleGrid); i++ {
		dp = append(dp, make([]int, len(obstacleGrid[0])))
	}
	dp[0][0] = 1
	for i, row := range obstacleGrid {
		for j, v := range row {
			// 如果当前是障碍物，则不可达
			if v == 1 {
				dp[i][j] = 0
				continue
			}
			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}
	return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
