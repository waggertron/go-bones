package util

import (
	"sync"

	"github.com/wagertron/go-bones/model"
)

type Worker struct {
	work     <-chan model.Work // receive only channel
	wg       *sync.WaitGroup
	workFunc func(model.Work)
}

func (w Worker) Process(work model.Work) {
	w.workFunc(work)
	w.wg.Done()
}

func (w Worker) Start() {
	for work := range w.work {
		w.Process(work)
	}
}

type WorkMill struct {
	count    int
	workFunc func(model.Work)
}

func NewWorkMill(count int, workFunc func(model.Work)) *WorkMill {
	return &WorkMill{
		count:    count,
		workFunc: workFunc,
	}
}

func (wm *WorkMill) Process(totalWork []model.Work) {
	// create new work channel
	workChan := make(chan model.Work)
	// create new wait group
	var wg sync.WaitGroup

	// start wm.count worker processing async
	for i := 0; i < wm.count; i++ {
		w := Worker{
			work:     workChan,
			wg:       &wg,
			workFunc: wm.workFunc,
		}
		go w.Start()
	}

	for _, work := range totalWork {
		wg.Add(1)
		workChan <- work // buffered channel does not block
	}

	wg.Wait()       // needed for final n jobs to finish
	close(workChan) // lets go funcs exit from range and return
	return
}
