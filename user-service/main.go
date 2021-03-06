package main

import (
	"flag"
	"log"
	"os"

	pbUser "github.com/hellodudu/shippy/proto/user"
	"github.com/hellodudu/shippy/user-service/handle"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/broker/grpc"
)

func main() {

	dbHost := flag.String("DB_HOST", "root:@(localhost:3306)/db_shippy", "helpful string")
	flag.Parse()
	os.Setenv("DB_HOST", *dbHost)

	// new micro service
	b := grpc.NewBroker()
	srv := micro.NewService(
		micro.Name("shippy.service.user"),
		micro.Broker(b),
	)
	srv.Init()

	h, err := handle.NewUserSrvHandler(srv)
	if err != nil {
		log.Fatalf("failed to call NewUserSrvHandler: %v", err)
	}

	defer h.Close()

	pbUser.RegisterUserServiceHandler(srv.Server(), h)

	if err := srv.Run(); err != nil {
		log.Fatalln("failed to serve:", err)
	}
}
