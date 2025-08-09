package main

import (
	"fmt"
)

// ゴールーチンリークの例
func main() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for s := range strings {
				// Do something interesting
				fmt.Println(s)
			}
		}()
		return completed
	}

	doWork(nil) // nilチャンネルを渡しているので、doWorkを含みゴールーチンはプロセスが生きている限り残ってしまう。
	// Perhaps more work is done here
	fmt.Println("Done.")
}
