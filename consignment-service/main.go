package main

import (
	"log"

	"github.com/hellodudu/shippy/consignment-service/handle"
	pbCons "github.com/hellodudu/shippy/proto/consignment"
	"github.com/micro/go-micro"
)

func main() {
	srv := micro.NewService(
		micro.Name("shippy.service.consignment"),
	)

	srv.Init()

	h, err := handle.NewConsSrvHandler()
	if err != nil {
		log.Fatalf("failed to call NewConsSrvHandler: %v", err)
	}

	defer h.Close()

	// service
	pbCons.RegisterShippingServiceHandler(srv.Server(), h)

	if err := srv.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
