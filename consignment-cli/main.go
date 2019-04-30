package main

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

	pbCons "github.com/hellodudu/shippy/proto/consignment"
	micro "github.com/micro/go-micro"
)

const (
	DEFAULT_INFO_FILE = "consignment.json"
)

func parseFile(fileName string) (*pbCons.Consignment, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var consignment *pbCons.Consignment
	err = json.Unmarshal(data, &consignment)
	if err != nil {
		return nil, errors.New("consignment.json file content error")
	}
	return consignment, nil
}

func main() {

	infoFile := DEFAULT_INFO_FILE
	if len(os.Args) > 1 {
		infoFile = os.Args[1]
	}

	service := micro.NewService(micro.Name("shippy.consignment.cli"))
	service.Init()

	client := pbCons.NewShippingServiceClient("shippy.service.consignment", service.Client())

	consignment, err := parseFile(infoFile)
	if err != nil {
		log.Fatalf("parse info file error: %v", err)
	}

	createResp, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("create consignment error: %v", err)
	}

	log.Printf("CreateConsignment response: %v", createResp)

	getResp, err := client.GetConsignments(context.Background(), &pbCons.GetRequest{})
	if err != nil {
		log.Fatalf("get consignment error: %v", err)
	}

	log.Printf("GetConsignment response : %v", getResp)
}
