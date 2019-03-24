package worker

/*

Example message:

{
	"type": "message_type",
	"payload": {
		"message": "hi"
	}
}

*/

// Job is one job that a worker works on
type Job struct {
	Type    string
	Payload struct {
		Message string
	}
}
