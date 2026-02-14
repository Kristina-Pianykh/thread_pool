package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	pool := CreateThreadPool(runtime.NumCPU())

	for i := range 20 {
		go pool.Execute(Task(func() {
			ticker(i)
		}))
	}

	fmt.Println("asynchronous")
	fmt.Println("finished main execution")

	v := <-signals
	fmt.Printf("received signal: %v\n", v)
	fmt.Println("Shutting down...")

	pool.Done()
}

func ticker(id int) {
	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()

	cnt := 0
	for {
		<-ticker.C
		cnt++
		fmt.Printf("tick: %d\n", id)
		if cnt == 3 {
			return
		}
	}
}
