package cmd

import (
	"fmt"
)

type HostType interface {
	getCommandName() string
	getRepoNames(owner string) []string
	queryApi(owner string) interface{}
	parseRepoNames(typedRepo interface{}) []string
}

func getRepositoriesAndOutput(hostType HostType, owner string, shouldShowCount bool) {
	repos := hostType.getRepoNames(owner)
	for _, repo := range repos {
		fmt.Println(repo)
	}
	if shouldShowCount {
		fmt.Println("")
		fmt.Println(fmt.Sprintf("%v has %v repositories on %v!", owner, len(repos), hostType.getCommandName()))
	}
}
