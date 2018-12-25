package lib

import (
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bramz/maahes/lib/commands"
	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

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
	connect, err := sql.Open("sqlite3", "data/maahes.db")
	if err != nil {
		fmt.Println(err)
	}
	db = connect

	valid := []string{"theo", "define", "say", "markov", "8ball", "qdb", }
	cmds := map[string]Cmd{
		"theo":   commands.TheoCmd{},
		"define": commands.DefineCmd{},
//		"say":    commands.SayCmd{},
//		"markov": commands.MarkovCmd{},
		"8ball":  commands.EightBallCmd{},
		"qdb":    commands.QdbCommand{db},
//		"sync":   commands.SyncCmd{},
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

	stmt, err := db.Prepare("INSERT INTO messages (content, user, discriminator, channel, server) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
	}

	_, err = stmt.Exec(message.Content, message.Author.Username, message.Author.Discriminator, message.ChannelID, message.ChannelID)
	if err != nil {
		fmt.Println(err)
	}
}
