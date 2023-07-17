/*
 * @lc app=leetcode.cn id=97 lang=golang
 * @lcpr version=21909
 *
 * [97] 交错字符串
 */

// @lc code=start
package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n, t := len(s1), len(s2), len(s3)
	if n+m != t {
		return false
	}

	f := make([]bool, n+1)

	f[0] = true

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			p := i + j - 1
			if i > 0 {
				f[j] = f[j] && s1[i-1] == s3[p]
			}
			if j > 0 {
				f[j] = f[j] || (f[j-1] && s2[j-1] == s3[p])
			}
		}
	}
	return f[n]
}

/*
func isInterleave(s1 string, s2 string, s3 string) bool {
    n, m, t := len(s1), len(s2), len(s3)
    if (n + m) != t {
        return false
    }
    f := make([][]bool, n + 1)
    for i := 0; i <= n; i++ {
        f[i] = make([]bool, m + 1)
    }
    f[0][0] = true
    for i := 0; i <= n; i++ {
        for j := 0; j <= m; j++ {
            p := i + j - 1
            if i > 0 {
                f[i][j] = f[i][j] || (f[i-1][j] && s1[i-1] == s3[p])
            }
            if j > 0 {
                f[i][j] = f[i][j] || (f[i][j-1] && s2[j-1] == s3[p])
            }
        }
    }
    return f[n][m]
}*/

// @lc code=end

/*
// @lcpr case=start
// "aabcc"\n"dbbca"\n"aadbbcbcac"\n
// @lcpr case=end

// @lcpr case=start
// "aabcc"\n"dbbca"\n"aadbbbaccc"\n
// @lcpr case=end

// @lcpr case=start
// ""\n""\n""\n
// @lcpr case=end

*/
