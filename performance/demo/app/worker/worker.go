package worker

import (
	"encoding/json"
	"log"

	"github.com/miloskrca/golang-training/performance/demo/app/rabbitmq"
)

type Res struct {
	Message string
}

func Run(client string, jobs <-chan Job, results chan<- rabbitmq.Message) {
	for job := range jobs {
		// log.Printf("worker for client %s received a job %+v", client, job)

		// do the work
		res := Res{
			Message: job.Payload.Message + "_" + client,
		}

		// prepare for sending
		payload, err := json.Marshal(res)
		if err != nil {
			log.Printf("could not marshal result, client %s: %v", client, err)
			continue
		}
		msg := rabbitmq.Message{
			QueueName: client,
			Type:      "rabbitmq_message_type",
			Payload:   payload,
		}

		// send
		results <- msg
	}
}
