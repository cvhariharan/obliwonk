package providers

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/cvhariharan/obliwonk/config"
)

type mathProvider struct {
	url string
}

func NewMathProvider(config config.Config) Provider {
	return &mathProvider{
		url: config.MathProviderUrl,
	}
}

func (m *mathProvider) GetContent() ([]byte, error) {
	resp, err := http.Get(m.url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return r, nil
}
