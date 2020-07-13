package main

import (
	"context"
	"log"

	"github.com/cvhariharan/obliwonk/actions"
	"github.com/cvhariharan/obliwonk/config"
	"github.com/cvhariharan/obliwonk/providers"
	"github.com/google/go-github/github"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"golang.org/x/oauth2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var config config.Config
	err = envconfig.Process("obliwonk", &config)

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.GithubToken},
	)

	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	p := providers.NewMathProvider(config)

	_, err = actions.CreateRepoIfNew(ctx, client, config)
	if err != nil {
		log.Fatal(err)
	}

	_, err = actions.UpdateReadMe(ctx, client, p, config)
	if err != nil {
		log.Fatal(err)
	}
}
