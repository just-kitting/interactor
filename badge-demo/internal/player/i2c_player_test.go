package player

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"badgesnake/internal/protocol"

	"github.com/BattlesnakeOfficial/rules/client"
)

func TestI2CPlayerMetadataRoundTrip(t *testing.T) {
	var callCount int
	var writeArgs []string
	var readArgs []string

	metadataBody, err := json.Marshal(client.SnakeMetadataResponse{
		APIVersion: "1",
		Author:     "Ada",
		Color:      "#112233",
	})
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}
	responseFrame, err := protocol.NewResponseFrame(protocol.StatusSuccess, metadataBody).Encode()
	if err != nil {
		t.Fatalf("Encode() error = %v", err)
	}

	p, err := NewI2CPlayer(I2CConfig{
		Bus:     "i2c-1",
		Address: 0x42,
		Transfer: func(_ context.Context, args ...string) ([]byte, error) {
			callCount++
			switch callCount {
			case 1:
				writeArgs = append([]string(nil), args...)
				return []byte(""), nil
			case 2:
				readArgs = append([]string(nil), args...)
				return []byte(asTransferOutput(responseFrame)), nil
			default:
				t.Fatalf("unexpected transfer call %d", callCount)
				return nil, nil
			}
		},
	})
	if err != nil {
		t.Fatalf("NewI2CPlayer() error = %v", err)
	}

	metadata, err := p.Metadata(context.Background())
	if err != nil {
		t.Fatalf("Metadata() error = %v", err)
	}

	if metadata.Author != "Ada" {
		t.Fatalf("Author mismatch: got %q want %q", metadata.Author, "Ada")
	}
	if len(writeArgs) < 4 || writeArgs[2] != "1" || writeArgs[3] != "w6@0x42" {
		t.Fatalf("unexpected write args: %v", writeArgs)
	}
	if len(readArgs) < 4 || readArgs[3] != "r512@0x42" {
		t.Fatalf("unexpected read args: %v", readArgs)
	}
}

func TestI2CPlayerReturnsTransportStatusError(t *testing.T) {
	var callCount int
	responseFrame, err := protocol.NewResponseFrame(protocol.StatusBusy, nil).Encode()
	if err != nil {
		t.Fatalf("Encode() error = %v", err)
	}

	p, err := NewI2CPlayer(I2CConfig{
		Bus:     "1",
		Address: 0x42,
		Transfer: func(_ context.Context, args ...string) ([]byte, error) {
			callCount++
			if callCount == 1 {
				return []byte(""), nil
			}
			return []byte(asTransferOutput(responseFrame)), nil
		},
	})
	if err != nil {
		t.Fatalf("NewI2CPlayer() error = %v", err)
	}

	_, err = p.Metadata(context.Background())
	if err == nil {
		t.Fatal("Metadata() error = nil, want busy status error")
	}
	if !strings.Contains(err.Error(), "busy") {
		t.Fatalf("Metadata() error = %v, want busy message", err)
	}
}

func asTransferOutput(buf []byte) string {
	parts := make([]string, 0, len(buf))
	for _, b := range buf {
		parts = append(parts, strings.ToLower("0x"+hexByte(b)))
	}
	return strings.Join(parts, " ")
}

func hexByte(value byte) string {
	const digits = "0123456789abcdef"
	return string([]byte{digits[value>>4], digits[value&0x0f]})
}
