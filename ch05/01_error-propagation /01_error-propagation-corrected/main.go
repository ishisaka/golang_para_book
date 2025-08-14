// main.go
// エラーハンドリングの例。
// 実際にはサードパーティーのエラーハンドリング用ライブラリを使うのが良い。
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime/debug"
)

type MyError struct {
	Inner      error
	Message    string
	StackTrace string
	Misc       map[string]interface{}
}

func wrapError(err error, messagef string, msgArgs ...interface{}) MyError {
	return MyError{
		Inner:      err, // ここで包んでいるエラーを保管します。何が起きたのか調査する必要があるときに低水準のエラーをいつでも見れるようにしておきます。
		Message:    fmt.Sprintf(messagef, msgArgs...),
		StackTrace: string(debug.Stack()),        // この行はエラーが作られたときにスタックトレースを記録するためのものです。より洗練されたエラー型であればwrapErrorのスタックフレームを省略するでしょう。
		Misc:       make(map[string]interface{}), // ここで雑多な情報を保管するための場所を作ります。エラーの診断をする際に助けになる並行処理のIDやスタックトレースのハッシュ、あるいは他のコンテキストに関する情報を保管します。
	}
}

func (err MyError) Error() string {
	return err.Message
}

// "lowlevel" module

type LowLevelErr struct {
	error
}

func isGloballyExec(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, LowLevelErr{(wrapError(err, err.Error()))} // os.Statの呼び出しから発生する生のエラーをカスタマイズしたエラーで内包しています。
	}
	return info.Mode().Perm()&0100 == 0100, nil
}

// "intermediate" module

type IntermediateErr struct {
	error
}

func runJob(id string) error {
	const jobBinPath = "/bad/job/binary"
	isExecutable, err := isGloballyExec(jobBinPath)
	if err != nil {
		return IntermediateErr{wrapError(
			err,
			"cannot run job %q: requisite binaries not available",
			id,
		)} // lowlevelモジュールからのエラーを渡します。
	} else if isExecutable == false {
		return wrapError(
			nil,
			"cannot run job %q: requisite binaries are not executable",
			id,
		)
	}

	return exec.Command(jobBinPath, "--id="+id).Run()
}

func handleError(key int, err error, message string) {
	log.SetPrefix(fmt.Sprintf("[logID: %v]: ", key))
	log.Printf("%#v", err) // 何が起きたかを掘り下げる必要が出てきたときのためにすべてのエラーをログ出力しています。
	fmt.Printf("[%v] %v", key, message)
}

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	err := runJob("1")
	if err != nil {
		msg := "There was an unexpected issue; please report this as a bug."
		if _, ok := err.(IntermediateErr); ok { // ここでエラーが期待した型かどうかを確認しています。もしそうであれば、きちんとした形式のエラーなので、メッセージを単純にそのままユーザーに渡せます。
			msg = err.Error()
		}
		handleError(1, err, msg) // この行ではログとエラーメッセージをID1として紐付けています。IDは単調増加させることもできますし、一意なIDにするためにGUIDにしてもいいでしょう。
	}
}
