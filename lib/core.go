package lib

import (
	"fmt"

	"github.com/bramz/maahes/lib/commands"

	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
)

const (
	cmds := map[string]Cmd{
		"theo":   commands.TheoCmd{},
		"define": commands.DefineCmd{},
	}
)

func StartSession() {
	discord, err := discordgo.New("Bot <yourtoken>")
	if err != nil {
		fmt.Println(err)
		return
	}

	discord.AddHandler(Parser)

	err = discord.Open()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer discord.Close()

	<-make(chan struct{})
}

func Parser(session *discordgo.Session, message *discordgo.MessageCreate) {
	maahes, err := session.User("@me")
	if err != nil {
		fmt.Println(err)
		return
	}

	author := message.Author
	if author.ID == maahes.ID || author.Bot {
		return
	}

	content := message.Content

	if string(content[0]) == "!" {
		ct := content[1:]
		out := cmds[ct]
		session.ChannelMessageSend(message.ChannelID, out.Handle([]string{ct}))
	}

	// colorize terminal output
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()

	fmt.Printf("#%s <%s>: %s\n", green(message.ChannelID), red(message.Author), blue(message.Content))
}
