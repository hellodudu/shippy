package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	pbCons "github.com/hellodudu/shippy/proto/consignment"
	pbVesl "github.com/hellodudu/shippy/proto/vessel"
	micro "github.com/micro/go-micro"
)

const (
	CONSIGNMENT_FILE = "consignments.json"
	VESSEL_FILE      = "vessels.json"
)

func createConsignment(s micro.Service) {
	client := pbCons.NewShippingServiceClient("shippy.service.consignment", s.Client())

	// read consignment file
	data, err := ioutil.ReadFile(CONSIGNMENT_FILE)
	if err != nil {
		log.Fatal(err)
	}

	var consignments []*pbCons.Consignment
	err = json.Unmarshal(data, &consignments)
	if err != nil {
		log.Fatal("consignments.json file content error")
	}

	for _, v := range consignments {
		createResp, err := client.CreateConsignment(context.Background(), v)
		if err != nil {
			log.Fatalf("create consignment error: %v", err)
		}

		log.Println("CreateConsignment response: %v", createResp)
	}

	getResp, err := client.GetConsignments(context.Background(), &pbCons.GetRequest{})
	if err != nil {
		log.Fatalf("GetConsignments error: %v", err)
	}

	log.Printf("GetConsignments response : %v", getResp)
}

func createVessel(s micro.Service) {
	client := pbVesl.NewVesselServiceClient("shippy.service.vessel", s.Client())

	// read vessel file
	data, err := ioutil.ReadFile(VESSEL_FILE)
	if err != nil {
		log.Fatal(err)
	}

	var vessels []*pbVesl.Vessel
	err = json.Unmarshal(data, &vessels)
	if err != nil {
		log.Fatal("vessels.json file content error")
	}

	for _, v := range vessels {
		createResp, err := client.Create(context.Background(), v)
		if err != nil {
			log.Fatalf("create vessels error: %v", err)
		}

		log.Println("CreateVessels response: %v", createResp)
	}
}

func main() {

	service := micro.NewService(micro.Name("shippy.consignment.cli"))
	service.Init()

	createConsignment(service)
	createVessel(service)

}
