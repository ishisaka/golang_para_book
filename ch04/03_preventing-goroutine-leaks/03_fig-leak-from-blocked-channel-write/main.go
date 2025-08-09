package main

import (
	"fmt"
	"math/rand"
)

// ゴルーチンがチャネルに対して書き込みを行おうとしてブロックしてしまう例
func main() {
	newRandStream := func() <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.") // ゴルーチンが無事に終了した場合にメッセージを表示します。が、実際には表示されません。
			defer close(randStream)
			for {
				randStream <- rand.Int() // ３回目以後読み込まれていないのでここでロックしてリークしてしまう、
			}
		}()

		return randStream
	}

	randStream := newRandStream()
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
}
