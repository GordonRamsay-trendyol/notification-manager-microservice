package service

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/GordonRamsay-trendyol/notification-manager-microservice/messaging"
	"github.com/GordonRamsay-trendyol/notification-manager-microservice/model"
	"github.com/GordonRamsay-trendyol/notification-manager-microservice/repository"
	"github.com/GordonRamsay-trendyol/notification-manager-microservice/util"
)

type UserService interface {
	CreateUser(user *model.User) error
	UpdateNotificationSettings(userID int64, settings model.NotificationSettings) error
	UpdatePushID(userID int64, pushID string) error
}

func NewUserService() UserService {
	repository := repository.NewUserRepository()
	service := &userServiceImpl{userRepository: repository}

	reader := messaging.NewReader(newEventReaderConfig(UserUpdateTopic), service)
	go reader.Listen()

	return service
}

type userServiceImpl struct {
	userRepository repository.UserRepository
}

func (service *userServiceImpl) UpdateNotificationSettings(userID int64, settings model.NotificationSettings) error {
	user, err := service.userRepository.FindById(userID)

	if err != nil {
		log.Printf("user not found, err: %v\n", err)
		return errors.New("user not found")
	}

	user.NotificationSettings = settings
	service.userRepository.Update(*user)

	return nil
}

func (service *userServiceImpl) UpdatePushID(userID int64, pushID string) error {
	return nil
}

func (service *userServiceImpl) CreateUser(user *model.User) error {
	service.userRepository.Save(*user)
	return nil
}

func (service *userServiceImpl) Receive(msg []byte) error {
	// TODO: No need to update user changes...
	userUpdateMsg := userUpdateMsg{}

	if err := json.Unmarshal(msg, &userUpdateMsg); err != nil {
		return errors.New("invalid update message error")
	}

	noNeedUpdate := !util.ContainsStrList(userUpdateMsg.UpdatedFields, "EMAIL", "PHONE_NUMBER")

	if noNeedUpdate {
		return nil
	}

	// update current user information

	return nil
}

type userUpdateMsg struct {
	ID            string   `json:"id"`
	Email         string   `json:"email"`
	PhoneNumber   string   `json:"phoneNumber"`
	UpdatedFields []string `json:"updatedFields"`
}
