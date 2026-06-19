package player

import (
	"context"
	"strings"

	"github.com/BattlesnakeOfficial/rules/client"
)

type SimulatedPlayer struct {
	Name     string
	MoveMode string
	Color    string
	Head     string
	Tail     string
	Author   string
	Version  string
}

func (p SimulatedPlayer) Metadata(context.Context) (client.SnakeMetadataResponse, error) {
	return client.SnakeMetadataResponse{
		APIVersion: "1",
		Author:     defaultString(p.Author, "BadgeSnake"),
		Color:      defaultString(p.Color, "#a85f20"),
		Head:       defaultString(p.Head, "default"),
		Tail:       defaultString(p.Tail, "default"),
		Version:    defaultString(p.Version, "sim"),
	}, nil
}

func (p SimulatedPlayer) Start(context.Context, client.SnakeRequest) ([]byte, error) {
	return []byte(`{"transport":"sim"}`), nil
}

func (p SimulatedPlayer) Move(_ context.Context, request client.SnakeRequest) (client.MoveResponse, error) {
	return client.MoveResponse{
		Move:  p.moveForTurn(request.Turn),
		Shout: "simulated",
	}, nil
}

func (p SimulatedPlayer) End(context.Context, client.SnakeRequest) ([]byte, error) {
	return []byte(`{"transport":"sim"}`), nil
}

func (p SimulatedPlayer) moveForTurn(turn int) string {
	switch strings.ToLower(strings.TrimSpace(p.MoveMode)) {
	case "up", "down", "left", "right":
		return strings.ToLower(strings.TrimSpace(p.MoveMode))
	case "clockwise":
		return []string{"up", "right", "down", "left"}[turn%4]
	case "counterclockwise":
		return []string{"up", "left", "down", "right"}[turn%4]
	default:
		return "up"
	}
}

func defaultString(value string, fallback string) string {
	if strings.TrimSpace(value) == "" {
		return fallback
	}
	return value
}
