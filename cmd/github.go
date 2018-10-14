// Copyright © 2018 Moustafa Baiou <mbaiou_1@yahoo.com>

package cmd

import (
	"context"
	"fmt"
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
		repos := getRepoNames(argOwner)
		for _, repo := range repos {
			fmt.Println(repo)
		}
		if argShowCount {
			fmt.Println("")
			fmt.Println(fmt.Sprintf("%v has %v repositories on github!", argOwner, len(repos)))
		}
	},
}

func getRepoNames(owner string) []string {
	var repos []string
	// get all pages of results
	allRepos := queryAPI(owner)
	if allRepos != nil {
		repos = parseRepoNames(allRepos)
	}
	return repos
}

func queryAPI(owner string) []*github.Repository {
	client := github.NewClient(nil)
	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}
	var allRepos []*github.Repository
	for {
		pageOfRepos, resp, githubError := client.Repositories.List(context.Background(), owner, opt)
		if githubError != nil {
			return nil
		}
		allRepos = append(allRepos, pageOfRepos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return allRepos
}

func parseRepoNames(repositories []*github.Repository) []string {
	var listOfRepoNames []string
	for _, repository := range repositories {
		listOfRepoNames = append(listOfRepoNames, *repository.Name)
	}
	return listOfRepoNames
}

func init() {
	rootCmd.AddCommand(githubCmd)
}
