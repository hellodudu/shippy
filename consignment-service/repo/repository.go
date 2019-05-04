package repo

import (
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
	return r, err
}

func (repo *repository) collection() *mgo.Collection {
	return repo.session.DB(DB_NAME).C(CON_COLLECTION)
}

func (repo *repository) Close() {
	repo.session.Close()
}

func (repo *repository) Create(c *pbCons.Consignment) error {
	return repo.collection().Insert(c)
}

func (repo *repository) GetAll() ([]*pbCons.Consignment, error) {
	var cons []*pbCons.Consignment
	err := repo.collection().Find(nil).All(&cons)
	return cons, err
}
