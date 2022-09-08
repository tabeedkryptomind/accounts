package app

import (
	"net/http"

	"github.com/form3tech-oss/interview-accountapi/logs"
	"github.com/form3tech-oss/interview-accountapi/models"
	"github.com/form3tech-oss/interview-accountapi/service"
	"github.com/gin-gonic/gin"
)

var log = logs.NewLogger()

type AccountServer struct {
	srv service.AccountService
}

func (s *AccountServer) CreateAccount(ctx *gin.Context) {
	var ac models.AccountData
	log.Info("Creating account...")

	if err := ctx.ShouldBindJSON(&ac); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "invalid json",
		})

		return
	}

	resp, err := s.srv.CreateAccount(ac)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "cannot create account",
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "account": resp})
}

func (s *AccountServer) GetAccounts(ctx *gin.Context) {
	log.Info("GetAccounts")
	ac, err := s.srv.FetchAccount()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  http.NotFound,
			"message": "no data found",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "account": ac})
}

func (s *AccountServer) DeleteAccount(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"status": http.StatusNotImplemented})
}
