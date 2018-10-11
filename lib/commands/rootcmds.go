package commands

type RootCmds struct {
}

func (r RootCmds) Handle(content []string) string {
	var quitmsg string
	if content[0] == "quit" {
		quitmsg = "Meow, goodbye master!"
	}
	return quitmsg
}
