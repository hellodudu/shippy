package main

import (
	"log"

	pbVesl "github.com/hellodudu/shippy/proto/vessel"
	"github.com/hellodudu/shippy/vessel-service/handle"
	"github.com/micro/go-micro"
)

var filename = "vessels.json"

func main() {

	// new micro service
	srv := micro.NewService(micro.Name("shippy.service.vessel"))
	srv.Init()

	h, err := handle.NewVeslSrvHandler()
	if err != nil {
		log.Fatalf("failed to call NewConsSrvHandler: %v", err)
	}

	defer h.Close()

	pbVesl.RegisterVesselServiceHandler(srv.Server(), h)

	if err := srv.Run(); err != nil {
		log.Fatalln("failed to serve:", err)
	}
}
