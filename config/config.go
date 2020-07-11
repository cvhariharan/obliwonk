package config

type Config struct {
	GithubToken     string `split_words:"true" required="true"`
	Username        string `split_words:"true" required="true"`
	Readme          string `split_words:"true" default="README.md"`
	CommitMessage   string `split_words:"true" default="Added via Obliwonk"`
	ContentFooter   string `split_words:"true"`
	MathProviderUrl string `split_words:"true" default="http://numbersapi.com/random/math"`
	JokeProviderUrl string `split_words:"true" default="https://official-joke-api.appspot.com/random_joke"`
}
