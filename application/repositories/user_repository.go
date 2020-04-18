package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/smartdog23/golang-desafio/domain"
	"log"
)

type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
}

type UserRepositoryDb struct {
	Db *gorm.DB
}

func (repo UserRepositoryDb) Insert (user *domain.User) (*domain.User, error) {

	err := user.Prepare()

	if err != nil {
		log.Fatalf("Error during the user validation: %v", err)
		return user, err
	}

	err = repo.Db.Create(user).Error

	if err != nil {
		log.Fatalf("Error to persist user: %v", err)
		return user, err
	}

	return user, nil
}

//type UserRepositoryMemory struct {
//
//}
//
//type UserRepositoryTxt struct {
//
//}