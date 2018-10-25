package lib

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bramz/maahes/lib/commands"
	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
)

func StartSession(token string) {
	discord, err := discordgo.New("Bot " + token)
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

	// handle signal calls
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	defer discord.Close()
	//<-make(chan struct{})
}

func Parser(session *discordgo.Session, message *discordgo.MessageCreate) {
	valid := []string{"theo", "define", "say"}
	cmds := map[string]Cmd{
		"theo":   commands.TheoCmd{},
		"define": commands.DefineCmd{},
		"say":    commands.SayCmd{},
	}

	if string(message.Content[0]) == "!" {
		content := strings.Split(string(message.Content[1:]), " ")
		for v := range valid {
			if content[0] == valid[v] {
				out := cmds[content[0]]
				session.ChannelMessageSend(message.ChannelID, out.Handle(content))
			} else {
				// do nothing
			}
		}
	}

	// colorize terminal output
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()

	fmt.Printf("#%s <%s>: %s\n", green(message.ChannelID), red(message.Author), blue(message.Content))
}
