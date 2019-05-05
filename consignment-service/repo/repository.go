package repo

import (
	"fmt"

	"github.com/hellodudu/shippy/consignment-service/db"
	pbCons "github.com/hellodudu/shippy/proto/consignment"
	"gopkg.in/mgo.v2"
)

const (
	DB_NAME        = "shippy"
	CON_COLLECTION = "consignments"
)

type IRepository interface {
	Create(consignment *pbCons.Consignment) error
	GetAll() ([]*pbCons.Consignment, error)
	Close()
}

type repository struct {
	consignments []*pbCons.Consignment
	session      *mgo.Session
}

func NewRepository() (IRepository, error) {
	r := &repository{
		consignments: make([]*pbCons.Consignment, 0),
	}

	var err error
	r.session, err = db.NewSession("localhost:27017")
	if err != nil {
		return nil, err
	}

	s := r.session.Copy()
	defer s.Close()
	r.collection(s).Find(nil).All(&r.consignments)

	return r, err
}

func (repo *repository) collection(s *mgo.Session) *mgo.Collection {
	return s.DB(DB_NAME).C(CON_COLLECTION)
}

func (repo *repository) Close() {
	repo.session.Close()
}

func (repo *repository) Create(c *pbCons.Consignment) error {

	for _, value := range repo.consignments {
		if value.Id == c.Id {
			return fmt.Errorf("consignment id: %s existed", c.Id)
		}
	}

	s := repo.session.Copy()
	defer s.Close()

	repo.consignments = append(repo.consignments, c)

	return repo.collection(s).Insert(c)
}

func (repo *repository) GetAll() ([]*pbCons.Consignment, error) {
	return repo.consignments, nil
}
