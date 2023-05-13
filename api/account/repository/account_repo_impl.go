package repo

import (
	"fmt"
	"surprise/account"
	"surprise/model"

	"gorm.io/gorm"
)

type AccountRepoImpl struct {
	DB *gorm.DB
}

func NewAccountRepo(DB *gorm.DB) account.AccountRepo {
	return &AccountRepoImpl{DB}
}

func (ar *AccountRepoImpl) Create(account *model.Account) (*model.Account, error) {
	err := ar.DB.Save(&account).Error
	if err != nil {
		fmt.Printf("[AccountRepoImpl.Create] error execute query %v \n", err)
		return nil, fmt.Errorf("failed insert data")
	}
	return account, nil
}