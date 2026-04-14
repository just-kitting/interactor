package simtransport

import (
	"context"
	"encoding/json"
	"testing"

	"badgesnake/internal/player"
	"badgesnake/internal/protocol"

	"github.com/BattlesnakeOfficial/rules/client"
)

func TestHandleMoveFrame(t *testing.T) {
	adapter := InMemoryAdapter{
		Player: player.FixturePlayer{
			MoveDecision: client.MoveResponse{
				Move:  "left",
				Shout: "fixture",
			},
		},
	}

	request := client.SnakeRequest{
		Game: client.Game{
			ID:      "game-1",
			Timeout: 500,
		},
		Turn: 1,
		Board: client.Board{
			Height: 11,
			Width:  11,
			Snakes: []client.Snake{},
			Food:   []client.Coord{},
		},
		You: client.Snake{
			ID: "player-1",
		},
	}

	payload, err := json.Marshal(request)
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}

	response, err := adapter.HandleFrame(context.Background(), protocol.NewRequestFrame(protocol.TokenMove, payload))
	if err != nil {
		t.Fatalf("HandleFrame() error = %v", err)
	}

	if response.Code != protocol.StatusSuccess {
		t.Fatalf("status mismatch: got %d want %d", response.Code, protocol.StatusSuccess)
	}

	var decoded client.MoveResponse
	if err := json.Unmarshal(response.Payload, &decoded); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}

	if decoded.Move != "left" {
		t.Fatalf("move mismatch: got %q want %q", decoded.Move, "left")
	}
	if decoded.Shout != "fixture" {
		t.Fatalf("shout mismatch: got %q want %q", decoded.Shout, "fixture")
	}
}

func TestHandleUnsupportedFrame(t *testing.T) {
	adapter := InMemoryAdapter{
		Player: player.FixturePlayer{},
	}

	response, err := adapter.HandleFrame(context.Background(), protocol.NewRequestFrame(0xff, nil))
	if err != nil {
		t.Fatalf("HandleFrame() error = %v", err)
	}

	if response.Code != protocol.StatusUnsupported {
		t.Fatalf("status mismatch: got %d want %d", response.Code, protocol.StatusUnsupported)
	}
}
