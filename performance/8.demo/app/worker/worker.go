package worker

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/miloskrca/golang-training/performance/demo/app/rabbitmq"
)

// Res is the result of the workers work on a Job
type Res struct {
	Time    time.Time
	Message string
}

// Run starts waiting for jobs that come from the `jobs` channel,
// does the work and sends results to the `results` channel
func Run(consumer string, jobs <-chan Job, results chan<- rabbitmq.Message) {
	for job := range jobs {
		// log.Printf("worker for consumer %s received a job %+v", consumer, job)

		// do the work
		res := Res{
			Time:    time.Now(),
			Message: fmt.Sprintf("%s processed for consumer: %s", job.Payload.Message, consumer),
		}

		// prepare for sending
		payload, err := json.Marshal(res)
		if err != nil {
			log.Printf("could not marshal result, consumer %s: %v", consumer, err)
			continue
		}
		msg := rabbitmq.Message{
			QueueName: consumer,
			Type:      "rabbitmq_message_type",
			Payload:   payload,
		}

		// send
		results <- msg
	}
}
