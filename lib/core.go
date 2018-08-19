package lib

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func StartSession() {
	discord, err := discordgo.New("Bot <yourtoken>")
	if err != nil {
		fmt.Println(err)
		return
	}

	discord.AddHandler(Commands)

	err = discord.Open()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer discord.Close()

	<-make(chan struct{})
}

func Commands(session *discordgo.Session, message *discordgo.MessageCreate) {
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

	theo, err := ioutil.ReadFile("data/theo.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := strings.Split(string(theo), "\n")

	if content == "!theo" {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		i := r.Intn(len(lines) - 1)
		line := lines[i]
		session.ChannelMessageSend(message.ChannelID, line)
	}

	fmt.Printf("Message: %+v\n", message.Message)
}
