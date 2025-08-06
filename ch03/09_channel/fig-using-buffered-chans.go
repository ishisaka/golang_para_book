package main

import (
	"bytes"
	"fmt"
	"os"
)

// バッファ付きチャネルの例
func main() {
	var stdoutBuff bytes.Buffer         // バッファ
	defer stdoutBuff.WriteTo(os.Stdout) // 最後にバッファを標準出力に出力

	intStream := make(chan int, 4) // キャパシティ4のチャンエルを作る
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done.")
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received %v.\n", integer)
	}
}
