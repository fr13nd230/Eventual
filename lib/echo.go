package lib

import (
	"encoding/json"

	mael "github.com/jepsen-io/maelstrom/demo/go"
)

func Echo(n *mael.Node) {
	n.Handle("echo", func(msg mael.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		body["type"] = "echo_ok"
		return n.Reply(msg, body)
	})
}
