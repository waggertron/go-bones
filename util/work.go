package util

// do not communicate by sharing memory, share memory by communicating

import (
	"sync"

	"github.com/wagertron/go-bones/model"
)

type Worker struct {
	work     <-chan model.Work // receive only channel
	wg       *sync.WaitGroup   // to parallel count work done and in progress
	workFunc func(model.Work)  // a generic work function
}

func (w Worker) Process(work model.Work) {
	w.workFunc(work)
	w.wg.Done()
}

func (w Worker) Start() {
	// will block until work added, then proceed until channel closed
	for work := range w.work { // blocks until work is added into the channel, when all work is exhuasted, blocks until channel closure 
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
	// create new work channel for this process
	workChan := make(chan model.Work)
	// create new wait group for this process, for each item of work, add to the wait group
	var wg sync.WaitGroup

	// start wm.count workers processing async
	for i := 0; i < wm.count; i++ {
		w := Worker{
			work:     workChan,
			wg:       &wg,
			workFunc: wm.workFunc,
		}
		go w.Start() // intiates worker processing, but does not block the main thread
	}

	for _, work := range totalWork {
		wg.Add(1)
		workChan <- work // blocks until worker recieves 
	}

	wg.Wait()
	close(workChan) // lets go funcs exit from range and return
	return
}
