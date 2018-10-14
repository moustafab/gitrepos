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
	Short: "list repos at github.com",
	Long: `list repos at github.com that the current user has access to based on the owner
For example:

gitrepos github -o moustafab
gitrepos github --owner moustafab`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("github called with owner", owner)
		repos := getRepoNames(owner)
		for _, repo := range repos {
			fmt.Println(repo)
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
