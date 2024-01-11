package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	ZuiShell "github.com/chauhoangminhnguyen/material-zui-go/shell"
	ZuiS "github.com/chauhoangminhnguyen/material-zui-go/string"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if err := doAction(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
	// doAction("push")
}

func doAction(action string) error {
	action = strings.TrimSuffix(action, "\n")
	switch action {
	case "push":
		return push()
	}
	return nil
}

func push() error {
	output, _ := ZuiShell.ExecuteWithOutput("git describe --tags --abbrev=0")
	output = strings.TrimSuffix(output, "\n")
	newVersion := "v" + ZuiS.UpgradeVersion(output)
	// pushCommand := fmt.Sprintf("git tag %s && git push && git push origin %s", newVersion, newVersion)
	pushCommand := fmt.Sprintf("git tag %s", newVersion)
	fmt.Println(pushCommand)
	errs := ZuiShell.Execute(pushCommand)
	return errs[0]
}

// func execute(command string) error {
// 	command = strings.TrimSuffix(command, "\n")
// 	args := strings.Split(command, " ")
// 	cmd := exec.Command(args[0], args[1:]...)
// 	cmd.Stderr = os.Stderr
// 	cmd.Stdout = os.Stdout
// 	return cmd.Run()
// }

// func executeWithOutput(command string) (string, error) {
// 	command = strings.TrimSuffix(command, "\n")
// 	args := strings.Split(command, " ")
// 	cmd := exec.Command(args[0], args[1:]...)
// 	cmd.Stderr = os.Stderr

// 	output, err := cmd.Output()
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(output), nil
// }

// func execInput(input string) error {
// 	input = strings.TrimSuffix(input, "\n")

// 	cmd := exec.Command(input)

// 	cmd.Stderr = os.Stderr
// 	cmd.Stdout = os.Stdout

// 	return cmd.Run()
// }
