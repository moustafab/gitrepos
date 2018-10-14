// Copyright © 2018 Moustafa Baiou <mbaiou_1@yahoo.com>

package cmd

import (
	"context"
	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
)

// githubCmd represents the github command
var githubCmd = &cobra.Command{
	Use:   "github",
	Short: "list owner's repos at github.com",
	Long: `list owner's repos at github.com that the current user has access to
For example:

gitrepos github -o moustafab
gitrepos github --owner moustafab`,
	Run: func(cmd *cobra.Command, args []string) {
		getRepositoriesAndOutput(GithubHost{cmd.Name()}, argOwner, argShowCount)
	},
}

func init() {
	rootCmd.AddCommand(githubCmd)
}

type GithubHost struct {
	name string
}

func (gh GithubHost) getCommandName() string {
	return gh.name
}

func (gh GithubHost) getRepoNames(owner string) []string {
	var repos []string
	// get all pages of results
	allRepos := gh.queryApi(owner)
	if allRepos != nil {
		repos = gh.parseRepoNames(allRepos)
	}
	return repos
}

func (gh GithubHost) queryApi(owner string) interface{} {
	client := github.NewClient(nil)
	options := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}
	var allRepos []*github.Repository
	for {
		pageOfRepos, resp, githubError := client.Repositories.List(context.Background(), owner, options)
		if githubError != nil {
			return nil
		}
		allRepos = append(allRepos, pageOfRepos...)
		if resp.NextPage == 0 {
			break
		}
		options.Page = resp.NextPage
	}
	return allRepos
}

func (gh GithubHost) parseRepoNames(unTypedRepositories interface{}) []string {
	var listOfRepoNames []string
	repositories, ok := unTypedRepositories.([]*github.Repository)
	if ok {
		for _, repository := range repositories {
			listOfRepoNames = append(listOfRepoNames, *repository.Name)
		}
	}
	return listOfRepoNames
}
