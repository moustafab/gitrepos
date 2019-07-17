// Copyright © 2018 Moustafa Baiou <mbaiou_1@yahoo.com>

package cmd

import (
	"fmt"
	"github.com/crossid/bitbucket-golang-api"
	"github.com/spf13/cobra"
)

// bitbucketCmd represents the bitbucket command
var bitbucketCmd = &cobra.Command{
	Use:   "bitbucket",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		getRepositoriesAndOutput(BitbucketHost{cmd.Name()}, argOwner, argIsOrg, argShowCount, argAccessToken)
	},
}

func init() {
	rootCmd.AddCommand(bitbucketCmd)
}

type BitbucketHost struct {
	name string
}

func (bb BitbucketHost) getCommandName() string {
	return bb.name
}

func (bb BitbucketHost) getRepoNames(owner string, accessToken string, ownerType OwnerType) []string {
	var repos []string
	// get all pages of results
	allRepos := bb.queryApi(owner, accessToken, ownerType)
	if allRepos != nil {
		repos = bb.parseRepoNames(allRepos)
	}
	return repos
}

func (bb BitbucketHost) queryApi(owner string, accessToken string, ownerType OwnerType) interface{} {
	client := bitbucket.NewV2()
	options := bitbucket.ListReposOpts{
		Pagelen: 500,
	}
	var allRepos []map[string]interface{}
	response, bitbucketError := client.Repositories.ListPublic(options)
	if bitbucketError != nil {
		fmt.Println("Error accessing bitbucket", bitbucketError)
		return nil
	}
	allRepos = append(allRepos, response.Values...)
	return allRepos
}

func (bb BitbucketHost) parseRepoNames(unTypedRepositories interface{}) []string {
	var listOfRepoNames []string
	repositories, ok := unTypedRepositories.([]map[string]interface{})
	if ok {
		for _, repository := range repositories {
			listOfRepoNames = append(listOfRepoNames, repository["full_name"].(string))
		}
	}
	return listOfRepoNames
}
