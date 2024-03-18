package leetcode1976

// 下述思路解决失败，搜索会重会漏，暂未排查问题
// 官方思路 dijkstra算法

const (
	maxCost = 1e9
	mod     = 1e9 + 7
)

type Plan struct {
	// 提案的花费
	cost int
	// 相同花费的数量
	nums int
}

type Target struct {
	// 节点编号
	number int
	// 到节点的花费
	cost int
}

// 广搜+图dp
// dp[i] 表示从0到i的提案 最少时间Plan
// 状态-> 广搜图的遍历
//      当前node m-> 目标node n
// 决策:
/*
   1. if dp[m].cost + cost(m->n) == dp[n].cost, dp[n].nums++
   2. if dp[m].cost + cost(m->n) < dp[n.cost], dp[n] = Plan(dp[m].cost+cost(m->n), 1)
   3. if dp[m].cost + cost(m->n) > dp[n].cost, 劣势方案，舍弃
*/
func countPaths(n int, roads [][]int) int {
	// 构建出边数组，graph[i] 存储所有i可以到达的节点及花费 Target{n, cost}
	graph := make([][]Target, n)
	for _, road := range roads {
		// road[0] <-> road[1], cost: road[2]
		graph[road[0]] = append(graph[road[0]], Target{number: road[1], cost: road[2]})
		graph[road[1]] = append(graph[road[1]], Target{number: road[0], cost: road[2]})
	}
	// 构建状态dp
	dp := make([]Plan, n)
	// 初始状态，dp[0] 为 Plan{cost:0, nums:1}，没有方案, 其他：Plan{cost:max, nums:1}
	for i := range dp {
		dp[i] = Plan{cost: maxCost, nums: 1}
	}
	dp[0] = Plan{cost: 0, nums: 1}
	// bfs
	queue := []int{0}
	// 由于无向图，graph节点存储双向，搜索会重复。记忆化，访问过的节点就不访问咯
	visited := make([]bool, n)
	visited[0] = true
	for len(queue) > 0 {
		cur := queue[0]
		targets := graph[cur]
		queue = queue[1:]
		for _, target := range targets {
			// 更新dp, 节点入队
			if dp[cur].cost+target.cost == dp[target.number].cost {
				dp[target.number].nums = (dp[target.number].nums + 1) % mod
			} else if dp[cur].cost+target.cost < dp[target.number].cost {
				// 特殊边界，如果目标值被降低了，需要重新入队，更新答案（巨坑，找了好久）
				visited[target.number] = false
				dp[target.number] = Plan{
					cost: dp[cur].cost + target.cost,
					nums: 1,
				}
			}
			if visited[target.number] {
				// 如果节点入队过，就不用再入队了
				continue
			}
			queue = append(queue, target.number)
			visited[target.number] = true
		}
	}
	return dp[n-1].nums
}
