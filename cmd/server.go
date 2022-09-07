package cmd

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	app "github.com/form3tech-oss/interview-accountapi/app/server"
	"github.com/form3tech-oss/interview-accountapi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


const (
	port = ":8081"
)

var DB *gorm.DB

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
	}

	srv.Run(port)
}





func Setup(){
	host := os.Getenv("PSQL_HOST")
	pgPort := os.Getenv("PSQL_PORT")

	dsn := "host=postgresql user=root password=password dbname=interview_accountapi port=5432"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = DB.AutoMigrate(models.AccountData{})
	if err != nil {
		fmt.Println(err)
	}
}