/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// getUsersCmd represents the getUsers command
var getUsersCmd = &cobra.Command{
	Use:   "getUsers",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := os.ReadFile("data.json")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(data))
	},
}

func init() {
	rootCmd.AddCommand(getUsersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getUsersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getUsersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
