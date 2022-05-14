package bcmds

import (
	"context"
	"os/exec"
	"strings"
	"time"
)

func Docker(cmds ...string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	cmd := exec.CommandContext(ctx, "/usr/local/bin/docker", cmds...)
	s, err := cmd.CombinedOutput()

	return s, err
}

func DockerDir(dir string, cmds ...string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "/usr/local/bin/docker", cmds...)
	cmd.Dir = dir
	s, err := cmd.CombinedOutput()

	return s, err
}

func Bash(cmds string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	t := append([]string{"-c"}, strings.Fields(cmds)...)

	cmd := exec.CommandContext(ctx, "bash", t...)
	s, err := cmd.CombinedOutput()

	return s, err
}
