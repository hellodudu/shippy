package repo

import (
	"fmt"
	"log"
	"os"

	pbUser "github.com/hellodudu/shippy/proto/user"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type IRepository interface {
	Create(*pbUser.User) error
	Get(id string) (*pbUser.User, error)
	GetAll() ([]*pbUser.User, error)
	GetByEmailAndPassword(*pbUser.User) (*pbUser.User, error)
	Close()
}

type repository struct {
	users []*pbUser.User
	db    *gorm.DB
}

func NewRepository() (IRepository, error) {
	r := &repository{
		users: make([]*pbUser.User, 0),
	}

	dsn := os.Getenv("DB_HOST")
	if len(dsn) == 0 {
		log.Fatal("cannot get env DB_HOST")
	}

	var err error
	r.db, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal("failed to connect database")
	}

	r.db.AutoMigrate(&r.users)
	r.db.Find(&r.users)

	return r, nil
}

func (repo *repository) Close() {
	repo.db.Close()
}

func (repo *repository) Create(u *pbUser.User) error {
	for _, value := range repo.users {
		if value.Id == u.Id {
			return fmt.Errorf("user id: %s existed", u.Id)
		}
	}

	repo.users = append(repo.users, u)
	return repo.db.Create(&u).Error
}

func (repo *repository) Get(id string) (*pbUser.User, error) {
	for _, val := range repo.users {
		if val.Id == id {
			return val, nil
		}
	}

	return nil, fmt.Errorf("cannot find user by id:%s", id)
}

func (repo *repository) GetAll() ([]*pbUser.User, error) {
	return repo.users, nil
}

func (repo *repository) GetByEmailAndPassword(u *pbUser.User) (*pbUser.User, error) {
	for _, val := range repo.users {
		if val.Id == u.Id {
			return val, nil
		}
	}

	return nil, fmt.Errorf("cannot get user by email and password")
}
