package handle

import (
	"context"
	"encoding/json"
	"log"

	pbUser "github.com/hellodudu/shippy/proto/user"
	"github.com/hellodudu/shippy/user-service/repo"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"golang.org/x/crypto/bcrypt"
)

type UserSrvHandler struct {
	r repo.IRepository
	s micro.Service
	p micro.Publisher
}

func NewUserSrvHandler(s micro.Service) (*UserSrvHandler, error) {
	h := &UserSrvHandler{
		s: s,
	}

	var err error
	if h.r, err = repo.NewRepository(); err != nil {
		log.Fatalf("failed to call NewRepository(): %v", err)
	}

	h.p = micro.NewPublisher("user.create", s.Client())
	if h.p == nil {
		log.Fatalf("failed to call NewPublisher")
	}

	return h, err
}

func (h *UserSrvHandler) Broker() broker.Broker {
	return h.s.Options().Broker
}

func (h *UserSrvHandler) Close() {
	h.r.Close()
}

func (h *UserSrvHandler) Create(ctx context.Context, req *pbUser.User, resp *pbUser.Response) error {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	req.Password = string(hashedPwd)
	if err := h.r.Create(req); err != nil {
		return nil
	}

	h.p.Publish(ctx, req)

	resp.User = req
	return nil
}

func (h *UserSrvHandler) Auth(ctx context.Context, req *pbUser.User, resp *pbUser.Token) error {
	u, err := h.r.GetByEmailAndPassword(req)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		return err
	}
	// t, err := h.tokenService.Encode(u)
	// if err != nil {
	// 	return err
	// }
	resp.Token = "`x_2nam"
	return nil
}

func (h *UserSrvHandler) ValidateToken(ctx context.Context, req *pbUser.Token, resp *pbUser.Token) error {
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

func (h *UserSrvHandler) publishEvent(user *pbUser.User) error {
	body, err := json.Marshal(user)
	if err != nil {
		return err
	}

	msg := &broker.Message{
		Header: map[string]string{
			"id": user.Id,
		},
		Body: body,
	}

	log.Println("publish event:", msg)

	if err := h.Broker().Publish("user.created", msg); err != nil {
		log.Fatalf("user create publish failed: %v\n", err)
	}
	return nil
}
