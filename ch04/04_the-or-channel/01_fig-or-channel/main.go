package main

import (
	"fmt"
	"time"
)

// orチャネルの例
func main() {
	var or func(channels ...<-chan interface{}) <-chan interface{}
	or = func(channels ...<-chan interface{}) <-chan interface{} { // 関数orを定義
		switch len(channels) {
		case 0: // 可変長引数のスライスが空の場合には、単純にnilチャネルを返します。
			return nil
		case 1: // 可変長引数のスライスが1つしか要素を持っていない場合にはその要素を返すだけ
			return channels[0]
		}

		orDone := make(chan interface{})
		go func() { // ここが関数の本体で、再帰が発生する部分です。ゴルーチンを作って、ブロックすることなく作ったチャネルにメッセージを受け取れるようにします。
			defer close(orDone)

			switch len(channels) {
			case 2: // 再帰のやり方のせいで、orへの各再帰呼出しは少なくとも2つのチャネルを持っています。ゴルーチンの数を制限するために、2つしかチャネルがなかった場合の特別な条件を設定します。
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default: // スライスの3番目以降のチャネルから再帰的にorチャネルを作成して、そこからselectを行います。この再帰関係はスライスの残りの部分をorチャネルに分解して、最初のシグナルが返ってくる木構造を形成します。またorDoneチャネルも渡して、木構造の上位の部分が終了したら下位の部分も終了するようにしています。
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...): // <6>
				}
			}
		}()
		return orDone
	}
	sig := func(after time.Duration) <-chan interface{} { // この関数は単純にafterで指定された時間が経過したら閉じられるチャネルを生成します。
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now() // or関数から返されるチャネルがいつブロックされ始めたかを大まかに追跡します。
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v", time.Since(start)) // チャネルへの読み込みまでにかかった時間を表示します。
}
