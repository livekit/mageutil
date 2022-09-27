package mageutil

import (
	"context"
	"os"
	"os/exec"
	"strings"
)

func Command(ctx context.Context, command string) *exec.Cmd {
	args := strings.Split(command, " ")
	cmd := exec.CommandContext(ctx, args[0], args[1:]...)
	ConnectStd(cmd)
	return cmd
}

func CommandDir(ctx context.Context, dir, command string) *exec.Cmd {
	args := strings.Split(command, " ")
	cmd := exec.CommandContext(ctx, args[0], args[1:]...)
	cmd.Dir = dir
	ConnectStd(cmd)
	return cmd
}

func Run(ctx context.Context, commands ...string) error {
	for _, command := range commands {
		if err := Command(ctx, command).Run(); err != nil {
			return err
		}
	}
	return nil
}

func RunDir(ctx context.Context, dir string, commands ...string) error {
	for _, command := range commands {
		if err := CommandDir(ctx, dir, command).Run(); err != nil {
			return err
		}
	}
	return nil
}

func Pipe(first, second string) error {
	a1 := strings.Split(first, " ")
	c1 := exec.Command(a1[0], a1[1:]...)

	c1.Stderr = os.Stderr
	p, err := c1.StdoutPipe()
	if err != nil {
		return err
	}

	a2 := strings.Split(second, " ")
	c2 := exec.Command(a2[0], a2[1:]...)

	c2.Stdin = p
	c2.Stdout = os.Stdout
	c2.Stderr = os.Stderr

	if err = c1.Start(); err != nil {
		return err
	}
	if err = c2.Start(); err != nil {
		return err
	}
	if err = c1.Wait(); err != nil {
		return err
	}
	if err = c2.Wait(); err != nil {
		return err
	}
	return nil
}

func ConnectStd(cmd *exec.Cmd) {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
}
