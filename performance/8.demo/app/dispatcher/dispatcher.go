package dispatcher

import (
	"github.com/miloskrca/golang-training/performance/demo/app/worker"
)

// Dispatcher schedules every incoming job to all registered workers
type Dispatcher struct {
	workers map[string]chan<- worker.Job
}

// New creates a new dispatcher
func New() *Dispatcher {
	return &Dispatcher{
		workers: make(map[string]chan<- worker.Job),
	}
}

// Reqister registers a workers job channel so the dispacher can send it jobs
func (d *Dispatcher) Reqister(name string, jobCh chan<- worker.Job) {
	d.workers[name] = jobCh
}

// Dispatch sends the same jobs to all registered workers
func (d *Dispatcher) Dispatch(job worker.Job) {
	// log.Printf("dispatching %d jobs", len(d.workers))
	for _, jobCh := range d.workers {
		jobCh <- job
	}
}

// Stop closes all job channels
func (d *Dispatcher) Stop() {
	for worker, jobCh := range d.workers {
		close(jobCh)
		delete(d.workers, worker)
	}
}
