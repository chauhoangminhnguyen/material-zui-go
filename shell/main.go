package mz_tmp

import (
	"os"
	"os/exec"
	"strings"
)

func ExecuteSingle(command string) error {
	command = strings.Trim(command, "\n")
	command = strings.Trim(command, " ")
	args := strings.Split(command, " ")
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func ExecuteWithOutput(command string) (string, error) {
	command = strings.Trim(command, "\n")
	args := strings.Split(command, " ")
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func Execute(command string) []error {
	command = strings.Trim(command, "\n")
	commands := strings.Split(command, "&&")
	return ExecuteMultiple(commands)
}

func ExecuteMultiple(commands []string) []error {
	errors := []error{}
	for _, v := range commands {
		err := ExecuteSingle(v)
		errors = append(errors, err)
		if err != nil {
			return errors
		}
	}
	return errors
}

func ExecuteMultipleSkipError(commands []string) []error {
	errors := []error{}
	for _, v := range commands {
		errors = append(errors, ExecuteSingle(v))
	}
	return errors
}
