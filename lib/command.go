package lib

type Command struct {
	trigger string
	handler CmdHandler
}

type CmdHandler interface {
	Name() string
	Register(string, string)
}

func CmdRegister(ch CmdHandler) *Command {
	return &Command{
		trigger: string(ch.Name()),
		handler: ch,
	}
}

func (c *Command) Register(ct string) {
	c.handler.Register(ct, c.trigger)
}
