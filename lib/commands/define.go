package commands

import (
	"fmt"

	urbandict "github.com/davidscholberg/go-urbandict"
)

type DefineCmd struct {
}

func (d DefineCmd) Handle(content []string) string {
	def, err := urbandict.Define(string(content[1]))
	if err != nil {
		fmt.Println(err)
	}

	return def.Definition
}
