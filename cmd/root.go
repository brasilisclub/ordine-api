/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var logo = `
Ordine API
======================================================
                _ _                       _____ _____
               | (_)                /\   |  __ \_   _|
   ___  _ __ __| |_ _ __   ___     /  \  | |__) || |
  / _ \| '__/ _| | | '_ \ / _ \   / /\ \ |  ___/ | |
 | (_) | | | (_| | | | | |  __/  / ____ \| |    _| |_
  \___/|_|  \__,_|_|_| |_|\___| /_/    \_\_|   |_____|
======================================================
🚀 Powered by brasilis for Monsters! 🚀
`

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ordine-api",
	Short: "",
	Long:  logo,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
