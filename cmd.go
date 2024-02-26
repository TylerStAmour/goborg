package goborg

import (
	"encoding/json"
	"os/exec"
)

// Exec will execute a Borg command with the given arguments and produce a JSON output.
func (b *BorgEnv) Exec(command string, args ...string) (map[string]any, error) {
	// one-liner to combine command and args, with the command coming first
	combinedArgs := append([]string{command, "--log-json"}, args...)

	cmd := exec.Command("borg", combinedArgs...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	var jsonOut map[string]any
	err = json.Unmarshal(out, &jsonOut)
	if err != nil {
		return nil, err
	}

	return jsonOut, nil
}
