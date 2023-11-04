package repository

import (
	"project/internal/model"
)

// type Conn struct {
// 	db *gorm.DB
// }

// func NewConn(db *gorm.DB) (*Conn, error) {
// 	if db == nil {
// 		return nil, errors.New("db connection not given")
// 	}

// 	return &Conn{db: db}, nil

// }

// type Users interface {
// 	CreateUser(model.User) (model.User, error)
// 	FetchUserByEmail(string) (model.User, error)
// }

func (r *Repo) CreateUser(u model.User) (model.User, error) {
	err := r.db.Create(&u).Error
	if err != nil {
		return model.User{}, err
	}
	return u, nil
}

func (r *Repo) FetchUserByEmail(s string) (model.User, error) {
	var u model.User
	tx := r.db.Where("email=?", s).First(&u)
	if tx.Error != nil {
		return model.User{}, nil
	}
	return u, nil

}
