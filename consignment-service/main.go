package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	pbCons "github.com/hellodudu/shippy/proto/consignment"
	pbVesl "github.com/hellodudu/shippy/proto/vessel"
	"github.com/micro/go-micro"
)

const (
	PORT = ":50051"
)

type IRepository interface {
	Create(consignment *pbCons.Consignment) (*pbCons.Consignment, error)
	GetAll() []*pbCons.Consignment
}

type Repository struct {
	consignments []*pbCons.Consignment
}

func (repo *Repository) Create(consignment *pbCons.Consignment) (*pbCons.Consignment, error) {
	repo.consignments = append(repo.consignments, consignment)
	return consignment, nil
}

func (repo *Repository) GetAll() []*pbCons.Consignment {
	return repo.consignments
}

type service struct {
	repo       IRepository
	veslClient pbVesl.VesselServiceClient
}

func (s *service) CreateConsignment(ctx context.Context, req *pbCons.Consignment, out *pbCons.CreateConsignmentResponse) error {

	spec := &pbVesl.Specification{Weight: req.Weight}
	resp, err := s.veslClient.FindAvailable(ctx, spec)
	if err != nil {
		return err
	}

	log.Println("vesl client call FindAvailable response:", resp)

	if len(resp.Vessels) == 0 {
		errs := fmt.Sprintf("can't find available vessel with specification:%v", spec)
		log.Println(errs)
		return errors.New(errs)
	}

	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	out.Created = true
	out.Consignment = consignment
	return nil
}

func (s *service) GetConsignments(ctx context.Context, _ *pbCons.GetRequest, out *pbCons.GetConsignmentsResponse) error {
	out.Consignments = s.repo.GetAll()
	return nil
}

func main() {
	s := &service{repo: &Repository{}}

	srv := micro.NewService(
		micro.Name("shippy.service.consignment"),
	)

	srv.Init()

	// client
	s.veslClient = pbVesl.NewVesselServiceClient("shippy.service.vessel", nil)

	// service
	pbCons.RegisterShippingServiceHandler(srv.Server(), s)

	if err := srv.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
