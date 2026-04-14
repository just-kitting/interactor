package player

import (
	"context"

	"github.com/BattlesnakeOfficial/rules/client"
)

type Player interface {
	Metadata(context.Context) (client.SnakeMetadataResponse, error)
	Start(context.Context, client.SnakeRequest) ([]byte, error)
	Move(context.Context, client.SnakeRequest) (client.MoveResponse, error)
	End(context.Context, client.SnakeRequest) ([]byte, error)
}
