package util

import (
	"sync"

	"github.com/wagertron/go-bones/model"
)

type Worker struct {
	work <-chan model.Work
	wg   *sync.WaitGroup
}

func (w Worker) Process(model.Work) {
	w.wg.Done()
}

func (w Worker) Start() {
	for work := range w.work {
		w.Process(work)
	}
}

type WorkMill struct {
	workers  []Worker
	workChan chan<- model.Work
	wg       *sync.WaitGroup
}

func NewWorkMill(n int) *WorkMill {
	workers := make([]Worker, n)
	workChan := make(chan model.Work, n)

	var wg sync.WaitGroup

	for i := 0; i <= n; i++ {
		workers[i] = Worker{
			work: workChan,
			wg:   &wg,
		}
	}

	return &WorkMill{
		workers:  workers,
		workChan: workChan,
		wg:       &wg,
	}
}

func (wm *WorkMill) Process(totalWork []model.Work) {
	var wg sync.WaitGroup

	for _, worker := range wm.workers {
		go worker.Start()
	}

	for _, work := range totalWork {
		wg.Add(1)
		wm.workChan <- work
	}

	wg.Wait() //needed for final n jobs to finish

	return
}
