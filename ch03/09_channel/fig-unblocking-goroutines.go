package main

import (
	"fmt"
	"sync"
)

// チャンネルを閉じることで複数のゴールーチンにシグナルを送信する例
func main() {
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin // チャンネルが読み込めるようになるまでゴールーチンは待機
			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutines...")
	close(begin) // チャンネルを閉じることで全てのゴールーチンを開放
	wg.Wait()
}
