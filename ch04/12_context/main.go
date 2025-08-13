package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background()) // mainがcontext.Background()で新しいContextを作り、それをcontext.WithCancelでキャンセルできるようにしています。
	defer cancel()

	wg.Go(func() {

		if err := printGreeting(ctx); err != nil {
			fmt.Printf("cannot print greeting: %v\n", err)
			cancel() // この行ではprintGreetingからエラーが返ってきたらmainがContextをキャンセルするようにしています。
		}
	})

	wg.Go(func() {
		if err := printFarewell(ctx); err != nil {
			fmt.Printf("cannot print farewell: %v\n", err)
		}
	})

	wg.Wait()
}

func printGreeting(ctx context.Context) error {
	greeting, err := genGreeting(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", greeting)
	return nil
}

func printFarewell(ctx context.Context) error {
	farewell, err := genFarewell(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", farewell)
	return nil
}

func genGreeting(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second) // genGreetingがContextをcontext.WithTimeoutで囲んでいます。これによって1秒後に戻されたContextを自動的にキャンセルして、それによってgenGreetingが以後Contextを渡すあらゆる子供、ここではつまりlocaleをキャンセルします。
	defer cancel()

	switch locale, err := locale(ctx); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "hello", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func genFarewell(ctx context.Context) (string, error) {
	switch locale, err := locale(ctx); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "goodbye", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func locale(ctx context.Context) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err() // この行はContextがキャンセルされた理由を返します。このエラーはmainまで伝搬され、これが2でのキャンセルを発生させます。
	case <-time.After(1 * time.Minute):
	}
	return "EN/US", nil
}
