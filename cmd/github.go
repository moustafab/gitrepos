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
		repos := getRepoNames(owner)
		for _, repo := range repos {
			fmt.Println(repo)
		}
		if showCount {
			fmt.Println("")
			fmt.Println(fmt.Sprintf("%v has %v repositories on github!", owner, len(repos)))
		}
	},
}

func getRepoNames(s string) []string {
	var repos []string
	client := github.NewClient(nil)

	fullRepositories, _, githubError := client.Repositories.List(context.Background(), owner, nil)
	if githubError != nil {
		return nil
	}

	if fullRepositories != nil {
		repos = parseRepoNames(fullRepositories)
	}

	return repos
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
