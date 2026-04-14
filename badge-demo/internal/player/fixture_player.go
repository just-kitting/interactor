package player

import (
	"context"

	"github.com/BattlesnakeOfficial/rules/client"
)

type FixturePlayer struct {
	MetadataResponse client.SnakeMetadataResponse
	MoveDecision     client.MoveResponse
	StartBody        []byte
	EndBody          []byte
}

func (p FixturePlayer) Metadata(context.Context) (client.SnakeMetadataResponse, error) {
	return p.MetadataResponse, nil
}

func (p FixturePlayer) Start(context.Context, client.SnakeRequest) ([]byte, error) {
	return append([]byte(nil), p.StartBody...), nil
}

func (p FixturePlayer) Move(context.Context, client.SnakeRequest) (client.MoveResponse, error) {
	return p.MoveDecision, nil
}

func (p FixturePlayer) End(context.Context, client.SnakeRequest) ([]byte, error) {
	return append([]byte(nil), p.EndBody...), nil
}
