package mz_string

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
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

type ZString string

// type ZString interface {
// 	string
// 	error
// }

func ZuiString(value string) *ZString {
	result := ZString(value)
	return &result
}

func (value *ZString) Get() string {
	return string(*value)
}

func (value *ZString) ToLower() *ZString {
	*value = ZString(strings.ToLower(string(*value)))
	return value
}

func (value *ZString) ToUpper() *ZString {
	*value = ZString(strings.ToUpper(string(*value)))
	return value
}

func (value *ZString) ReplaceWithRegex(expr string, repl string) *ZString {
	reg, err := regexp.Compile(expr)
	if err != nil {
		log.Error(err)
		return value
	}
	*value = ZString(reg.ReplaceAllString(value.Get(), repl))
	return value
}

func (value *ZString) ReplaceAll(src string, repl string) *ZString {
	*value = ZString(strings.ReplaceAll(value.Get(), src, repl))
	return value
}
