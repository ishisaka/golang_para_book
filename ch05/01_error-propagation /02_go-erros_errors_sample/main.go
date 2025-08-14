// main.go
// github.com/go-errors/errorsを使ったエラーハンドリングの簡単な例
package main

import (
	"fmt"

	"github.com/go-errors/errors"
)

var Crashed = errors.Errorf("oh dear")

func Crash() error {
	return errors.New(Crashed)
}

func main() {
	err := Crash()
	if err != nil {
		if errors.Is(err, Crashed) {
			fmt.Println(err.(*errors.Error).ErrorStack())
		} else {
			panic(err)
		}
	}
}
