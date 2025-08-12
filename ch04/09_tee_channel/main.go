package main

import "fmt"

// or-doneチャネルとteeチャネルの例
func main() {
	// リピーター
	repeat := func(
		done <-chan interface{},
		values ...interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()
		return valueStream
	}
	// テイク、リピーターから受け取ったものをnum回処理する。
	take := func(
		done <-chan interface{},
		valueStream <-chan interface{},
		num int,
	) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}
	done := make(chan interface{})
	defer close(done)

	/*
		以下の関数は「or-done」パターンの実装で、キャンセル可能な転送用ラッパーです。
		入力チャネル c から値を受け取り、出力チャネル valStream に中継しますが、
		次のいずれかが起きたら即座に終了します。
		- done チャネルが閉じられた（キャンセル）とき
		- 入力チャネル c が閉じられた（読み尽くした）とき
	*/
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
		tee は入力チャネル in から受け取った各値を、2つの出力チャネル out1 と out2 の
		両方に送る「分岐（tee）」を実装しています。
			- キャンセル用の done を尊重し、or-done パターン経由で in を読みます。
			- 値1つにつき必ず2回送信し、片方が遅くても両方に届けるまでブロックします（バックプレッシャーが掛かる）。
	*/
	tee := func(
		done <-chan interface{},
		in <-chan interface{},
	) (_, _ <-chan interface{}) {
		out1 := make(chan interface{})
		out2 := make(chan interface{})
		go func() {
			defer close(out1)
			defer close(out2)
			for val := range orDone(done, in) {
				var out1, out2 = out1, out2 // 2つのチャネルのコピー変数としてout1とout2という2つのローカル変数を用意します。
				for i := 0; i < 2; i++ {    // 1つのselect文を使ってout1とout2への書き込みがお互いにブロックしないようにします。両方のチャネルに確実に書き込まれるように、select文を2回繰り返します。
					select {
					case out1 <- val:
						out1 = nil // チャネルへの書き込みが終わったら、コピー変数にnilを代入して、それ以降の書き込みをブロックしてもう片方のチャネルへの書き込みができるようにします。
					case out2 <- val:
						out2 = nil // チャネルへの書き込みが終わったら、コピー変数にnilを代入して、それ以降の書き込みをブロックしてもう片方のチャネルへの書き込みができるようにします。
					}
				}
			}
		}()
		return out1, out2
	}

	out1, out2 := tee(done, take(done, repeat(done, 1, 2), 4))

	for val1 := range out1 {
		fmt.Printf("out1: %v, out2: %v\n", val1, <-out2)
	}
}
