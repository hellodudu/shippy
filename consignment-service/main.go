package main

import (
	// 导如 protoc 自动生成的包
	"context"
	"log"

	pbCons "github.com/hellodudu/shippy/proto/consignment"
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
	repo IRepository
}

func (s *service) CreateConsignment(ctx context.Context, req *pbCons.Consignment, out *pbCons.CreateConsignmentResponse) error {
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
	server := micro.NewService(
		micro.Name("shippy.service.consignment"),
	)

	server.Init()
	repo := &Repository{}

	pbCons.RegisterShippingServiceHandler(server.Server(), &service{repo})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
