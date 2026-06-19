package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"badgesnake/internal/gamestate"
	"badgesnake/internal/livegame"
	"badgesnake/internal/player"
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
	case "ui-live":
		runUILive(args[1:])
	default:
		usage()
	}
}

func runUISim(args []string) {
	statePath := "/tmp/badgesnake/state.json"
	commandPath := "/tmp/badgesnake/command.json"
	seed := time.Now().UnixNano()
	doneHold := 1500 * time.Millisecond
	matchupID := ""

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
		case "--matchup":
			i++
			if i >= len(args) {
				fatal(fmt.Errorf("missing value for --matchup"))
			}
			matchupID = args[i]
		default:
			fatal(fmt.Errorf("unknown ui-sim option: %s", args[i]))
		}
	}

	sim := gamestate.NewSimulator(seed)
	if matchupID != "" {
		if err := sim.SetMatchup(matchupID); err != nil {
			fatal(err)
		}
		sim.Reset()
	}
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

func runUILive(args []string) {
	statePath := "/tmp/badgesnake/state.json"
	commandPath := "/tmp/badgesnake/command.json"
	seed := time.Now().UnixNano()
	doneHold := 1500 * time.Millisecond
	timeoutMS := 500
	bus := "1"
	addr := 0x42
	maxResponseLen := 512
	liveName := "Zepto"
	opponentName := "Clocky"
	opponentMode := "clockwise"
	probeIntervalMS := 1000

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--state-file":
			i++
			statePath = requireArg(args, i, "--state-file")
		case "--command-file":
			i++
			commandPath = requireArg(args, i, "--command-file")
		case "--seed":
			i++
			seed = parseInt64Arg(requireArg(args, i, "--seed"), "--seed")
		case "--bus":
			i++
			bus = strings.TrimSpace(requireArg(args, i, "--bus"))
		case "--addr":
			i++
			addr = parseHexIntArg(requireArg(args, i, "--addr"), "--addr")
		case "--max-response-len":
			i++
			maxResponseLen = parseIntArg(requireArg(args, i, "--max-response-len"), "--max-response-len")
		case "--timeout-ms":
			i++
			timeoutMS = parseIntArg(requireArg(args, i, "--timeout-ms"), "--timeout-ms")
		case "--live-name":
			i++
			liveName = requireArg(args, i, "--live-name")
		case "--opponent-name":
			i++
			opponentName = requireArg(args, i, "--opponent-name")
		case "--opponent-mode":
			i++
			opponentMode = requireArg(args, i, "--opponent-mode")
		case "--probe-interval-ms":
			i++
			probeIntervalMS = parseIntArg(requireArg(args, i, "--probe-interval-ms"), "--probe-interval-ms")
		default:
			fatal(fmt.Errorf("unknown ui-live option: %s", args[i]))
		}
	}

	livePlayer, err := player.NewI2CPlayer(player.I2CConfig{
		Bus:            bus,
		Address:        addr,
		MaxResponseLen: maxResponseLen,
	})
	if err != nil {
		fatal(err)
	}

	session, err := livegame.NewSession(livegame.Config{
		Seed:             seed,
		TimeoutMS:        timeoutMS,
		LiveName:         liveName,
		OpponentName:     opponentName,
		OpponentMoveMode: opponentMode,
		LiveBus:          bus,
		LiveAddress:      addr,
		ProbeInterval:    time.Duration(probeIntervalMS) * time.Millisecond,
	}, livePlayer, player.SimulatedPlayer{
		Name:     opponentName,
		MoveMode: opponentMode,
	})
	if err != nil {
		fatal(err)
	}

	lastSeq := int64(-1)
	lastStepAt := time.Now()
	_ = os.Remove(commandPath)

	for {
		cmd, ok, err := gamestate.ReadCommand(commandPath)
		if err != nil {
			fatal(err)
		}
		if ok && cmd.Seq != lastSeq {
			session.ApplyCommand(cmd)
			lastSeq = cmd.Seq
			lastStepAt = time.Now()
		}

		snapshot := session.Snapshot()
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutMS+250)*time.Millisecond)
		switch snapshot.Mode {
		case "WAITING":
			if err := session.EnsureStarted(ctx, time.Now()); err != nil {
				cancel()
				fatal(err)
			}
			snapshot = session.Snapshot()
			if snapshot.Mode == "LIVE" {
				lastStepAt = time.Now()
			}
		case "LIVE":
			if time.Since(lastStepAt) >= time.Duration(snapshot.StepMS)*time.Millisecond {
				if err := session.Step(ctx); err != nil {
					cancel()
					fatal(err)
				}
				lastStepAt = time.Now()
				snapshot = session.Snapshot()
			}
		case "DONE":
			if time.Since(lastStepAt) >= doneHold {
				session.Reset()
				lastStepAt = time.Now()
				snapshot = session.Snapshot()
			}
		}
		cancel()

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
	fmt.Println("usage: badgesnake <tokens|frame-example|ui-sim|ui-live>")
}

func requireArg(args []string, index int, name string) string {
	if index >= len(args) {
		fatal(fmt.Errorf("missing value for %s", name))
	}
	return args[index]
}

func parseIntArg(value string, name string) int {
	parsed, err := strconv.Atoi(value)
	if err != nil {
		fatal(fmt.Errorf("invalid value for %s: %w", name, err))
	}
	return parsed
}

func parseInt64Arg(value string, name string) int64 {
	parsed, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		fatal(fmt.Errorf("invalid value for %s: %w", name, err))
	}
	return parsed
}

func parseHexIntArg(value string, name string) int {
	parsed, err := strconv.ParseInt(value, 0, 16)
	if err != nil {
		fatal(fmt.Errorf("invalid value for %s: %w", name, err))
	}
	return int(parsed)
}
