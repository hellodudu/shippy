package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/hellodudu/shippy/consignment-service/repo"
	pbCons "github.com/hellodudu/shippy/proto/consignment"
	pbVesl "github.com/hellodudu/shippy/proto/vessel"
	"github.com/micro/go-micro"
)

const (
	PORT = ":50051"
)

type service struct {
	r repo.IRepository
	c pbVesl.VesselServiceClient
}

func (s *service) CreateConsignment(ctx context.Context, req *pbCons.Consignment, out *pbCons.CreateConsignmentResponse) error {

	spec := &pbVesl.Specification{Weight: req.Weight}
	resp, err := s.c.FindAvailable(ctx, spec)
	if err != nil {
		return err
	}

	log.Println("vesl client call FindAvailable response:", resp)

	if len(resp.Vessels) == 0 {
		errs := fmt.Sprintf("can't find available vessel with specification:%v", spec)
		log.Println(errs)
		return errors.New(errs)
	}

	if err := s.r.Create(req); err != nil {
		return err
	}

	out.Created = true
	out.Consignment = req
	return nil
}

func (s *service) GetConsignments(ctx context.Context, _ *pbCons.GetRequest, out *pbCons.GetConsignmentsResponse) error {
	var err error
	out.Consignments, err = s.r.GetAll()
	return err
}

func main() {
	s := &service{}
	var err error
	if s.r, err = repo.NewRepository(); err != nil {
		log.Fatalf("failed to call NewRepository(): %v", err)
	}

	srv := micro.NewService(
		micro.Name("shippy.service.consignment"),
	)

	srv.Init()

	// client
	s.c = pbVesl.NewVesselServiceClient("shippy.service.vessel", nil)

	// service
	pbCons.RegisterShippingServiceHandler(srv.Server(), s)

	if err := srv.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
