package main

import (
	"fmt"
	"time"
)

// ゴルーチンのリークを防ぐ
// 親のゴルーチンから子のゴルーチンへキャンセルを送れるようにする
// 慣習としてDoneというチャネル名を与え、それを子のゴルーチンの第１引数にする。
func main() {
	doWork := func(
		done <-chan interface{},
		strings <-chan string,
	) <-chan interface{} { // doneチャネルをdoWork関数に渡します。慣例として、このチャネルは第1引数にします。
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					// Do something interesting
					fmt.Println(s)
				case <-done: // この行はどこにでもに存在するfor-selectパターンを使っています。
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	terminated := doWork(done, nil)

	go func() { // 1秒以上経過したらdoWorkの中で生成されたゴルーチンをキャンセルする他のゴルーチンを生成します。
		// Cancel the operation after 1 second.
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()

	<-terminated // ここでdoWorkから生成されたゴルーチンがメインゴルーチンとつながります。
	fmt.Println("Done.")
}
