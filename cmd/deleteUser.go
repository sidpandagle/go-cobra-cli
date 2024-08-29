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

// deleteUserCmd represents the deleteUser command
var deleteUserCmd = &cobra.Command{
	Use:   "deleteUser",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			data, err := os.ReadFile("data.json")
			if err != nil {
				fmt.Println(err)
				return
			}

			var users []User
			var deleteIndex int
			json.Unmarshal(data, &users)
			for index, user := range users {
				if user.Name == args[0] {
					deleteIndex = index
				}
			}

			users = append(users[:deleteIndex], users[deleteIndex+1:]...)

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
			fmt.Println(args[0] + " deleted")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteUserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteUserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteUserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
