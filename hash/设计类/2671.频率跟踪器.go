/*
 * @lc app=leetcode.cn id=2671 lang=golang
 * @lcpr version=21909
 *
 * [2671] 频率跟踪器
 */

// @lc code=start
package main

type FrequencyTracker struct {
	cnt  map[int]int // 总共个数
	freq map[int]int // 有多少个
}

func Constructor() (_ FrequencyTracker) {
	return FrequencyTracker{map[int]int{}, map[int]int{}}
}

func (f *FrequencyTracker) Add(number int) {
	f.freq[f.cnt[number]]--
	f.cnt[number]++
	f.freq[f.cnt[number]]++
}

func (f *FrequencyTracker) DeleteOne(number int) {
	if f.cnt[number] == 0 {
		return // 不删除任何内容
	}
	f.freq[f.cnt[number]]--
	f.cnt[number]--
	f.freq[f.cnt[number]]++
}

func (f *FrequencyTracker) HasFrequency(frequency int) bool {
	return f.freq[frequency] > 0
}

/**
 * Your FrequencyTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * obj.DeleteOne(number);
 * param_3 := obj.HasFrequency(frequency);
 */
// @lc code=end
