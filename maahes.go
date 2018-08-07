package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
	trigger string
	mid     string
)

func main() {
	d, err := discordgo.New("Bot <yourtoken>")

	if err != nil {
		fmt.Println("failed to create discord session", err)
	}

	mihos, err := d.User("@me")

	if err != nil {
		fmt.Println("failed to access account", err)
	}

	mid = mihos.ID
	d.AddHandler(handleCmd)
	err = d.Open()

	if err != nil {
		fmt.Println("unable to establish connection", err)
	}

	defer d.Close()

	trigger = "!"
	<-make(chan struct{})
}

func handleCmd(d *discordgo.Session, msg *discordgo.MessageCreate) {
	user := msg.Author
	if user.ID == mid || user.Bot {
		return
	}

	content := msg.Content

	if content == "!test" {
		d.ChannelMessageSend(msg.ChannelID, "Testing..")
	}

	fmt.Printf("Message: %+v\n", msg.Message)
}
