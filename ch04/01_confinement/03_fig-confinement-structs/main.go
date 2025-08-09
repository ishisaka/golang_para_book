package main

import (
	"bytes"
	"fmt"
	"sync"
)

// レキシカルスコープによる拘束の例
func main() {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()

		var buff bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")
	go printData(&wg, data[:3]) // dataの中の先頭の3バイトを含んだスライスを渡します。
	go printData(&wg, data[3:]) // dataの中の後半の3バイトを含んだスライスを渡します。

	wg.Wait()
}
