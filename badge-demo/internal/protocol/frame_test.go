package protocol

import "testing"

func TestEncodeDecodeFrameRoundTrip(t *testing.T) {
	original := NewRequestFrame(TokenMove, []byte(`{"move":"up"}`))

	encoded, err := original.Encode()
	if err != nil {
		t.Fatalf("Encode() error = %v", err)
	}

	decoded, err := DecodeFrame(encoded)
	if err != nil {
		t.Fatalf("DecodeFrame() error = %v", err)
	}

	if decoded.Version != original.Version {
		t.Fatalf("Version mismatch: got %d want %d", decoded.Version, original.Version)
	}
	if decoded.Code != original.Code {
		t.Fatalf("Code mismatch: got %d want %d", decoded.Code, original.Code)
	}
	if decoded.Flags != original.Flags {
		t.Fatalf("Flags mismatch: got %d want %d", decoded.Flags, original.Flags)
	}
	if string(decoded.Payload) != string(original.Payload) {
		t.Fatalf("Payload mismatch: got %q want %q", string(decoded.Payload), string(original.Payload))
	}
}

func TestDecodeFrameRejectsShortBuffer(t *testing.T) {
	if _, err := DecodeFrame([]byte{1, 2, 3}); err == nil {
		t.Fatal("DecodeFrame() expected error for short buffer")
	}
}

func TestLookupEndpoint(t *testing.T) {
	endpoint, ok := LookupEndpoint(TokenStart)
	if !ok {
		t.Fatal("LookupEndpoint(TokenStart) returned ok=false")
	}
	if endpoint.Path != "/start" {
		t.Fatalf("Path mismatch: got %q want %q", endpoint.Path, "/start")
	}
}
