package repository

import "github.com/tienducitt/go-restful/src/model"

type IUserRepository interface {
	GetAll() ([]model.User, error)
	Get(id int64) (*model.User, error)
	Create(u *model.User) error
	Update(u *model.User) error
	Delete(u *model.User) error
}
