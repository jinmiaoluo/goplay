package main

// 通过回调函数判断 arithmetic progression(等差数列) geometric progression(等比数列)
// 如果不用回调函数. 我们需要单独实现等差数列和等比数列的比较函数. 比如 APExamineProgression 和 GPExamineProgression
// 通过使用回调函数. 将不同数列的判断逻辑通过函数值传入. 这样, 在调用时. 只需传入不同的函数值即可实现不同的操作. 提高代码的复用

import (
	"fmt"
	"io"
	"os"
)

var conditionmap = map[string]func(int, int) int{
	"ap": func(i1 int, i2 int) int {
		return i2 - i1
	},
	"gp": func(i1 int, i2 int) int {
		return i2 / i1
	},
}

type sequence []int

func (s sequence) readinput(r io.Reader) {
	for j := 0; j < len(s); j++ {
		fmt.Fscan(r, &s[j])
	}
}

func (s sequence) examineprogression(condition func(a, b int) int) bool {
	if len(s) < 3 {
		return false
	}
	d := condition(s[1], s[0])
	for i := 2; i < len(s); i++ {
		if condition(s[i], s[i-1]) != d {
			return false
		}
	}
	return true
}

func main() {
	var seqlen int
	fmt.Scan(&seqlen)

	s := make(sequence, seqlen)
	s.readinput(os.Stdin)

	switch {
	case s.examineprogression(conditionmap["ap"]):
		fmt.Println("ap")
	case s.examineprogression(conditionmap["gp"]):
		fmt.Println("gp")
	default:
		fmt.Println("random")
	}
}
