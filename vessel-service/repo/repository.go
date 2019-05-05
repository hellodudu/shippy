package repo

import (
	"fmt"

	pbVesl "github.com/hellodudu/shippy/proto/vessel"
	"github.com/hellodudu/shippy/vessel-service/db"
	"gopkg.in/mgo.v2"
)

const (
	DB_NAME         = "shippy"
	VESL_COLLECTION = "vessels"
)

type IRepository interface {
	Create(v *pbVesl.Vessel) error
	FindAvailable(s *pbVesl.Specification) ([]*pbVesl.Vessel, error)
	Close()
}

type repository struct {
	vessels []*pbVesl.Vessel
	session *mgo.Session
}

func NewRepository() (IRepository, error) {
	r := &repository{
		vessels: make([]*pbVesl.Vessel, 0),
	}

	var err error
	r.session, err = db.NewSession("localhost:27017")
	if err != nil {
		return nil, err
	}

	s := r.session.Copy()
	defer s.Close()
	r.collection(s).Find(nil).All(&r.vessels)

	return r, err
}

func (repo *repository) collection(s *mgo.Session) *mgo.Collection {
	return s.DB(DB_NAME).C(VESL_COLLECTION)
}

func (repo *repository) Close() {
	repo.session.Close()
}

func (repo *repository) Create(v *pbVesl.Vessel) error {
	for _, value := range repo.vessels {
		if value.Id == v.Id {
			return fmt.Errorf("vessel id: %s existed", v.Id)
		}
	}

	repo.vessels = append(repo.vessels, v)

	s := repo.session.Copy()
	defer s.Close()

	return repo.collection(s).Insert(v)
}

func (repo *repository) FindAvailable(s *pbVesl.Specification) ([]*pbVesl.Vessel, error) {
	ret := make([]*pbVesl.Vessel, 0)

	for _, val := range repo.vessels {
		if val.MaxWeight >= s.Weight {
			ret = append(ret, val)
		}
	}

	return ret, nil
}
