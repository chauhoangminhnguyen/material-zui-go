package mz_tmp

import (
	"os"
	"os/exec"
	"strings"
)

func executeSingle(command string) error {
	command = strings.Trim(command, "\n")
	command = strings.Trim(command, " ")
	args := strings.Split(command, " ")
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func Execute(commandInput string) []error {
	commandInput = strings.Trim(commandInput, "\n")
	commands := strings.Split(commandInput, "&&")
	return ExecuteMultiple(commands)
}

func ExecuteMultiple(command []string) []error {
	errors := []error{}
	for _, v := range command {
		err := executeSingle(v)
		errors = append(errors, err)
		if err != nil {
			return errors
		}
	}
	return errors
}

func ExecuteMultipleSkipError(command []string) []error {
	errors := []error{}
	for _, v := range command {
		errors = append(errors, executeSingle(v))
	}
	return errors
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
