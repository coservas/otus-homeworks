package main

import (
	"os"
	"os/exec"
)

func RunCmd(cmd []string, env Environment) (returnCode int) {
	cmdName := cmd[0]
	command := exec.Command(cmdName, cmd[1:]...)

	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr

	for name, value := range env {
		if value.NeedRemove {
			os.Unsetenv(name)
			break
		}

		if _, ok := os.LookupEnv(name); ok {
			os.Unsetenv(name)
		}

		os.Setenv(name, value.Value)
	}

	_ = command.Run()

	return command.ProcessState.ExitCode()
}
