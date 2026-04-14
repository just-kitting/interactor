package simtransport

import (
	"context"
	"encoding/json"
	"fmt"

	"badgesnake/internal/player"
	"badgesnake/internal/protocol"

	"github.com/BattlesnakeOfficial/rules/client"
)

type InMemoryAdapter struct {
	Player player.Player
}

func (a InMemoryAdapter) HandleFrame(ctx context.Context, req protocol.Frame) (protocol.Frame, error) {
	if a.Player == nil {
		return protocol.Frame{}, fmt.Errorf("player is nil")
	}

	switch req.Code {
	case protocol.TokenInfo:
		body, err := a.handleInfo(ctx)
		if err != nil {
			return protocol.NewResponseFrame(protocol.StatusInternal, nil), err
		}
		return protocol.NewResponseFrame(protocol.StatusSuccess, body), nil
	case protocol.TokenStart:
		return a.handleSnakeRequest(ctx, req.Payload, a.Player.Start)
	case protocol.TokenMove:
		return a.handleMoveRequest(ctx, req.Payload)
	case protocol.TokenEnd:
		return a.handleSnakeRequest(ctx, req.Payload, a.Player.End)
	default:
		return protocol.NewResponseFrame(protocol.StatusUnsupported, nil), nil
	}
}

func (a InMemoryAdapter) handleInfo(ctx context.Context) ([]byte, error) {
	response, err := a.Player.Metadata(ctx)
	if err != nil {
		return nil, err
	}
	return json.Marshal(response)
}

func (a InMemoryAdapter) handleSnakeRequest(
	ctx context.Context,
	payload []byte,
	fn func(context.Context, client.SnakeRequest) ([]byte, error),
) (protocol.Frame, error) {
	request, err := decodeSnakeRequest(payload)
	if err != nil {
		return protocol.NewResponseFrame(protocol.StatusBadRequest, nil), err
	}

	body, err := fn(ctx, request)
	if err != nil {
		return protocol.NewResponseFrame(protocol.StatusInternal, nil), err
	}

	return protocol.NewResponseFrame(protocol.StatusSuccess, body), nil
}

func (a InMemoryAdapter) handleMoveRequest(ctx context.Context, payload []byte) (protocol.Frame, error) {
	request, err := decodeSnakeRequest(payload)
	if err != nil {
		return protocol.NewResponseFrame(protocol.StatusBadRequest, nil), err
	}

	response, err := a.Player.Move(ctx, request)
	if err != nil {
		return protocol.NewResponseFrame(protocol.StatusInternal, nil), err
	}

	body, err := json.Marshal(response)
	if err != nil {
		return protocol.NewResponseFrame(protocol.StatusInternal, nil), err
	}

	return protocol.NewResponseFrame(protocol.StatusSuccess, body), nil
}

func decodeSnakeRequest(payload []byte) (client.SnakeRequest, error) {
	var request client.SnakeRequest
	if err := json.Unmarshal(payload, &request); err != nil {
		return client.SnakeRequest{}, err
	}
	return request, nil
}
