package handle

import (
	"context"
	"log"

	pbUser "github.com/hellodudu/shippy/proto/user"
	"github.com/hellodudu/shippy/user-service/repo"
)

type UserSrvHandler struct {
	r repo.IRepository
}

func NewUserSrvHandler() (*UserSrvHandler, error) {
	h := &UserSrvHandler{}

	var err error
	if h.r, err = repo.NewRepository(); err != nil {
		log.Fatalf("failed to call NewRepository(): %v", err)
	}

	return h, err
}

func (h *UserSrvHandler) Close() {
	h.r.Close()
}

func (h *UserSrvHandler) Create(ctx context.Context, req *pbUser.User, resp *pbUser.Response) error {
	if err := h.r.Create(req); err != nil {
		return nil
	}
	resp.User = req
	return nil
}

func (h *UserSrvHandler) Get(ctx context.Context, req *pbUser.User, resp *pbUser.Response) error {
	u, err := h.r.Get(req.Id)
	if err != nil {
		return err
	}
	resp.User = u
	return nil
}

func (h *UserSrvHandler) GetAll(ctx context.Context, req *pbUser.Request, resp *pbUser.Response) error {
	users, err := h.r.GetAll()
	if err != nil {
		return err
	}
	resp.Users = users
	return nil
}

func (h *UserSrvHandler) Auth(ctx context.Context, req *pbUser.User, resp *pbUser.Token) error {
	_, err := h.r.GetByEmailAndPassword(req)
	if err != nil {
		return err
	}
	resp.Token = "`x_2nam"
	return nil
}

func (h *UserSrvHandler) ValidateToken(ctx context.Context, req *pbUser.Token, resp *pbUser.Token) error {
	return nil
}
