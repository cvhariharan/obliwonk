package actions

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/cvhariharan/obliwonk/config"
	"github.com/cvhariharan/obliwonk/providers"
	"github.com/google/go-github/github"
)

func UpdateReadMe(ctx context.Context, client *github.Client, p providers.Provider, config config.Config) (*github.RepositoryContentResponse, error) {
	fetchedReadMe, resp, readMeErr := client.Repositories.GetReadme(ctx, config.Username, config.Username, nil)
	// Providers provide the content to be added to README
	c, err := p.GetContent()
	if err != nil {
		return nil, err
	}
	fmt.Println("Content: ", string(c))
	// Append if additional content in config
	var content bytes.Buffer
	content.Write(c)

	if config.ContentFooter != "" {
		content.Write([]byte(config.ContentFooter))
	}

	readme := &github.RepositoryContentFileOptions{
		Message: github.String(config.CommitMessage),
		Content: content.Bytes(),
	}

	// Check status code instead of error handling
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
		return nil, err
	}
	return r, readMeErr
}

func CreateRepoIfNew(ctx context.Context, client *github.Client, config config.Config) (*github.Repository, error) {
	_, resp, err := client.Repositories.Get(ctx, config.Username, config.Username)
	if resp.StatusCode == http.StatusNotFound {
		repo := &github.Repository{
			Name:    github.String(config.Username),
			Private: github.Bool(config.RepoPrivate),
		}
		r, _, err := client.Repositories.Create(ctx, "", repo)
		if err != nil {
			return nil, err
		}
		return r, nil
	}

	return nil, err
}
