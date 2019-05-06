package main

import (
	"log"
	"os"

	pbUser "github.com/hellodudu/shippy/proto/user"
	"github.com/hellodudu/shippy/user-service/handle"
	"github.com/micro/go-micro"
)

func main() {

	os.Setenv("DB_HOST", "root:@(localhost:3306)/db_shippy")

	// new micro service
	srv := micro.NewService(micro.Name("shippy.service.user"))
	srv.Init()

	h, err := handle.NewUserSrvHandler()
	if err != nil {
		log.Fatalf("failed to call NewUserSrvHandler: %v", err)
	}

	defer h.Close()

	pbUser.RegisterUserServiceHandler(srv.Server(), h)

	if err := srv.Run(); err != nil {
		log.Fatalln("failed to serve:", err)
	}
}