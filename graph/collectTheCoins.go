package main

func collectTheCoins(coins []int, edges [][]int) int {
	n := len(coins)
	g := make([][]int, n)
	degree := make([]int, n)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
		degree[x]++
		degree[y]++
	}

	var q []int
	rest := n // 删除树中所有无金币的叶子节点，直到树中所有的叶子节点都是含有金币的
	for i := 0; i < n; i++ {
		if degree[i] == 1 && coins[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		degree[u]--
		rest--
		for _, v := range g[u] {
			degree[v]--
			if degree[v] == 1 && coins[v] == 0 {
				q = append(q, v)
			}
		}
	}

	// 删除树中所有的叶子节点, 连续删除2次
	for j := 0; j < 2; j++ {
		var q []int
		for i := 0; i < n; i++ {
			if degree[i] == 1 {
				q = append(q, i)
			}
		}
		for len(q) > 0 {
			u := q[0]
			q = q[1:]
			degree[u]--
			rest--
			for _, v := range g[u] {
				degree[v]--
			}
		}
	}

	if rest == 0 {
		return 0
	}
	return (rest - 1) * 2
}
