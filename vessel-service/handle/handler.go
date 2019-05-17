package handle

import (
	"context"
	"encoding/json"
	"log"

	pbUser "github.com/hellodudu/shippy/proto/user"
	pbVesl "github.com/hellodudu/shippy/proto/vessel"
	"github.com/hellodudu/shippy/vessel-service/repo"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
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

	h.Broker().Subscribe("user.created", h.subHandle)

	return h, err
}

func (v *VeslSrvHandler) subHandle(pub broker.Publication) error {
	log.Println("sub handle topic:", pub.Topic())
	log.Println("sub message header:", pub.Message().Header)

	u := &pbUser.User{}
	json.Unmarshal(pub.Message().Body, u)
	log.Println("sub message body:", pub.Message().Body)
	log.Println("sub message json unmarshal:", u)
	return nil
}

func (v *VeslSrvHandler) Broker() broker.Broker {
	return v.s.Options().Broker
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
