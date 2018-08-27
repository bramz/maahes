package lib

import (
	"fmt"

	"github.com/bramz/maahes/lib/commands"
	"github.com/bwmarrin/discordgo"
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

	cmds := map[string]func() string{
		"test": commands.TestCmd,
		"theo": commands.TheoCmd,
		//		"quit": commands.QuitCmd,
	}

	if string(content[0]) == "!" {
		trigger := content[1:]
		return
	}

	fmt.Printf("Message: %+v\n", message.Message)
}
