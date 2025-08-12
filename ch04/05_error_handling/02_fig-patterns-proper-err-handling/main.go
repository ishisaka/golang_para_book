package main

import (
	"fmt"
	"net/http"
)

// Goルーチンを使う場合のエラーハンドリングの方法例
func main() {
	type Result struct { // errorを含む型を作る
		Error    error
		Response *http.Response
	}
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result { // 先ほど作ったResult型のチャネルを返す
		results := make(chan Result)
		go func() {
			defer close(results)

			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				result = Result{Error: err, Response: resp} // Result型を初期化
				select {
				case <-done:
					return
				case results <- result: // Result型のインスタンスをチャネルに書き込む
				}
			}
		}()
		return results
	}

	done := make(chan interface{})
	defer close(done)

	urls := []string{"https://www.google.com", "https://badhost"}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil { // Goルーチンで発生したエラーをここでハンドリングする。
			fmt.Printf("error: %v", result.Error)
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}
