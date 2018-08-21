package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	trigger string
	mid     string
)

func main() {
	d, err := discordgo.New("Bot <yourtoken")

	if err != nil {
		fmt.Println("failed to create discord session", err)
		os.Exit(1)
	}

	mihos, err := d.User("@me")

	if err != nil {
		fmt.Println("failed to access account", err)
		os.Exit(1)
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

	theo, err := ioutil.ReadFile("data/theo.txt")
	if err != nil {
		fmt.Print(err)
	}

	lines := strings.Split(string(theo), "\n")

	if content == "!theo" {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		i := r.Intn(len(lines) - 1)
		line := lines[i]
		d.ChannelMessageSend(msg.ChannelID, line)
	}

	fmt.Printf("Message: %+v\n", msg.Message)
}
