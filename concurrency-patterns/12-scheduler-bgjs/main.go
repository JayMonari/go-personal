package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	fmt.Println("Process ID", os.Getpid())

	s := NewScheduler(5, 10)

	s.ListenForWork()

	fmt.Println("Listening for work")

	<-waitToExit()

	s.Exit()

	fmt.Println("Graceful shutdown")
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

//-

type Scheduler struct {
	workers  int
	msgCh    chan struct{}
	signalCh chan os.Signal
	wg       sync.WaitGroup
}

func NewScheduler(workers, buffer int) *Scheduler {
	return &Scheduler{
		workers:  workers,
		msgCh:    make(chan struct{}, buffer),
		signalCh: make(chan os.Signal, 1),
	}
}

func (s *Scheduler) ListenForWork() {
	go func() { // 1) Listen for messages to process
		signal.Notify(s.signalCh, syscall.SIGILL)
		for {
			<-s.signalCh

			s.msgCh <- struct{}{} // 2) Send to processing channel
		}
	}()

	s.wg.Add(s.workers)

	for i := 0; i < s.workers; i++ {
		go func(n int) {
			for {
				select {
				case _, open := <-s.msgCh: // 3) Wait for messages to process
					if !open {
						fmt.Printf("%d closing\n", n+1)
						s.wg.Done()
						return
					}

					fmt.Printf("%d<- Processing\n", n)
				}
			}
		}(i)
	}
}

func (s *Scheduler) Exit() {
	close(s.msgCh)
	s.wg.Wait()
}
