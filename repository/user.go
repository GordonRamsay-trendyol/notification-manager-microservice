package repository

import "github.com/GordonRamsay-Trendyol/notification-manager/model"

type UserRepository interface {
	Save(user model.User) model.User

	Update(user model.User) model.User

	FindById(id int64) (*model.User, error)
}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{}
}

type userRepositoryImpl struct {
}

func (repo *userRepositoryImpl) Save(user model.User) model.User {
	return model.User{}
}

func (repo *userRepositoryImpl) Update(user model.User) model.User {
	return model.User{}
}

func (repo *userRepositoryImpl) FindById(id int64) (*model.User, error) {
	return nil, nil
}
