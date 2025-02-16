package service

import (
	"sync"
	"web3.kz/solscan/config"
)

type RealExecutorPool struct {
	ExecutorsCount int
	Processor      Processor
	TaskQueue      chan Task
}

type Task func()

type Worker struct {
	Id       uint
	JobQueue chan Task
}

func (w *Worker) Run(wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		task := <-w.JobQueue
		config.Log.Infof("Worker-%d start execute task", w.Id)
		task()
	}()
}

func (ep *RealExecutorPool) Execute() {
	var wg sync.WaitGroup

	workers := make([]Worker, ep.ExecutorsCount)

	for i := 0; i < len(workers); i++ {
		workers = append(workers, Worker{
			Id:       uint(i),
			JobQueue: ep.TaskQueue,
		})
	}

	for _, w := range workers {
		wg.Add(1)
		w.Run(&wg)
	}

	wg.Wait()
}
