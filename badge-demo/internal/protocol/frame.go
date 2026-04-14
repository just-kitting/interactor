package protocol

import (
	"encoding/binary"
	"fmt"
)

const HeaderSize = 6

type Frame struct {
	Version uint8
	Code    uint8
	Flags   uint8
	Payload []byte
}

func NewRequestFrame(token uint8, payload []byte) Frame {
	return Frame{
		Version: Version,
		Code:    token,
		Flags:   0,
		Payload: payload,
	}
}

func NewResponseFrame(status uint8, payload []byte) Frame {
	return Frame{
		Version: Version,
		Code:    status,
		Flags:   0,
		Payload: payload,
	}
}

func (f Frame) Encode() ([]byte, error) {
	if len(f.Payload) > 0xffff {
		return nil, fmt.Errorf("payload too large: %d", len(f.Payload))
	}

	buf := make([]byte, HeaderSize+len(f.Payload))
	buf[0] = f.Version
	buf[1] = f.Code
	buf[2] = f.Flags
	buf[3] = 0
	binary.LittleEndian.PutUint16(buf[4:6], uint16(len(f.Payload)))
	copy(buf[6:], f.Payload)

	return buf, nil
}

func DecodeFrame(buf []byte) (Frame, error) {
	if len(buf) < HeaderSize {
		return Frame{}, fmt.Errorf("frame too short: %d", len(buf))
	}

	payloadLen := int(binary.LittleEndian.Uint16(buf[4:6]))
	if len(buf) != HeaderSize+payloadLen {
		return Frame{}, fmt.Errorf("frame length mismatch: got=%d want=%d", len(buf), HeaderSize+payloadLen)
	}

	payload := make([]byte, payloadLen)
	copy(payload, buf[6:])

	return Frame{
		Version: buf[0],
		Code:    buf[1],
		Flags:   buf[2],
		Payload: payload,
	}, nil
}
