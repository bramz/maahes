package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	URL = "https://api.urbandictionary.com/v0/define?term="
)

/*
type Response struct {
	Tags       []string `json:"tags"`
	ResultType string   `json:"result_type"`
	Results    []Result `json:"list"`
}

type Result struct {
	Definition  string `json:"definition"`
	Permalink   string `json:"permalink"`
	ThumbsUp    int    `json:"thumbs_up"`
	Author      string `json:"author"`
	Word        string `json:"word"`
	Defid       int    `json:"defid"`
	CurrentVote string `json:"current_vote"`
	Example     string `json:"example"`
	ThumbsDown  int    `json:"thumbs_down"`
}
*/
type SearchResult struct {
	Tags    []string
	Results []Results `json:"list"`
}

type Results struct {
	Author     string
	Definition string
	Example    string
}

func DefineCmd(search string) string {
	response, err := http.Get(URL + url.QueryEscape(search))
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Errorf("response was not valid: %d", response.StatusCode)
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

	first := result.Results
	def := first[0].Definition

	return def
}
