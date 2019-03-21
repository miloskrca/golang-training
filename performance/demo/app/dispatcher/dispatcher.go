package dispatcher

import "github.com/miloskrca/golang-training/performance/demo/app/worker"

type Dispatcher struct {
	clients map[string]chan<- worker.Job
}

func New() *Dispatcher {
	return &Dispatcher{make(map[string]chan<- worker.Job)}
}

func (d *Dispatcher) Reqister(client string, jobCh chan<- worker.Job) {
	d.clients[client] = jobCh
}

func (d *Dispatcher) Dispatch(job worker.Job) {
	// log.Printf("dispatching %d jobs", len(d.clients))
	for _, jobCh := range d.clients {
		jobCh <- job
	}
}
