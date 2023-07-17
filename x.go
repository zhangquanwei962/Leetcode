package main

import "fmt"

func main() {
	col := 26
	var ans string
	for col > 0 {
		x := (col - 1) % 26
		ans = string(x+'A') + ans
		col = (col - 1) / 26
	}
	fmt.Println(ans)
}
