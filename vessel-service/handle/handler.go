package handle

import (
	"context"
	"log"

	pbUser "github.com/hellodudu/shippy/proto/user"
	pbVesl "github.com/hellodudu/shippy/proto/vessel"
	"github.com/hellodudu/shippy/vessel-service/repo"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
)

type VeslSrvHandler struct {
	r repo.IRepository
	s micro.Service
}

func NewVeslSrvHandler(s micro.Service) (*VeslSrvHandler, error) {
	h := &VeslSrvHandler{s: s}

	var err error
	if h.r, err = repo.NewRepository(); err != nil {
		log.Fatalf("failed to call NewRepository(): %v", err)
	}

	if err := micro.RegisterSubscriber("user.create", s.Server(), h.Process); err != nil {
		log.Fatalf("failed to subscrib user.create: %v", err)
	}

	return h, err
}

func (v *VeslSrvHandler) Process(ctx context.Context, event *pbUser.User) error {
	md, _ := metadata.FromContext(ctx)
	log.Printf("[pubsub.1] Received event %+v with metadata %+v\n", event, md)
	// do something with event
	return nil
}

func (v *VeslSrvHandler) Close() {
	v.r.Close()
}

func (v *VeslSrvHandler) Create(ctx context.Context, vesl *pbVesl.Vessel, resp *pbVesl.CreateResp) error {
	resp.Success = true

	err := v.r.Create(vesl)
	if err != nil {
		resp.Success = false
	}

	log.Println("VeslSrvHandler Create result:", resp.Success)

	return err
}

func (v *VeslSrvHandler) FindAvailable(ctx context.Context, spec *pbVesl.Specification, resp *pbVesl.FindAvailableResp) error {

	var err error
	resp.Vessels, err = v.r.FindAvailable(spec)
	return err
}
