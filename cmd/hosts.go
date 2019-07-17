package cmd

import (
	"fmt"
)

type OwnerType string

const (
	org  OwnerType = "ORG"
	user OwnerType = "USER"
)

type HostType interface {
	getCommandName() string
	getRepoNames(owner string, accessToken string, ownerType OwnerType) []string
	queryApi(owner string, accessToken string, ownerType OwnerType) interface{}
	parseRepoNames(typedRepo interface{}) []string
}

func getRepositoriesAndOutput(hostType HostType, owner string, isOrg bool, shouldShowCount bool, accessToken string) {
	ownerType := user
	if isOrg {
		ownerType = org
	}
	repos := hostType.getRepoNames(owner, accessToken, ownerType)
	for _, repo := range repos {
		fmt.Println(repo)
	}
	if shouldShowCount {
		fmt.Println("")
		fmt.Println(fmt.Sprintf("%v has %v repositories on %v!", owner, len(repos), hostType.getCommandName()))
	}
}
