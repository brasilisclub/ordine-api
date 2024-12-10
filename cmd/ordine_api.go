package cmd

import (
	"ordine-api/config"
	docs "ordine-api/docs"
	"ordine-api/pkg/migrations"
	"ordine-api/pkg/routes"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// @title           Swagger Ordine API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
var ordineCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your application",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		migrations.Migrate()
		r := gin.Default()
		r.Use(config.CorsMiddleware())

		routes.Load(r)
		//docs.SwaggerInfo.BasePath = "/api/v1"
		docs.SwaggerInfo.BasePath = "/"
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
		r.Run()
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

	rootCmd.AddCommand(migrateCmd)
	// root command block
	rootCmd.AddCommand(ordineCmd)
}
