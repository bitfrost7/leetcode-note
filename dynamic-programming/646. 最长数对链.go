package dynamic_programming

import (
	"fmt"
	"sort"
)

//给你一个由 n 个数对组成的数对数组 pairs ，其中 pairs[i] = [lefti, righti] 且 lefti < righti 。
//
//现在，我们定义一种 跟随 关系，当且仅当 b < c 时，数对 p2 = [c, d] 才可以跟在 p1 = [a, b] 后面。我们用这种形式来构造 数对链 。
//
//找出并返回能够形成的 最长数对链的长度 。
//
//你不需要用到所有的数对，你可以以任何顺序选择其中的一些数对来构造。

//输入：pairs = [[1,2], [2,3], [3,4]]
//输出：2
//解释：最长的数对链是 [1,2] -> [3,4] 。

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})
	f := make([]int, 0)
	for _, v := range pairs {
		if i := sort.SearchInts(f, v[0]); i < len(f) {
			f[i] = min(f[i], v[1])
		} else {
			f = append(f, v[1])
		}
	}
	return len(f)
}

func TestFindLongestChain() {
	pairs := [][]int{
		{1, 2}, {2, 3}, {3, 4},
	}
	fmt.Printf("Input: pairs=%v", pairs)
	fmt.Printf("最长的数对链长度为:%d", findLongestChain(pairs))
}
