package main

import (
	"fmt"
	"time"
)

type Task func()

type Pool struct {
	maxWorkers int
	queue      chan Task
	done       chan struct{}
}

type Worker struct {
	id int
}

func main() {
	done := make(chan struct{})
	queue := make(chan Task, 100)
	for i := 0; i < 100; i++ {
		i := i
		queue <- func() {
			fmt.Println(i)
		}
	}

	pool := Pool{
		maxWorkers: 5,
		queue:      queue,
		done:       done,
	}

	pool.Start()
	time.Sleep(time.Second * 5)
	pool.Stop()
}

func (p Pool) Start() {
	for i := 0; i < p.maxWorkers; i++ {
		worker := Worker{i}
		go worker.Consume(p.queue, p.done)
	}
}

func (p Pool) Stop() {
	close(p.done)
}

func (w Worker) Consume(queue chan Task, stopCh chan struct{}) {
	for {
		select {
		case <-stopCh:
			fmt.Printf("worker %d stopped\n", w.id)
			return
		case task := <-queue:
			//fmt.Printf("worker %d do\n", w.id)
			//fmt.Printf("worker %d do %v\n", w.id, task)
			task()
		}
	}
}
