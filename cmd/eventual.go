package main

import (
	"encoding/json"
	"log"

	mael "github.com/jepsen-io/maelstrom/demo/go"
)

// type Message[T string] struct {
// 	Src  string         `json:"src"`
// 	Dest string         `json:"dest"`
// 	Body MessageBody[T] `json:"body"`
// }

// type MessageBody[T string] struct {
// 	Type      T       `json:"type"`
// 	MsgID     *string `json:"msg_id"`
// 	InReplyTo *string `json:"in_reply_to"`
// }

// func NewMessage[T string](Src, Dest string, Body MessageBody[T]) *Message[T] {
// 	return &Message[T]{
// 		Src:  Src,
// 		Dest: Dest,
// 		Body: MessageBody[T]{
// 			Type:      Body.Type,
// 			MsgID:     Body.MsgID,
// 			InReplyTo: Body.InReplyTo,
// 		},
// 	}
// }

func main() {
	n := mael.NewNode()

	n.Handle("echo", func(msg mael.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		body["type"] = "echo_ok"
		return n.Reply(msg, body)
	})

	if err := n.Run(); err != nil {
		log.Fatalf("Failed to run the node %v", err)
	}
}
