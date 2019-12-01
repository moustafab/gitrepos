// Copyright © 2018 Moustafa Baiou <mbaiou_1@yahoo.com>

package cmd

import (
	"context"
	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
	"log"
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
		getRepositoriesAndOutput(GithubHost{cmd.Name()}, argOwner, argIsOrg, argShowCount, argAccessToken)
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

func (gh GithubHost) getRepoNames(owner string, accessToken string, ownerType OwnerType) []string {
	var repos []string
	// get all pages of results
	allRepos := gh.queryApi(owner, accessToken, ownerType)
	if allRepos != nil {
		repos = gh.parseRepoNames(allRepos)
	}
	return repos
}

func enableSecurityAlerts(client *github.Client, owner string, allRepos []*github.Repository) {
	for _, repository := range allRepos {
		var repositoryName = *repository.Name
		_, githubError := client.Repositories.EnableVulnerabilityAlerts(context.Background(), owner, repositoryName)
		if githubError != nil {
			log.Printf("Something went wrong with %v\n", repositoryName)
			continue
		}
		log.Printf("%v successfully enabled vulnerability alerts!\n", repositoryName)
	}
}

func getGithubRepos(client *github.Client, owner string, ownerType OwnerType) []*github.Repository {
	ctx := context.Background()
	globalOptions := github.ListOptions{PerPage: 100}

	if ownerType == user {
		options := &github.RepositoryListOptions{
			ListOptions: globalOptions,
		}
		var allRepos []*github.Repository
		for {
			pageOfRepos, resp, githubError := client.Repositories.List(ctx, owner, options)
			if githubError != nil {
				return nil
			}
			allRepos = append(allRepos, pageOfRepos...)
			if resp.NextPage == 0 {
				println("no more pages to go through")
				break
			}
			options.Page = resp.NextPage
		}
		return allRepos
	}
	options := &github.RepositoryListByOrgOptions{
		ListOptions: globalOptions,
	}
	var allRepos []*github.Repository
	for {
		pageOfRepos, resp, githubError := client.Repositories.ListByOrg(ctx, owner, options)
		if githubError != nil {
			return nil
		}
		allRepos = append(allRepos, pageOfRepos...)
		if resp.NextPage == 0 {
			break
		}
		options.Page = resp.NextPage
	}
	enableSecurityAlerts(client, owner, allRepos)
	return allRepos
}

func (gh GithubHost) queryApi(owner string, accessToken string, ownerType OwnerType) interface{} {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tokenClient := oauth2.NewClient(ctx, ts)
	if len(accessToken) == 0 {
		tokenClient = nil
	}

	client := github.NewClient(tokenClient)
	allRepos := getGithubRepos(client, owner, ownerType)
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
