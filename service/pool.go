package service

import (
	"sync"
	"time"
	"web3.kz/solscan/config"
)

type RealExecutorPool struct {
	ExecutorsCount int
	Processor      Processor
}

type Task func()

type Worker struct {
	Id       uint
	JobQueue chan Task
}

var wg sync.WaitGroup

func (w *Worker) Run() {
	go func() {
		for task := range w.JobQueue {
			config.Log.Debugf("Worker-%d start execute task", w.Id)
			task()
			wg.Done()
		}
	}()
}

func (ep *RealExecutorPool) Execute() {
	taskQueue := make(chan Task)
	workers := make([]Worker, 5)
	for i := 1; i <= 5; i++ {
		workers[i-1] = Worker{
			Id:       uint(i),
			JobQueue: taskQueue,
		}
	}
	for _, w := range workers {
		w.Run()
	}
	go ep.schedule(taskQueue)
	wg.Wait()
	select {}
}

func (ep *RealExecutorPool) schedule(taskQueue chan Task) {
	config.Log.Info("Start analyse task")
	ticker := time.NewTicker(500 * time.Millisecond)

	defer ticker.Stop()

	for range ticker.C {
		wg.Add(1)
		config.Log.Debugf("<- Append task to queue")
		taskQueue <- func() {
			ep.Processor.Process()
		}
	}
}
