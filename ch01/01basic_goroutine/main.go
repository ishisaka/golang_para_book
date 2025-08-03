/*
これは結果が0になることを期待していますが、実際のところは動作させてみないと
わかりません。
*/
package main

func main() {
	var data int
	go func() {
		data++
	}()
	if data == 0 {
		println("data is still 0")
	} else {
		println("data has been incremented to", data)
	}
}
