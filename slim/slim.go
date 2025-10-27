package slim

import (
	"context"
	"os"
	"os/exec"
	"strings"

	"github.com/livekit/protocol/utils/must"
)

// Each arg is an independent bash script, potentially multi-line.
func Bashes(scripts ...string) {
	for _, script := range scripts {
		must.Do(bashSingleContextError(context.Background(), script))
	}
}

// Each arg is a line of a single bash script.
func Bash(lines ...string) {
	Bashes(strings.Join(lines, "\n"))
}

func bashSingleContextError(ctx context.Context, script string) error {
	cmd := exec.CommandContext(ctx, "/usr/bin/env", "bash", "-c", script)
	cmd.Stdin = nil
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
