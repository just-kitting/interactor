package main

import (
	"fmt"
	"os"

	"badgesnake/internal/protocol"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		usage()
		return
	}

	switch args[0] {
	case "tokens":
		for _, endpoint := range protocol.EndpointTable {
			fmt.Printf("0x%02x %s %s %s\n", endpoint.Token, endpoint.Method, endpoint.Path, endpoint.Name)
		}
	case "frame-example":
		frame := protocol.NewRequestFrame(protocol.TokenMove, []byte(`{"move":"up"}`))
		encoded, err := frame.Encode()
		if err != nil {
			fatal(err)
		}
		fmt.Printf("%x\n", encoded)
	default:
		usage()
	}
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func usage() {
	fmt.Println("usage: badgesnake <tokens|frame-example>")
}
