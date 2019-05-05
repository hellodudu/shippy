package handle

import (
	"context"
	"log"

	pbVesl "github.com/hellodudu/shippy/proto/vessel"
	"github.com/hellodudu/shippy/vessel-service/repo"
)

type VeslSrvHandler struct {
	r repo.IRepository
}

func NewVeslSrvHandler() (*VeslSrvHandler, error) {
	h := &VeslSrvHandler{}

	var err error
	if h.r, err = repo.NewRepository(); err != nil {
		log.Fatalf("failed to call NewRepository(): %v", err)
	}

	return h, err
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
