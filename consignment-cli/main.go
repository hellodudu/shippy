package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pbCons "github.com/hellodudu/shippy/proto/consignment"
	pbUser "github.com/hellodudu/shippy/proto/user"
	pbVesl "github.com/hellodudu/shippy/proto/vessel"
	"github.com/micro/cli"
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

	var consignments struct {
		Cons []*pbCons.Consignment `json:"consignments"`
	}

	err = json.Unmarshal(data, &consignments)
	if err != nil {
		log.Fatal("consignments.json file content error:", err)
	}

	for _, v := range consignments.Cons {
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

	var vessels struct {
		Vs []*pbVesl.Vessel `json:"vessels"`
	}

	err = json.Unmarshal(data, &vessels)
	if err != nil {
		log.Fatal("vessels.json file content error", err)
	}

	for _, v := range vessels.Vs {
		createResp, err := client.Create(context.Background(), v)
		if err != nil {
			log.Fatalf("create vessels error: %v", err)
		}

		log.Println("CreateVessels response: %v", createResp)
	}
}

func createUser() {
	client := pbUser.NewUserServiceClient("shippy.service.user", nil)
	service := micro.NewService()

	service.Init(
		micro.Action(func(c *cli.Context) {
			name := "dudu"
			email := "hellodudu86@gmail.com"
			password := "123qwe"
			company := "amazing"

			r, err := client.Create(context.TODO(), &pbUser.User{
				Id:       "1001",
				Name:     name,
				Email:    email,
				Password: password,
				Company:  company,
			})
			if err != nil {
				log.Fatalf("Could not create: %v", err)
			}
			log.Printf("Created: %v", r.User.Id)

			getAll, err := client.GetAll(context.Background(), &pbUser.Request{})
			if err != nil {
				log.Fatalf("Could not list users: %v", err)
			}
			for _, v := range getAll.Users {
				log.Println(v)
			}

			os.Exit(0)
		}),
	)

	if err := service.Run(); err != nil {
		log.Println(err)
	}
}

func main() {

	service := micro.NewService(micro.Name("shippy.consignment.cli"))
	service.Init()

	// createVessel(service)
	// createConsignment(service)
	createUser()
}
