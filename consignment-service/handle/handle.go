package handle

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/hellodudu/shippy/consignment-service/repo"
	pbCons "github.com/hellodudu/shippy/proto/consignment"
	pbVesl "github.com/hellodudu/shippy/proto/vessel"
)

type ConsSrvHandler struct {
	r repo.IRepository
	c pbVesl.VesselServiceClient
}

func NewConsSrvHandler() (*ConsSrvHandler, error) {
	h := &ConsSrvHandler{}

	var err error
	if h.r, err = repo.NewRepository(); err != nil {
		log.Fatalf("failed to call NewRepository(): %v", err)
	}

	h.c = pbVesl.NewVesselServiceClient("shippy.service.vessel", nil)

	return h, err
}

func (s *ConsSrvHandler) CreateConsignment(ctx context.Context, req *pbCons.Consignment, out *pbCons.CreateConsignmentResponse) error {

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
	defer s.r.Close()

	out.Created = true
	out.Consignment = req
	return nil
}

func (s *ConsSrvHandler) GetConsignments(ctx context.Context, _ *pbCons.GetRequest, out *pbCons.GetConsignmentsResponse) error {
	var err error
	out.Consignments, err = s.r.GetAll()
	defer s.r.Close()
	return err
}
