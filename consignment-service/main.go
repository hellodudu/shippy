package main

import (
	// 导如 protoc 自动生成的包
	"context"
	"log"

	pb "github.com/hellodudu/shippy/proto"
	"github.com/micro/go-micro"
)

const (
	PORT = ":50051"
)

type IRepository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

type Repository struct {
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.consignments = append(repo.consignments, consignment)
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

type service struct {
	repo IRepository
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, out *pb.CreateConsignmentResponse) error {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	out.Created = true
	out.Consignment = consignment
	return nil
}

func (s *service) GetConsignments(ctx context.Context, _ *pb.GetRequest, out *pb.GetConsignmentsResponse) error {
	out.Consignments = s.repo.GetAll()
	return nil
}

func main() {
	server := micro.NewService(
		micro.Name("shippy.service.consignment"),
	)

	server.Init()
	repo := &Repository{}

	pb.RegisterShippingServiceHandler(server.Server(), &service{repo})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
