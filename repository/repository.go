package repository

import "github.com/form3tech-oss/interview-accountapi/models"


type Account interface {
	Create(models.AccountData)(models.AccountData, error)

	Fetch()([]models.AccountData, error)
	
	Delete(int64)(models.AccountData, error)
}