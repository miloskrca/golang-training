package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/miloskrca/golang-training/performance/demo/app/dispatcher"
	"github.com/miloskrca/golang-training/performance/demo/app/rabbitmq"
	"github.com/miloskrca/golang-training/performance/demo/app/worker"

	_ "net/http/pprof"
)

var (
	cpuprofile string
	numQueues  int
)

func main() {
	flag.IntVar(&numQueues, "q", 10, "num of queues")
	flag.Parse()

	var queueNames []string
	for i := 0; i < numQueues; i++ {
		queueNames = append(queueNames, fmt.Sprintf("client_%d", i))
	}
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
	go rmq.Start()
	log.Println("rabbitmq service started...")

	d := dispatcher.New()

	results := make(chan rabbitmq.Message)
	for _, queue := range queueNames {
		jobs := make(chan worker.Job)
		d.Reqister(queue, jobs)
		go worker.Run(queue, jobs, results)
	}
	log.Println("workers started...")

	go func() {
		http.HandleFunc("/", handler(d))
		log.Println("server started...")
		log.Fatal(http.ListenAndServe(":8000", nil))
	}()

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
