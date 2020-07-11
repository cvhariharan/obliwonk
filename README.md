# Obliwonk
**Obliwonk** is a slightly over-engineered profile README updater. Profile README is a really cool feature that allows you to add a README to your GitHub profile. You simply create a repository with your username as the repo name and add a README to it.

**Obliwonk** automates this and provides the notion of *providers*. *Providers* basically provide an abstraction over any content provider (for eg. APIs). It is just a simple interface.
```go
type Provider interface {
	GetContent() (string, error)
}
```
Two providers are already included in the box, joke and math facts provider. There is also a utility function to randomly choose a provider.

### Sample Config
Add `.env` to the project dir with the following env keys
```
OBLIWONK_GITHUB_TOKEN=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
OBLIWONK_USERNAME=<github_username>
OBLIWONK_README=README.md
OBLIWONK_COMMIT_MESSAGE=Updated via Obliwonk
OBLIWONK_MATH_PROVIDER_URL=http://numbersapi.com/random/math
OBLIWONK_JOKE_PROVIDER_URL=https://official-joke-api.appspot.com/random_joke
```
Here the `OBLIWONK_GITHUB_TOKEN` is a personal access token. The default math and joke providers use the corresponding URLs, these fields already have default values and are optional.