# Obliwonk
[![Blog Link](https://img.shields.io/badge/post-link-blue)](https://blog.cvhariharan.me/projects/2020/07/12/obliwonk/)  

**Obliwonk** is a slightly over-engineered profile README updater. Profile README is a really cool feature that allows you to add a README to your GitHub profile. You simply create a repository with your username as the repo name and add a README to it.

**Obliwonk** automates this and provides the notion of *providers*. *Providers* basically provide an abstraction over any content provider (for eg. APIs). It is just a simple interface.
```go
type Provider interface {
	GetContent() ([]byte, error)
}
```
Two providers are already included in the box, joke and math facts provider. There is also a utility function to randomly choose a provider.

### Sample Config
Add `.env` to the project dir with the following env keys
```
OBLIWONK_GITHUB_TOKEN=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
OBLIWONK_USERNAME=<github_username>
OBLIWONK_README=README.md
OBLIWONK_REPO_PRIVATE=true
OBLIWONK_COMMIT_MESSAGE=Updated via Obliwonk
OBLIWONK_MATH_PROVIDER_URL=http://numbersapi.com/random/math
OBLIWONK_JOKE_PROVIDER_URL=https://official-joke-api.appspot.com/random_joke
```
Here the `OBLIWONK_GITHUB_TOKEN` is a personal access token. The default math and joke providers use the corresponding URLs, these fields already have default values and are optional.

### Instructions
Build a binary using `go build` and use crontab for scheduling. I personally use
```bash
docker build . -t obliwonk:latest
0 */2 * * * docker run obliwonk:latest
```

You can also use GitHub Actions to schedule this workflow. You can use the following workflow
```yaml
name: Go

on:
  schedule:
    - cron: 0 */2 * * *

jobs:

  build:
    name: Build
    runs-on: ubuntu-latest
    steps:

    - name: Set up Go 1.x
      uses: actions/setup-go@v2
      with:
        go-version: ^1.13
      id: go

    - name: Check out code into the Go module directory
      uses: actions/checkout@v2

    - name: Get dependencies
      run: |
        go get -v -t -d ./...
        if [ -f Gopkg.toml ]; then
            curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
            dep ensure
        fi
    - name: Build
      run: go build
    
    - name: Touch .env
      run: touch .env
    
    - name: Run obliwonk
      env:
        OBLIWONK_GITHUB_TOKEN: ${{ secrets.OBLIWONK_TOKEN }}
        OBLIWONK_USERNAME: ${{ secrets.OBLIWONK_USERNAME }}
        OBLIWONK_README: README.md 
        OBLIWONK_COMMIT_MESSAGE: ${{ secrets.OBLIWONK_COMMIT_MESSAGE }}
        OBLIWONK_MATH_PROVIDER_URL: ${{ secrets.OBLIWONK_MATH_PROVIDER_URL }}
        OBLIWONK_JOKE_PROVIDER_URL: ${{ secrets.OBLIWONK_JOKE_PROVIDER_URL }}
      run: ./obliwonk

```
Create a file `.github/workflows/go.yml`. You will also have to create the corresponding secrets. You would have to create a new personal access token and give it repo access instead of using the default workflows GitHub token. 

### TODO
- [ ] Add support for templates
- [ ] Over-engineer