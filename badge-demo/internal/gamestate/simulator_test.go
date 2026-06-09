package gamestate

import "testing"

func TestSimulatorResetAndSnapshot(t *testing.T) {
	sim := NewSimulator(1)
	snapshot := sim.Snapshot()

	if snapshot.BoardSize != DefaultBoardSize {
		t.Fatalf("board size = %d, want %d", snapshot.BoardSize, DefaultBoardSize)
	}
	if len(snapshot.Snakes) != 2 {
		t.Fatalf("len(snakes) = %d, want 2", len(snapshot.Snakes))
	}
	if len(snapshot.Foods) < 2 {
		t.Fatalf("len(foods) = %d, want at least 2", len(snapshot.Foods))
	}
	if snapshot.Title == "" {
		t.Fatal("title = empty, want matchup title")
	}
}

func TestSimulatorCommands(t *testing.T) {
	sim := NewSimulator(1)

	sim.ApplyCommand(ControlCommand{Command: "faster"})
	if sim.Snapshot().StepMS != DefaultStepMS-150 {
		t.Fatalf("step ms after faster = %d", sim.Snapshot().StepMS)
	}

	sim.ApplyCommand(ControlCommand{Command: "pause_toggle"})
	if sim.Snapshot().Mode != "PAUSED" {
		t.Fatalf("mode after pause = %q, want PAUSED", sim.Snapshot().Mode)
	}

	sim.ApplyCommand(ControlCommand{Command: "reset"})
	snapshot := sim.Snapshot()
	if snapshot.Turn != 0 {
		t.Fatalf("turn after reset = %d, want 0", snapshot.Turn)
	}
	if snapshot.Mode != "LIVE" {
		t.Fatalf("mode after reset = %q, want LIVE", snapshot.Mode)
	}
	if snapshot.MatchNumber < 2 {
		t.Fatalf("match number after reset = %d, want >= 2", snapshot.MatchNumber)
	}
}
