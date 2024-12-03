package cmd

import (
	"ordine-api/pkg/routes"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var ordineCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your application",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		r := gin.Default()
		routes.ProductRoutes(r)
		routes.OrdineRoutes(r)

		r.Run()
	},
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "A brief description of your application",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {

	rootCmd.AddCommand(migrateCmd)
	// root command block
	rootCmd.AddCommand(ordineCmd)
}
