package main

import (
	"fmt"
	"github.com/smartdog23/golang-desafio/application/repositories"
	"github.com/smartdog23/golang-desafio/application/usecases"
	"github.com/smartdog23/golang-desafio/domain"
	"github.com/smartdog23/golang-desafio/framework/utils"
	"log"
)

func main()  {

	db := utils.ConnectDB()

	user := domain.User{
		Name:     "Valdir",
		Email:    "valdir1@gmail.com",
		Password: "secret",
	}

	userRepo := repositories.UserRepositoryDb{Db:db}

	userUseCase := usecases.UserUseCase{userRepo}

	result, err := userUseCase.Create(&user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result.Name, result.Email, result.Token)

}
