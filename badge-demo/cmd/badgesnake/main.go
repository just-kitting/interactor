package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"badgesnake/internal/gamestate"
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
	case "ui-sim":
		runUISim(args[1:])
	default:
		usage()
	}
}

func runUISim(args []string) {
	statePath := "/tmp/badgesnake/state.json"
	commandPath := "/tmp/badgesnake/command.json"
	seed := time.Now().UnixNano()
	doneHold := 1500 * time.Millisecond

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--state-file":
			i++
			if i >= len(args) {
				fatal(fmt.Errorf("missing value for --state-file"))
			}
			statePath = args[i]
		case "--command-file":
			i++
			if i >= len(args) {
				fatal(fmt.Errorf("missing value for --command-file"))
			}
			commandPath = args[i]
		case "--seed":
			i++
			if i >= len(args) {
				fatal(fmt.Errorf("missing value for --seed"))
			}
			value, err := strconv.ParseInt(args[i], 10, 64)
			if err != nil {
				fatal(err)
			}
			seed = value
		default:
			fatal(fmt.Errorf("unknown ui-sim option: %s", args[i]))
		}
	}

	sim := gamestate.NewSimulator(seed)
	lastSeq := int64(-1)
	lastStepAt := time.Now()
	_ = os.Remove(commandPath)

	for {
		cmd, ok, err := gamestate.ReadCommand(commandPath)
		if err != nil {
			fatal(err)
		}
		if ok && cmd.Seq != lastSeq {
			sim.ApplyCommand(cmd)
			lastSeq = cmd.Seq
			lastStepAt = time.Now()
		}

		snapshot := sim.Snapshot()
		if snapshot.Mode == "LIVE" && time.Since(lastStepAt) >= time.Duration(snapshot.StepMS)*time.Millisecond {
			sim.Step()
			lastStepAt = time.Now()
			snapshot = sim.Snapshot()
		} else if snapshot.Mode == "DONE" && time.Since(lastStepAt) >= doneHold {
			sim.ApplyCommand(gamestate.ControlCommand{Command: "reset"})
			lastStepAt = time.Now()
			snapshot = sim.Snapshot()
		}

		if err := gamestate.WriteSnapshot(statePath, snapshot); err != nil {
			fatal(err)
		}

		time.Sleep(50 * time.Millisecond)
	}
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func usage() {
	fmt.Println("usage: badgesnake <tokens|frame-example|ui-sim>")
}
