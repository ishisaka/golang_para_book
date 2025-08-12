package main

import (
	"fmt"
)

// ブリッジチャネルの例
func main() {
	orDone := func(done, c <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for {
				select {
				case <-done:
					return
				case v, ok := <-c:
					if ok == false {
						return
					}
					select {
					case valStream <- v:
					case <-done:
					}
				}
			}
		}()
		return valStream
	}
	/*
		以下のコードは「ブリッジ（bridge）」と呼ばれるパターンの実装で、
		チャネルのチャネル（<-chan <-chan interface{}）から流れてくる
		「複数のサブストリーム」を1本の値ストリーム（<-chan interface{}）に順に
		直列化して流す役目を持ちます。要するに、チャネルを「平坦化」して1本に
		束ねる関数です
	*/
	bridge := func(
		done <-chan interface{},
		chanStream <-chan <-chan interface{},
	) <-chan interface{} {
		valStream := make(chan interface{}) // これはbridgeからすべての値を返すチャネルです。
		go func() {
			defer close(valStream)
			for { // このループはchanStreamからチャネルを剥ぎ取り、ネストされたループに渡します。
				var stream <-chan interface{}
				select {
				case maybeStream, ok := <-chanStream:
					if ok == false {
						return
					}
					stream = maybeStream
				case <-done:
					return
				}
				for val := range orDone(done, stream) { // このループは渡されたチャネルから値を読み込みvalStreamにその値を渡す役割を担います
					select {
					case valStream <- val:
					case <-done:
					}
				}
			}
		}()
		return valStream
	}
	genVals := func() <-chan <-chan interface{} {
		chanStream := make(chan (<-chan interface{}))
		go func() {
			defer close(chanStream)
			for i := 0; i < 10; i++ {
				stream := make(chan interface{}, 1)
				stream <- i
				close(stream)
				chanStream <- stream
			}
		}()
		return chanStream
	}

	for v := range bridge(nil, genVals()) {
		fmt.Printf("%v ", v)
	}
}
