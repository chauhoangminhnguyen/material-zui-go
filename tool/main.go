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
	currentTag, _ := ZuiShell.ExecuteWithOutput("git describe --tags --abbrev=0")
	currentTag = strings.TrimSuffix(currentTag, "\n")
	newVersion := "v" + ZuiS.UpgradeVersion(currentTag)
	pushCommand := fmt.Sprintf("git tag %s && git push && git push origin %s", newVersion, newVersion)
	fmt.Println(pushCommand)
	errs := ZuiShell.Execute(pushCommand)
	return errs[0]
}
