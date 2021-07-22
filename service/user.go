package service

import (
	"encoding/json"
	"errors"

	"github.com/GordonRamsay-Trendyol/notification-manager/messaging"
	"github.com/GordonRamsay-Trendyol/notification-manager/model"
	"github.com/GordonRamsay-Trendyol/notification-manager/util"
)

const (
	UserUpdateTopic = "update_user"
)

type UserService interface {
	UpdateNotificationSettings(userID int64, settings model.NotificationSettings) error
}

func NewUserService(config messaging.EventReaderConfig) UserService {
	service := &userServiceImpl{}

	reader := messaging.NewKafkaReader(config, service)
	reader.Listen()

	return service
}

type userServiceImpl struct {
	notificationSettingsRepo interface{}
}

func (service *userServiceImpl) UpdateNotificationSettings(userID int64, settings model.NotificationSettings) error {
	return nil
}

func (service *userServiceImpl) Receive(msg []byte) error {
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
