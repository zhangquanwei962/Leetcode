/*
 * @lc app=leetcode.cn id=1396 lang=golang
 * @lcpr version=21909
 *
 * [1396] 设计地铁系统
 */

// @lc code=start
package main

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		startInfo: make(map[int]Start),
		table:     make(map[StartEnd]SumAndAmount),
	}
}

type Start struct {
	StationName string
	Time        int
}

type StartEnd struct {
	StartStation string
	EndStation   string
}

type SumAndAmount struct {
	Sum    int
	Amount int
}

type UndergroundSystem struct {
	startInfo map[int]Start
	table     map[StartEnd]SumAndAmount
}

func (us *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	us.startInfo[id] = Start{
		StationName: stationName,
		Time:        t,
	}
}

func (us *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	startTime := us.startInfo[id].Time
	key := StartEnd{
		StartStation: us.startInfo[id].StationName,
		EndStation:   stationName,
	}
	us.table[key] = SumAndAmount{
		Sum:    us.table[key].Sum + (t - startTime),
		Amount: us.table[key].Amount + 1,
	}
}

func (us *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	key := StartEnd{
		StartStation: startStation,
		EndStation:   endStation,
	}
	sum := us.table[key].Sum
	amount := us.table[key].Amount
	return float64(sum) / float64(amount)
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
// @lc code=end
