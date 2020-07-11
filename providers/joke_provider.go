package providers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/cvhariharan/obliwonk/config"
)

type Joke struct {
	ID        int64  `json:"id"`
	Type      string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

type jokeProvider struct {
	url string
}

func NewJokeProvider(config config.Config) Provider {
	return &jokeProvider{
		url: config.JokeProviderUrl,
	}
}

func (j *jokeProvider) GetContent() (string, error) {
	resp, err := http.Get(j.url)
	if err != nil {
		return "", err
	}
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var joke Joke
	err = json.Unmarshal(r, &joke)
	if err != nil {
		return "", err
	}
	return joke.Setup + " ... " + joke.Punchline, nil
}
