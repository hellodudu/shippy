package main

import (
	// 导如 protoc 自动生成的包
	"context"
	"log"

	pb "github.com/hellodudu/shippy/proto/consignment"
	"github.com/micro/go-micro"
)

const (
	PORT = ":50051"
)

//
// 仓库接口
//
type IRepository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error) // 存放新货物
}

//
// 我们存放多批货物的仓库，实现了 IRepository 接口
//
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

//
// 定义微服务
//
type service struct {
	repo Repository
}

//
// service 实现 consignment.pb.go 中的 ShippingServiceServer 接口
// 使 service 作为 gRPC 的服务端
//
// 托运新的货物
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, out *pb.Response) error {
	// 接收承运的货物
	log.Println("recv req:", req)
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}
	out = &pb.Response{Created: true, Consignment: consignment, Consignments: s.repo.consignments}
	return nil
}

func (s *service) GetConsignments(ctx context.Context, _ *pb.GetRequest, out *pb.Response) error {
	out = &pb.Response{Created: false, Consignment: nil, Consignments: s.repo.consignments}
	return nil
}

func main() {
	server := micro.NewService(
		micro.Name("shippy.service.consignment"),
	)

	server.Init()
	repo := Repository{}

	// 向 rRPC 服务器注册微服务
	// 此时会把我们自己实现的微服务 service 与协议中的 ShippingServiceServer 绑定
	pb.RegisterShippingServiceHandler(server.Server(), &service{repo})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
