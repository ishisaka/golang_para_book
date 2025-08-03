/*
これは一見よさそうですが、実際には混乱をさらに加速させます。
*/
package main

import "time"

func main() {
	var data int
	go func() {
		data++
	}()
	time.Sleep(100 * time.Millisecond) // ゴルーチンが実行されるのを待つ
	if data == 0 {
		println("data is still 0")
	} else {
		println("data has been incremented to", data)
	}
}
