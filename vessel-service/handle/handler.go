package handle

import (
	"context"

	pbVesl "github.com/hellodudu/shippy/proto/vessel"
	"github.com/hellodudu/shippy/vessel-service/repo"
)

type VeslSrvHandler struct {
	r repo.IRepository
}

func (v *VeslSrvHandler) Create(ctx context.Context, vesl *pbVesl.Vessel, resp *pbVesl.CreateResp) error {
	resp.Success = true

	err := v.r.Create(vesl)
	if err != nil {
		resp.Success = false
	}

	return err
}

func (v *vesselService) FindAvailable(ctx context.Context, spec *pbVesl.Specification, resp *pbVesl.FindAvailableResp) error {

	var err error
	resp.Vessels, err = v.r.FindAvailable(spec)
	return err
}
