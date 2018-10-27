package commands

import (
	"math/rand"
	"strings"
	"time"
)

type EightBallCmd struct {
}

func (e EightBallCmd) Handle(content []string) string {
	var answers []string
	var question string
	var output string

	answers = []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it, yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}

	question = strings.Join(content[1:], " ")

	if len(question) >= 10 {
		rand.Seed(time.Now().UnixNano())
		output = answers[rand.Intn(len(answers))]
	} else {
		output = "you must ask a question"
	}

	return output

}
