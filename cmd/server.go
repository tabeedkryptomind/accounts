package cmd

import (

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	app "github.com/form3tech-oss/interview-accountapi/app/server"
	"github.com/form3tech-oss/interview-accountapi/logs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


const (
	port = ":8080"
)

var Client *gorm.DB

var accountsCmd = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
		serveUser()
	},
}

func serveUser(){
	srv  := gin.New()
	
	srv.Use(gin.Recovery(), gin.Logger())

	Setup()
	accounts := app.AccountServer{}

	v1 := srv.Group("/v1")
	{
		v1.GET("/",func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"status": "OK", "message": "Healthcheck successful"})
		})

		v1.POST("/create",accounts.CreateAccount)
		v1.GET("/ac", accounts.GetAccounts)
		v1.DELETE("/:id", accounts.DeleteAccount)
	}

	srv.Run(port)
}


func Setup(){
	log := logs.NewLogger()
	var err error

	dsn := "host=postgresql-db user=admin password=test dbname=interview_accountapi port=5432 sslmode=disable"

	Client, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Infof("cannot connect to postgres: %v", err)
	}

	log.Info("DATABASE Connected")
}