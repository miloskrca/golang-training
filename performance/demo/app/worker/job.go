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

type Job struct {
	Type    string
	Payload struct {
		Message string
	}
}
