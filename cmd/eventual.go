package main

import (
	"log"

	"github.com/fr13nd230/Eventual/lib"
	mael "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := mael.NewNode()

    lib.Echo(n)
    lib.UniqueIDs(n)

	if err := n.Run(); err != nil {
		log.Fatalf("Failed to run the node %v", err)
	}
}
