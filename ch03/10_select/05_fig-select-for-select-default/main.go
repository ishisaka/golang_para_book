package main

import (
	"fmt"
	"time"
)

// main関数は非同期の終了シグナルを監視しながらループ内で作業をシミュレートします。
func main() {
	done := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	workCounter := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}

		// ゴールーチンの終了を待つ間別の仕事をする。
		workCounter++
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workCounter)
}
