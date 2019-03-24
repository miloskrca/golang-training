package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime/pprof"
	"time"

	"github.com/miloskrca/golang-training/performance/demo/app/dispatcher"
	"github.com/miloskrca/golang-training/performance/demo/app/rabbitmq"
	"github.com/miloskrca/golang-training/performance/demo/app/worker"
	// _ "net/http/pprof"
)

var (
	cpuprofile string
	numQueues  int
)

func main() {
	flag.IntVar(&numQueues, "q", 10, "num of queues")
	flag.StringVar(&cpuprofile, "cpuprofile", "", "write cpu profile to file")
	flag.Parse()

	// collect cpu profile while running
	if cpuprofile != "" {
		f, err := os.Create(cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	// create a list of queue names that will be used by consumers
	var queueNames []string
	for i := 0; i < numQueues; i++ {
		queueNames = append(queueNames, fmt.Sprintf("consumer_%d", i))
	}

	// create and start a RabbitMQ service that waits on workers to finish
	// and sends every finished job to the corresponding queue
	config := rabbitmq.RabbitMQConf{
		Hostname: "localhost",
		Port:     "5672",
		User:     "guest",
		Password: "guest",
		VHost:    "performance",
		Queues:   queueNames,
		Exchange: "performance_exchange",
	}
	rmq, err := rabbitmq.New(config)
	if err != nil {
		log.Fatal(err)
	}
	go rmq.Run()
	log.Println("rabbitmq service started")

	d := dispatcher.New()

	// channel where the workers write their results
	results := make(chan rabbitmq.Message)

	// create a worker per consumer
	for _, queue := range queueNames {
		jobs := make(chan worker.Job)
		// regiter the job channel with the dispatcher
		d.Reqister(queue, jobs)
		// start a worker that waits for jobs on `jobs` channel and writes the reusults in `results` channel
		go worker.Run(queue, jobs, results)
	}
	log.Println("workers started")

	go func() {
		// register a handler that expects requests and dispatches jobs using the dispatcher
		http.HandleFunc("/", handler(d))
		log.Println("server started")
		log.Fatal(http.ListenAndServe(":8000", nil))
	}()

	// if we are profiling stop after a timeout
	if cpuprofile != "" {
		go func() {
			time.Sleep(60 * time.Second)
			log.Println("60s expired, stopping")
			d.Stop()
			close(results)
		}()
	}

	// for every result a worker did send a RabbitMQ message
	for msg := range results {
		rmq.Send(msg)
	}
}

func handler(d *dispatcher.Dispatcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var job worker.Job
		if err := json.Unmarshal(body, &job); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		d.Dispatch(job)
	}
}
