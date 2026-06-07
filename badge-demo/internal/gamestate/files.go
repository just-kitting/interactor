package gamestate

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

func WriteSnapshot(path string, snapshot Snapshot) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}

	tmpPath := path + ".tmp"
	body, err := json.Marshal(snapshot)
	if err != nil {
		return err
	}
	body = append(body, '\n')

	if err := os.WriteFile(tmpPath, body, 0o644); err != nil {
		return err
	}

	return os.Rename(tmpPath, path)
}

func ReadCommand(path string) (ControlCommand, bool, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return ControlCommand{}, false, nil
		}
		return ControlCommand{}, false, err
	}

	var cmd ControlCommand
	if err := json.Unmarshal(body, &cmd); err != nil {
		return ControlCommand{}, false, err
	}

	return cmd, true, nil
}
