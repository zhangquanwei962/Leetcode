/*
 * @lc app=leetcode.cn id=1044 lang=golang
 * @lcpr version=21909
 *
 * [1044] 最长重复子串
 */

// @lc code=start
// 有单调性，归纳法证明
// size满足，那么size-1的长度也满足
package main

const P int64 = 131

func longestDupSubstring(s string) (ans string) {
	n := len(s)
	h := make([]int64, n+1)
	p := make([]int64, n+1)
	p[0] = 1
	// 偏移了1，str从1开始
	for i := 1; i <= n; i++ {
		h[i] = h[i-1]*P + int64(s[i-1])
		p[i] = p[i-1] * P
	}

	get := func(l, r int) int64 {
		return h[r] - h[l-1]*p[r-l+1]
	}

	check := func(str string, length int) bool {
		n := len(str)
		set := make(map[int64]bool)
		for i := 1; i+length-1 <= n; i++ {
			j := i + length - 1
			cur := get(i, j)
			if set[cur] {
				ans = str[i-1 : j]
				return true
			}
			set[cur] = true
		}
		return false
	}

	left, right := -1, n

	for left+1 < right {
		mid := left + (right-left)>>1
		t := check(s, mid)
		if t {
			left = mid
		} else {
			right = mid
		}
	}
	return
}

// class Solution {
//     typedef unsigned long long ULL;
//     const int p = 131;
//     vector<ULL> H, P;
//     unordered_set<ULL> st;
//     string ans;
// public:
//     string longestDupSubstring(string s) {
//         int n = s.size();
//         H.resize(n + 1), P.resize(n + 1);
//         P[0] = 1;
//         for (int i = 1; i <= n; i ++ )
//         {
//             H[i] = H[i - 1] * p + s[i - 1];
//             P[i] = P[i - 1] * p;
//         }

//         int l = 0, r = n;
//         while (l < r)
//         {
//             int mid = l + r + 1 >> 1;
//             if (check(s, mid)) l = mid;
//             else r = mid - 1;
//         }
//         return r == 0 ? "" : ans;
//     }

//     bool check(string &s, int x)
//     {
//         for (int i = 1; i + x - 1 <= s.size(); i ++ )
//         {
//             ULL t = get(i, i + x - 1);
//             if (st.count(t))
//             {
//                 ans = s.substr(i - 1, x);
//                 return true;
//             }
//             else st.insert(t);
//         }
//         return false;
//     }

//     ULL get(int l, int r)
//     {
//         return H[r] - H[l - 1] * P[r - l + 1];
//     }
// };

// @lc code=end

/*
// @lcpr case=start
// "banana"\n
// @lcpr case=end

// @lcpr case=start
// "abcd"\n
// @lcpr case=end

*/
