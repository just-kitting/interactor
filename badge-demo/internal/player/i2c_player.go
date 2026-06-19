package player

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"badgesnake/internal/protocol"

	"github.com/BattlesnakeOfficial/rules/client"
)

const defaultI2CReadLength = 512

type I2CTransferFunc func(context.Context, ...string) ([]byte, error)

type I2CConfig struct {
	Bus            string
	Address        int
	MaxResponseLen int
	ForceTransfer  bool
	Transfer       I2CTransferFunc
}

type I2CPlayer struct {
	bus            string
	address        int
	maxResponseLen int
	forceTransfer  bool
	transfer       I2CTransferFunc
}

func NewI2CPlayer(cfg I2CConfig) (*I2CPlayer, error) {
	bus := strings.TrimSpace(cfg.Bus)
	if bus == "" {
		return nil, fmt.Errorf("i2c bus is required")
	}

	address := cfg.Address
	if address < 0 || address > 0x7f {
		return nil, fmt.Errorf("i2c address 0x%x out of range", address)
	}

	maxResponseLen := cfg.MaxResponseLen
	if maxResponseLen == 0 {
		maxResponseLen = defaultI2CReadLength
	}
	if maxResponseLen < protocol.HeaderSize {
		return nil, fmt.Errorf("max response length must be at least %d", protocol.HeaderSize)
	}

	transfer := cfg.Transfer
	if transfer == nil {
		transfer = defaultI2CTransfer
	}

	forceTransfer := cfg.ForceTransfer
	if !cfg.ForceTransfer {
		forceTransfer = true
	}

	return &I2CPlayer{
		bus:            strings.TrimPrefix(strings.TrimPrefix(bus, "i2c-"), "/dev/i2c-"),
		address:        address,
		maxResponseLen: maxResponseLen,
		forceTransfer:  forceTransfer,
		transfer:       transfer,
	}, nil
}

func (p *I2CPlayer) Metadata(ctx context.Context) (client.SnakeMetadataResponse, error) {
	body, err := p.exchange(ctx, protocol.TokenInfo, nil)
	if err != nil {
		return client.SnakeMetadataResponse{}, err
	}

	var response client.SnakeMetadataResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return client.SnakeMetadataResponse{}, err
	}
	return response, nil
}

func (p *I2CPlayer) Start(ctx context.Context, request client.SnakeRequest) ([]byte, error) {
	return p.exchangeSnakeRequest(ctx, protocol.TokenStart, request)
}

func (p *I2CPlayer) Move(ctx context.Context, request client.SnakeRequest) (client.MoveResponse, error) {
	body, err := p.exchangeSnakeRequest(ctx, protocol.TokenMove, request)
	if err != nil {
		return client.MoveResponse{}, err
	}

	var response client.MoveResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return client.MoveResponse{}, err
	}
	return response, nil
}

func (p *I2CPlayer) End(ctx context.Context, request client.SnakeRequest) ([]byte, error) {
	return p.exchangeSnakeRequest(ctx, protocol.TokenEnd, request)
}

func (p *I2CPlayer) exchangeSnakeRequest(ctx context.Context, token uint8, request client.SnakeRequest) ([]byte, error) {
	payload, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	return p.exchange(ctx, token, payload)
}

func (p *I2CPlayer) exchange(ctx context.Context, token uint8, payload []byte) ([]byte, error) {
	reqFrame, err := protocol.NewRequestFrame(token, payload).Encode()
	if err != nil {
		return nil, err
	}

	if _, err := p.transfer(ctx, p.writeArgs(reqFrame)...); err != nil {
		return nil, fmt.Errorf("i2c write failed on bus %s addr 0x%02x: %w", p.bus, p.address, err)
	}

	rawResponse, err := p.transfer(ctx, p.readArgs()...)
	if err != nil {
		return nil, fmt.Errorf("i2c read failed on bus %s addr 0x%02x: %w", p.bus, p.address, err)
	}

	respFrame, err := protocol.DecodeFramePrefix(parseI2CTransferOutput(rawResponse))
	if err != nil {
		return nil, fmt.Errorf("invalid i2c response frame from bus %s addr 0x%02x: %w", p.bus, p.address, err)
	}
	if respFrame.Version != protocol.Version {
		return nil, fmt.Errorf("unsupported i2c response version %d", respFrame.Version)
	}
	if respFrame.Code != protocol.StatusSuccess {
		return nil, fmt.Errorf("i2c request failed with status %s", describeI2CStatus(respFrame.Code))
	}

	return respFrame.Payload, nil
}

func (p *I2CPlayer) writeArgs(payload []byte) []string {
	args := make([]string, 0, 4+len(payload))
	if p.forceTransfer {
		args = append(args, "-f")
	}
	args = append(args, "-y", p.bus, fmt.Sprintf("w%d@0x%02x", len(payload), p.address))
	for _, b := range payload {
		args = append(args, fmt.Sprintf("0x%02x", b))
	}
	return args
}

func (p *I2CPlayer) readArgs() []string {
	args := make([]string, 0, 4)
	if p.forceTransfer {
		args = append(args, "-f")
	}
	args = append(args, "-y", p.bus, fmt.Sprintf("r%d@0x%02x", p.maxResponseLen, p.address))
	return args
}

func defaultI2CTransfer(ctx context.Context, args ...string) ([]byte, error) {
	bin := os.Getenv("BADGESNAKE_I2CTRANSFER_BIN")
	if strings.TrimSpace(bin) == "" {
		bin = "i2ctransfer"
	}

	cmd := exec.CommandContext(ctx, bin, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("%s %v: %w: %s", bin, args, err, strings.TrimSpace(string(output)))
	}
	return output, nil
}

func parseI2CTransferOutput(output []byte) []byte {
	fields := strings.Fields(string(output))
	if len(fields) == 0 {
		return nil
	}

	buf := make([]byte, 0, len(fields))
	for _, field := range fields {
		cleaned := strings.TrimSuffix(strings.TrimPrefix(strings.ToLower(strings.TrimSpace(field)), "0x"), ",")
		if cleaned == "" {
			continue
		}
		value, err := strconv.ParseUint(cleaned, 16, 8)
		if err != nil {
			return nil
		}
		buf = append(buf, byte(value))
	}
	return buf
}

func describeI2CStatus(status uint8) string {
	switch status {
	case protocol.StatusSuccess:
		return "success"
	case protocol.StatusBadRequest:
		return "bad-request"
	case protocol.StatusUnsupported:
		return "unsupported"
	case protocol.StatusBusy:
		return "busy"
	case protocol.StatusInternal:
		return "internal"
	default:
		return fmt.Sprintf("0x%02x", status)
	}
}
