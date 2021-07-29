package strexec

import (
	"context"
	"os/exec"

	"github.com/mattn/go-shellwords"
)

func Command(cmd string) (*exec.Cmd, error) {
	cmdSlice, err := sliceCmd(cmd)
	if err != nil {
		return nil, err
	}
	name, arg := cmdSlice[0], cmdSlice[1:]
	return exec.Command(name, arg...), nil
}

func CommandContext(ctx context.Context, cmd string) (*exec.Cmd, error) {
	cmdSlice, err := sliceCmd(cmd)
	if err != nil {
		return nil, err
	}
	name, arg := cmdSlice[0], cmdSlice[1:]
	return exec.CommandContext(ctx, name, arg...), nil
}

func sliceCmd(args string) ([]string, error) {
	slice, err := shellwords.Parse(args)
	if err != nil {
		return nil, err
	}
	return slice, nil
}
