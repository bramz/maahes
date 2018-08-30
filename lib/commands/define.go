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

func (d DefineCmd) Handle(search []string) string {
	response, err := http.Get(URL + url.QueryEscape(search[1]))
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
