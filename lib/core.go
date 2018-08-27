package lib

import (
	"fmt"

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

	commands := map[string]*Command{
		"test", TestCmd,
		"theo", TheoCmd,
		"quit", QuitCmd,
	}

	if string(content[0]) == "!" {
		trigger := content[1:]
		if trigger == commands[trigger]*Command {
			return commands[trigger]
		} else {
			fmt.Println("command does not exist")
		}
	}

	fmt.Printf("Message: %+v\n", message.Message)
}
