package main

import (
	"context"
	"fmt"
)

func main() {
	//ctx, cancel := context.WithCancel(context.Background())
	//
	//quit := make(chan os.Signal)
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	//go gg(ctx)
	//<-quit
	//cancel()
	//time.Sleep(1 * time.Second)

}

func gg(ctx context.Context) {
	fmt.Println("hello")

	<-ctx.Done()
	fmt.Println("world")

}
