package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type FamilyMember struct {
	Name    string
	Age     int
	Parents []string
}

// // Bar will be initialized if there, nil if not there
// type Foo struct {
//     Bar *Bar
// }

type Command struct {
	Text string
}

type Message struct {
	Text string
	From string
}

type IncomingMessage struct {
	Cmd *Command
	Msg *Message
}

func main() {
	var m FamilyMember
	b := []byte(`{"Name": "name", "Age": 32, "Parents": ["mom", "dad"]}`)
	if err := json.Unmarshal(b, &m); err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)
	fmt.Println("========================")

	var incomingMessage1 IncomingMessage
	b1 := []byte(`{"Cmd": {"Text": "command"}}`)
	if err := json.Unmarshal(b1, &incomingMessage1); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", incomingMessage1)
	fmt.Printf("\tcommand: %s\n", incomingMessage1.Cmd.Text)
	fmt.Println("========================")

	var incomingMessage2 IncomingMessage
	b2 := []byte(`{"Msg": {"Text": "message", "From": "me"}}`)
	if err := json.Unmarshal(b2, &incomingMessage2); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", incomingMessage2)
	fmt.Printf("\tfrom:    %s\n\tmessage: %s\n", incomingMessage2.Msg.From, incomingMessage2.Msg.Text)
	fmt.Println("========================")
}
