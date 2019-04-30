package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	pbVesl "github.com/hellodudu/shippy/proto/vessel"
	"github.com/micro/go-micro"
)

var filename = "vessels.json"

type vesselService struct {
	Vessels []*pbVesl.Vessel
}

func (v *vesselService) FindAvailable(ctx context.Context, spec *pbVesl.Specification, resp *pbVesl.Response) error {
	for _, val := range v.Vessels {
		if val.MaxWeight >= spec.Weight {
			resp.Vessels = append(resp.Vessels, val)
		}
	}
	log.Println("call FindAvailable spec:", spec, ", resp:", resp)
	return nil
}

func main() {
	v := &vesselService{}

	// parse from json
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("read ", filename, " failed:", err)
	}

	if err := json.Unmarshal(b, v); err != nil {
		log.Fatalln("json unmarshal failed:", err)
	}

	// new micro service
	srv := micro.NewService(micro.Name("shippy.service.vessel"))
	srv.Init()
	pbVesl.RegisterVesselServiceHandler(srv.Server(), v)

	if err := srv.Run(); err != nil {
		log.Fatalln("failed to serve:", err)
	}
}
