package actions

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/cvhariharan/obliwonk/config"
	"github.com/cvhariharan/obliwonk/providers"
	"github.com/google/go-github/github"
)

func UpdateReadMe(ctx context.Context, client *github.Client, p providers.Provider, config config.Config) (*github.RepositoryContentResponse, error) {
	fetchedReadMe, resp, err := client.Repositories.GetReadme(ctx, config.Username, config.Username, nil)
	if err != nil {
		return nil, err
	}

	// Providers provide the content to be added to README
	content, err := p.GetContent()
	if err != nil {
		return nil, err
	}
	fmt.Println("Content: ", content)
	// Append if additional content in config
	if config.ContentFooter != "" {
		content += config.ContentFooter
	}

	readme := &github.RepositoryContentFileOptions{
		Message: github.String(config.CommitMessage),
		Content: []byte(content),
	}

	if resp.StatusCode == http.StatusNotFound {
		// Create README
		r, _, err := client.Repositories.CreateFile(ctx, config.Username, config.Username, config.Readme, readme)
		if err != nil {
			return nil, err
		}
		return r, nil
	}

	// Update the existing README
	readme.SHA = fetchedReadMe.SHA
	r, _, err := client.Repositories.UpdateFile(ctx, config.Username, config.Username, config.Readme, readme)
	if err != nil {
		log.Println(err)
	}
	return r, nil
}

func CreateRepoIfNew(ctx context.Context, client *github.Client, config config.Config) (*github.Repository, error) {
	_, resp, err := client.Repositories.Get(ctx, config.Username, config.Username)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusNotFound {
		repo := &github.Repository{
			Name:    github.String(config.Username),
			Private: github.Bool(true),
		}
		r, _, err := client.Repositories.Create(ctx, "", repo)
		if err != nil {
			return nil, err
		}
		fmt.Println(r)
		return r, nil
	}
	return nil, nil
}
