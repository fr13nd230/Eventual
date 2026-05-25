package lib

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"sync"
	"time"

	mael "github.com/jepsen-io/maelstrom/demo/go"
)

func UniqueIDs(n *mael.Node) {
	n.Handle("generate", func(msg mael.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		
		body["type"] = "generate_ok"
		body["id"] = generateUniqueID()

		return n.Reply(msg, body)
	})
}

func generateUniqueID() string {
    lock := sync.RWMutex{}
    lock.RLock()
    defer lock.RUnlock()
	ts := time.Now().Unix()
	suffix, _ := rand.Int(rand.Reader, big.NewInt(ts))
	return fmt.Sprintf("%v%v", ts, suffix)
}
