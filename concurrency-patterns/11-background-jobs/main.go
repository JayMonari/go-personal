package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// XXX: use go build, using go run will not get the correct PID.
func main() {
	fmt.Println("Process ID", os.Getpid())

	listenForWork()

	<-waitToExit()

	fmt.Println("Exited gracefully.")
}

func listenForWork() {
	const workersN = 5

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGTERM)

	workersCh := make(chan struct{}, workersN)

	go func() {
		for {
			<-sc

			workersCh <- struct{}{}
		}
	}()

	go func() {
		var workers int

		for range workersCh {
			workerID := (workers % workersN) + 1
			workers++

			fmt.Printf("%d<-\n", workerID)

			go func() {
				doWork(workerID)
			}()

		}
	}()
}

func waitToExit() <-chan struct{} {
	runCh := make(chan struct{}, 1)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)

	go func() {
		defer close(runCh)
		<-sc
	}()

	return runCh
}

func doWork(id int) {
	fmt.Printf("<-%d starting\n", id)
	time.Sleep(3 * time.Second)
	fmt.Printf("<-%d completed\n", id)
}
