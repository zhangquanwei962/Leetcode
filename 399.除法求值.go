/*
 * @lc app=leetcode.cn id=399 lang=golang
 * @lcpr version=21909
 *
 * [399] 除法求值
 */

// @lc code=start
package main

import (
	"math/big"
)

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	id := map[string]int{}
	for _, eq := range equations {
		a, b := eq[0], eq[1]
		if _, has := id[a]; !has {
			id[a] = len(id)
		}
		if _, has := id[b]; !has {
			id[b] = len(id)
		}
	}

	graph := make([][]*big.Float, len(id))
	for i := range graph {
		graph[i] = make([]*big.Float, len(id))
		for j := range graph[i] {
			graph[i][j] = new(big.Float)
			if i == j {
				graph[i][j].SetFloat64(1.0)
			}
		}
	}
	for i, eq := range equations {
		v, w := id[eq[0]], id[eq[1]]
		graph[v][w].SetFloat64(values[i])
		graph[w][v].Quo(big.NewFloat(1.0), graph[v][w])
	}

	for k := range graph {
		for i := range graph {
			for j := range graph {
				if graph[i][k].Sign() != 0 && graph[k][j].Sign() != 0 {
					graph[i][j].Mul(graph[i][k], graph[k][j])
				}
			}
		}
	}

	ans := make([]float64, len(queries))
	for i, q := range queries {
		start, hasS := id[q[0]]
		end, hasE := id[q[1]]
		if !hasS || !hasE || graph[start][end].Sign() == 0 {
			ans[i] = -1
		} else {
			ans[i], _ = graph[start][end].Float64()
		}
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [["a","b"],["b","c"]]\n[2.0,3.0]\n[["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]\n
// @lcpr case=end

// @lcpr case=start
// [["a","b"],["b","c"],["bc","cd"]]\n[1.5,2.5,5.0]\n[["a","c"],["c","b"],["bc","cd"],["cd","bc"]]\n
// @lcpr case=end

// @lcpr case=start
// [["a","b"]]\n[0.5]\n[["a","b"],["b","a"],["a","c"],["x","y"]]\n
// @lcpr case=end

*/
