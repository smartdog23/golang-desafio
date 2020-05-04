package main

import (
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/smartdog23/golang-desafio/application/repositories"
	"github.com/smartdog23/golang-desafio/application/usecases"
	"github.com/smartdog23/golang-desafio/framework/pb"
	"github.com/smartdog23/golang-desafio/framework/servers"
	"github.com/smartdog23/golang-desafio/framework/utils"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

var db *gorm.DB

func main()  {

	db = utils.ConnectDB()
	db.LogMode(true)

	port := flag.Int("port", 0, "Choose the server port")
	flag.Parse()
	log.Printf("start server on port %d", *port)

	userServer := setUpUserServer()

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userServer)
	reflection.Register(grpcServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}

}

func setUpUserServer() *servers.UserServer {
	userServer := servers.NewUserServer()
	userRepository := repositories.UserRepositoryDb{Db:db}
	userServer.UserUseCase = usecases.UserUseCase{UserRepository: userRepository}
	return userServer
}