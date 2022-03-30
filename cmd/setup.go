/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"regexp"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup the Migadu CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		usernamePrompt := promptui.Prompt{
			Label: "Account Email",
			Validate: func(input string) error {

				emailPattern := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

				if !emailPattern.MatchString(input) {
					return errors.New("Invalid email address")
				}
				return nil
			},
		}

		username, err := usernamePrompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		apiKeyPrompt := promptui.Prompt{
			Label: "API Key",
			Validate: func(input string) error {
				if len(input) == 0 {
					return errors.New("API Key is required")
				}
				return nil
			},
		}

		apiKey, err := apiKeyPrompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		fmt.Printf("Account Email: %q\n", username)
		fmt.Printf("API Keu: %q\n", apiKey)

	},
}

func init() {
	rootCmd.AddCommand(setupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
