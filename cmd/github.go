// Copyright © 2018 Moustafa Baiou <mbaiou_1@yahoo.com>


package cmd

import (
	"fmt"

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
		fmt.Println("github called")
	},
}

func init() {
	rootCmd.AddCommand(githubCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// githubCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// githubCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
