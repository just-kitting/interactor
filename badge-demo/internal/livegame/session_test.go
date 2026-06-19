package livegame

import (
	"context"
	"errors"
	"testing"
	"time"

	"badgesnake/internal/player"

	"github.com/BattlesnakeOfficial/rules/client"
)

func TestSessionStartsAndSteps(t *testing.T) {
	session, err := NewSession(Config{
		Seed:          1,
		TimeoutMS:     500,
		LiveName:      "Zepto",
		OpponentName:  "Clocky",
		LiveBus:       "1",
		LiveAddress:   0x42,
		ProbeInterval: time.Millisecond,
	}, player.FixturePlayer{
		MetadataResponse: client.SnakeMetadataResponse{
			APIVersion: "1",
			Color:      "#000000",
		},
		MoveDecision: client.MoveResponse{Move: "right"},
	}, player.SimulatedPlayer{
		Name:     "Clocky",
		MoveMode: "clockwise",
	})
	if err != nil {
		t.Fatalf("NewSession() error = %v", err)
	}

	if err := session.EnsureStarted(context.Background(), time.Now()); err != nil {
		t.Fatalf("EnsureStarted() error = %v", err)
	}

	snapshot := session.Snapshot()
	if snapshot.Mode != "LIVE" {
		t.Fatalf("mode = %q, want LIVE", snapshot.Mode)
	}
	if len(snapshot.Snakes) != 2 {
		t.Fatalf("len(snakes) = %d, want 2", len(snapshot.Snakes))
	}

	if err := session.Step(context.Background()); err != nil {
		t.Fatalf("Step() error = %v", err)
	}

	snapshot = session.Snapshot()
	if snapshot.Turn != 1 {
		t.Fatalf("turn = %d, want 1", snapshot.Turn)
	}
	if len(snapshot.Foods) == 0 {
		t.Fatal("foods = 0, want at least one food")
	}
}

func TestSessionWaitsWhenLiveMetadataFails(t *testing.T) {
	session, err := NewSession(Config{
		LiveName:      "Zepto",
		OpponentName:  "Clocky",
		LiveBus:       "1",
		LiveAddress:   0x42,
		ProbeInterval: time.Millisecond,
	}, failingPlayer{err: errors.New("no ack")}, player.SimulatedPlayer{
		Name:     "Clocky",
		MoveMode: "clockwise",
	})
	if err != nil {
		t.Fatalf("NewSession() error = %v", err)
	}

	if err := session.EnsureStarted(context.Background(), time.Now()); err != nil {
		t.Fatalf("EnsureStarted() error = %v", err)
	}

	snapshot := session.Snapshot()
	if snapshot.Mode != "WAITING" {
		t.Fatalf("mode = %q, want WAITING", snapshot.Mode)
	}
	if snapshot.EventText == "" {
		t.Fatal("event text = empty, want wait reason")
	}
}

type failingPlayer struct {
	err error
}

func (p failingPlayer) Metadata(context.Context) (client.SnakeMetadataResponse, error) {
	return client.SnakeMetadataResponse{}, p.err
}

func (p failingPlayer) Start(context.Context, client.SnakeRequest) ([]byte, error) {
	return nil, p.err
}

func (p failingPlayer) Move(context.Context, client.SnakeRequest) (client.MoveResponse, error) {
	return client.MoveResponse{}, p.err
}

func (p failingPlayer) End(context.Context, client.SnakeRequest) ([]byte, error) {
	return nil, p.err
}
