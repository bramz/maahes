package lib

type Command struct {
	trigger string
	handler CmdHandler
}

type CmdHandler interface {
	Name() string
	Register(map[string]*Command)
}

func AddCmd(ch CmdHandler) *Command {
	return &Command{
		trigger: string(ch.Name()),
		handler: ch,
	}
}

func (c *Command) Register(cm map[string]*Command) {
	c.handler.Register(cm)
	return
}

func (c *Command) Name() string {
	return string(c.trigger)
}
