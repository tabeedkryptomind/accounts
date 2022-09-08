package store

import(
	 "github.com/form3tech-oss/interview-accountapi/models"
	 "github.com/form3tech-oss/interview-accountapi/cmd"
)

type store struct {}

func (s *store)Create(ac models.AccountData)error{
	tx:= cmd.Client.Create(&ac)
	return tx.Error
}

func (s *store)Fetch()([]models.AccountData, error){
		var ac []models.AccountData
		tx:= cmd.Client.Find(&ac)
		return ac, tx.Error
}

func (s *store)Delete(int64)(models.AccountData, error){
	panic("not implemented")
}