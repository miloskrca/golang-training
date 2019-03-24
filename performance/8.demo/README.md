# Performance demo

## Commands
* `make setup` - run RabbitMQ
* `make run` - run the app
* `make load` - run the requests
* `make consume` - run the consumer that reads from RabbitMQ and displays the number of messages

## Architecture

`make   [load]         [app]              [setup]      [consume]`

                 |--> worker_0 -->|    |-----------|--> consumer_0 -->|    
                 |--> worker_1 -->|    |           |--> consumer_1 -->|    
        /doc --->|--> worker_2 -->|--> | RabbitMQ  |--> consumer_2 -->|--> status
                 |--> worker_3 -->|    |           |--> consumer_3 -->|    
                 |--> worker_4 -->|    |-----------|--> consumer_4 -->|    

## Job
{
	"type": "message_type",
	"payload": {
		"message": "hi"
	}
}

## Status

Output shown by the consumer:

```
client_6 received messages: 12
client_8 received messages: 12
client_5 received messages: 12
client_7 received messages: 12
client_0 received messages: 12
client_9 received messages: 12
client_3 received messages: 12
client_2 received messages: 12
client_1 received messages: 12
client_4 received messages: 12
Total received messages: 120
```