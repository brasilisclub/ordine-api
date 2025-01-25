/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"ordine-api/config"
	"ordine-api/pkg/http"
	"ordine-api/pkg/migrations"
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
	Use:   "entrypoint",
	Short: "All commands available of your application",
	Long:  logo + `This application default action is startt webserver interface`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(logo+"Service http started, listening port: %v and log_level: %v\n", config.Envs.PORT, config.Envs.LOG_LEVEL)
		http.GetServer().Run()
	},
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "A brief description of your application",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		migrations.Migrate()
	},
}

func init() {
	// root command block
	rootCmd.AddCommand(migrateCmd)
}

// osExit interface to allow testing
var osExit = os.Exit

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		osExit(1)
	}
}
