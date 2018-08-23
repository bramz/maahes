package lib

type Command struct {
	trigger string
	handler CmdHandler
}

type CmdHandler interface {
	Name() string
	Register(string)
}

func AddCmd(ch CmdHandler) *Command {
	return &Command{
		trigger: string(ch.Name()),
		handler: ch,
	}
}

func (c *Command) Register(ct string) {
	c.handler.Register(ct)
	c.trigger = c.handler.Load(ct)
}

func (c *Command) Name() string {
	return string(c.trigger)
}

type Commands struct {
	name map[string]string
}

func NewCommand() *Commands {
	return &Commands{
		name: map[string]string
	}
}

func (cmds *Commands) Register(ct map[string]string) string {
	cmds.name[ct] = ct
	return cmds.name[ct]
}


