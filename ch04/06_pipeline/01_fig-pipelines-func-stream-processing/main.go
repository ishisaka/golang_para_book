package main

import (
	"fmt"
)

// パイプライン処理の単純な例
func main() {
	// パイプラインのステージその1
	multiply := func(value, multiplier int) int {
		return value * multiplier
	}
	// パイプラインのステージその2
	add := func(value, additive int) int {
		return value + additive
	}

	ints := []int{1, 2, 3, 4}
	for _, v := range ints {
		// パイプラインのステージを接続
		fmt.Println(multiply(add(multiply(v, 2), 1), 2))
	}
}
