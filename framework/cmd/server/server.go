package main

import (
	"fmt"
	"github.com/smartdog23/golang-desafio/application/repositories"
	"github.com/smartdog23/golang-desafio/domain"
	"github.com/smartdog23/golang-desafio/framework/utils"
	"log"
)

func main()  {

	db := utils.ConnectDB()

	user := domain.User{
		Name:     "Valdir",
		Email:    "valdirpl@gmail.com",
		Password: "secret",
	}

	userRepo := repositories.UserRepositoryDb{Db:db}

	result, err := userRepo.Insert(&user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result.Name, result.Email, result.Token)

}
