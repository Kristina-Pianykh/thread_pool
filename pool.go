package main

import (
	"context"
	"fmt"
	"sync"
)

type Task func()

type ThreadPool struct {
	taskQueue  chan Task
	taskWg     sync.WaitGroup
	workerWg   sync.WaitGroup
	cancelFunc context.CancelFunc
}

func (p *ThreadPool) Execute(task Task) {
	p.taskWg.Add(1)
	p.taskQueue <- task
}

func (p *ThreadPool) Done() {
	p.WaitTasks()
	fmt.Println("shutting down workers")
	p.cancelFunc()
	p.workerWg.Wait()
	// close(p.taskQueue)
}

func CreateThreadPool(size int) *ThreadPool {
	queue := make(chan Task, 100)
	p := ThreadPool{
		taskQueue: queue,
	}

	ctx, cancel := context.WithCancel(context.Background())
	p.cancelFunc = cancel

	for i := range size {
		worker := &Worker{ID: i}
		p.workerWg.Go(func() {
			worker.Launch(ctx, queue, &p.taskWg, &p.workerWg)
		})
	}

	return &p
}

func (p *ThreadPool) WaitTasks() {
	p.taskWg.Wait()
}

type Worker struct {
	ID int
}

func (w *Worker) Launch(ctx context.Context, queue chan Task, taskWg *sync.WaitGroup, workerWg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d shutting down\n", w.ID)
			workerWg.Done()
			return
		case task := <-queue:
			task()
			taskWg.Done()
		}
	}
}
