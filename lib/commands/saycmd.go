package commands

import "strings"

type SayCmd struct {
}

func (s SayCmd) Handle(content []string) string {
	var msg string
	if content[0] == "say" {
		msg = strings.Join(content[1:], " ")
	}
	return msg
}
