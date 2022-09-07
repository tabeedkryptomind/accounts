package app

import (
	"net/http"

	"github.com/form3tech-oss/interview-accountapi/models"
	"github.com/form3tech-oss/interview-accountapi/service"
	"github.com/gin-gonic/gin"
)




type AccountServer struct {
	srv service.AccountService
}

func(s *AccountServer)CreateAccount(ctx * gin.Context){
	var ac models.AccountData

	if err := ctx.ShouldBindJSON(&ac); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "invalid json",
		})

		return
	}

	resp, err := s.srv.CreateAccount(ac)
	if err != nil {

	}

	ctx.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "account":resp})
}