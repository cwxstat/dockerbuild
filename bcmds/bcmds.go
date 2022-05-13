package bcmds

import (
	"context"
	"os/exec"
	"time"
)

func Docker(cmds ...string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	cmd := exec.CommandContext(ctx, "/usr/local/bin/docker", cmds...)
	s, err := cmd.CombinedOutput()

	return s, err
}
