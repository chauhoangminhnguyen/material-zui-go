package mz_string

import (
	"fmt"
	"regexp"
	"strconv"
)

func replaceWithRegex(expr string, src string, repl string) (string, error) {
	reg, err := regexp.Compile(expr)
	if err != nil {
		return "", err
	}
	return reg.ReplaceAllString(src, repl), nil
}

func ConvertToVersion(num int) string {
	major := num / 100
	minor := (num % 100) / 10
	patch := num % 10
	return fmt.Sprintf("%d.%d.%d", major, minor, patch)
}

func UpgradeVersion(version string) string {
	textVersion, _ := replaceWithRegex("\\D+", version, "")
	versionNumber, _ := strconv.Atoi(textVersion)
	return ConvertToVersion(versionNumber + 1)
}
