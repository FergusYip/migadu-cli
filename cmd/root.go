/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var configFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "migadu-cli",
	Short: "A CLI for Migadu",
	Long:  `This application is a tool to manage mailboxes, identities, aliases, and rewrites for your Migadu account.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.migadu-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/.migadu.json)")
}

func initConfig() {
	if configFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName(".migadu")
		viper.SetConfigType("json")
		viper.AddConfigPath("$HOME")
		viper.AddConfigPath(".")
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Println("Missing configuration. Run migadu-cli setup")

			os.Exit(1) // TODO Be smart and run it for the user
		} else {
			// Config file was found but another error was produced
			fmt.Println("Invalid configuration file: ", viper.ConfigFileUsed())
			os.Exit(1)
		}
	} else {
		fmt.Println("Using configuration file: ", viper.ConfigFileUsed())
	}
}
