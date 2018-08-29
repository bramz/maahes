package commands

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func TheoCmd(content string) string {
	theo, err := ioutil.ReadFile("data/theo.txt")
	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Split(string(theo), "\n")

	r := rand.New(rand.NewSource(time.Now().Unix()))
	i := r.Intn(len(lines) - 1)
	line := lines[i]
	return line
}
