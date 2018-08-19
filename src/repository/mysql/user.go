package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/tienducitt/go-restful/src/model"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAll() (users []model.User, err error) {
	err = r.db.Find(&users).Error
	return
}

func (r *UserRepository) Get(id int64) (user *model.User, err error) {
	user = &model.User{}
	if err = r.db.First(user, id).Error; gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	return
}

func (r *UserRepository) Create(u *model.User) error {
	return r.db.Create(u).Error
}

func (r *UserRepository) Update(u *model.User) error {
	return r.db.Save(u).Error
}

func (r *UserRepository) Delete(u *model.User) error {
	return r.db.Delete(u).Error
}
