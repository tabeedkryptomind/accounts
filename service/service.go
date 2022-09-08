package service

import (
	"github.com/form3tech-oss/interview-accountapi/models"
	"github.com/form3tech-oss/interview-accountapi/repository"
)

type AccountService struct {
	accRepo repository.Account
}

func (a * AccountService)CreateAccount(ac models.AccountData)(models.AccountData,  error) {
	err := a.accRepo.Create(ac)
	return ac, err
}

func (a * AccountService)FetchAccount()([]models.AccountData, error) {
	ac , err := a.accRepo.Fetch()
	return ac, err
}

func (a * AccountService)DeleteAccount(int64 )(models.AccountData, error){
	panic("not implemented")	
}
