package repository

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/GordonRamsay-trendyol/notification-manager-microservice/model"
)

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
	b, _ := json.Marshal(user)
	query := fmt.Sprintf("INSERT INTO `users` (KEY, VALUE) VALUES ('%d', %v)",
		user.ID, string(b))

	_, err := cls.Query(query, nil)

	if err != nil {
		log.Printf("sql error, err: %v\n", err)
		return model.User{}
	}

	return user
}

func (repo *userRepositoryImpl) Update(user model.User) model.User {
	b, _ := json.Marshal(user)
	query := fmt.Sprintf("UPDATE `users` USE KEYS '%d' SET users=%v",
		user.ID, string(b))

	_, err := cls.Query(query, nil)

	if err != nil {
		log.Printf("sql error, err: %v\n", err)
		return model.User{}
	}
	return user
}

func (repo *userRepositoryImpl) FindById(id int64) (*model.User, error) {
	query := fmt.Sprintf("SELECT * FROM `users` USE KEYS '%d'", id)

	rows, err := cls.Query(query, nil)

	if err != nil {
		log.Printf("sql error, err: %v\n", err)
		return nil, err
	}

	result := bucketSingleResult{}
	b := rows.Raw().NextBytes()

	if err = json.Unmarshal(b, &result); err != nil {
		return nil, err
	}

	if result.User != nil {
		return result.User, nil
	}

	return nil, err
}

type bucketSingleResult struct {
	User *model.User `json:"users"`
}
