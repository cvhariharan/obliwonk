package providers

import (
	"math/rand"
	"time"

	"github.com/cvhariharan/obliwonk/config"
)

type Provider interface {
	GetContent() ([]byte, error)
}

func GetRandomProvider(config config.Config) Provider {
	var allProviders []Provider
	allProviders = append(allProviders, NewMathProvider(config))
	allProviders = append(allProviders, NewJokeProvider(config))

	rand.Seed(time.Now().UnixNano())
	return allProviders[rand.Intn(len(allProviders))]
}
