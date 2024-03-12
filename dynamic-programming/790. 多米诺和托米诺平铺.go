package dynamic_programming

// https://leetcode.cn/problems/domino-and-tromino-tiling
// 有两种形状的瓷砖：一种是 2 x 1 的多米诺形，另一种是形如 "L" 的托米诺形。两种形状都可以旋转。
// 给定整数 n ，返回可以平铺 2 x n 的面板的方法的数量。返回对 109 + 7 取模 的值。
//
// 平铺指的是每个正方形都必须有瓷砖覆盖。两个平铺不同，当且仅当面板上有四个方向上的相邻单元中的两个，使得恰好有一个平铺有一个瓷砖占据两个正方形
func numTilings(n int) int {
	const mod int = 1e9 + 7
	f := make([][4]int, n+1)
	f[0][3] = 1
	for i := 1; i < n+1; i++ {
		f[i][0] = f[i-1][3]
		f[i][1] = (f[i-1][0] + f[i-1][2]) % mod
		f[i][2] = (f[i-1][0] + f[i-1][1]) % mod
		f[i][3] = (f[i-1][0] + f[i-1][1] + f[i-1][2] + f[i-1][3]) % mod
	}
	return f[n][3] % mod
}
