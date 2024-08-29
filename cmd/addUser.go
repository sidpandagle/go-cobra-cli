/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// addUserCmd represents the addUser command
var addUserCmd = &cobra.Command{
	Use:   "addUser",
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

		var users []User
		json.Unmarshal(data, &users)

		if len(args) == 1 {
			user := User{Name: args[0]}
			users = append(users, user)
		}

		file, err := os.Open("data.json")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		jdata, jerr := json.Marshal(users)
		if jerr != nil {
			return
		}
		werr := os.WriteFile("data.json", jdata, 0644)
		if werr != nil {
			return
		}
		fmt.Println(args[0] + " added")
	},
}

func init() {
	rootCmd.AddCommand(addUserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addUserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addUserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
