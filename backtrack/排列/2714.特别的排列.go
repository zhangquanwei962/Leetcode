package main

func specialPerm(nums []int) (ans int) {
	const mod int = 1e9 + 7
	n := len(nums)
	m := 1 << n

	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i == 0 {
			return 1
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}

		defer func() { *p = res }()

		for k, x := range nums {
			if x>>k&1 > 0 && (nums[j]%x == 0 || x%nums[j] == 0) {
				res = (res + dfs(i^(1<<k), k)) % mod
			}
		}
		return
	}

	for j := range nums {
		ans = (ans + dfs((m-1)^(1<<j), j)) % mod
	}
	return
}
