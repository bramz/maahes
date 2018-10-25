package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	URL = "https://api.urbandictionary.com/v0/define?term="
)

type DefineCmd struct {
}

type SearchResult struct {
	Tags    []string
	Results []Results `json:"list"`
}

type Results struct {
	Author     string
	Definition string
	Example    string
}

func (d DefineCmd) Handle(content []string) string {
	var strsearch string
	last := content[len(content)-1]

	if _, err := strconv.Atoi(last); err == nil {
		content = content[:len(content)-1]
	}

	strsearch = strings.Join(content[1:], " ")
	response, err := http.Get(URL + url.QueryEscape(strsearch))
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		output := "no results."
		return output
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	result := &SearchResult{}

	err = json.Unmarshal(data, result)
	if err != nil {
		fmt.Println(err)
	}

	results := result.Results

	if len(results) <= 0 {
		output := "no results were found."
		return output
	}

	if len(last) >= 2 {
		output := strsearch + " [0/" + strconv.Itoa(len(results)-1) + "]: " + results[0].Definition
		return output
	}

	if n, err := strconv.Atoi(last); err == nil {
		output := strsearch + " [" + strconv.Itoa(n) + "/" + strconv.Itoa(len(results)-1) + "]: " + results[n].Definition
		return output
	} else {
		output := strsearch + " [0/" + strconv.Itoa(len(results)-1) + "]: " + results[0].Definition
		return output
	}
}
