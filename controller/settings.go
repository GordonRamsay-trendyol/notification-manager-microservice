package controller

import (
	"github.com/GordonRamsay-Trendyol/notification-manager/controller/dto"
	"github.com/GordonRamsay-Trendyol/notification-manager/model"
	"github.com/GordonRamsay-Trendyol/notification-manager/service"
)

type UserController struct {
	userService service.UserService
}

func (controller *UserController) updateNotificationSettings(userID int64, request dto.UpdateNotificationRequest) {
	settings := model.NotificationSettings{
		AllowPushNotification: request.AllowPushNotification,
		AllowSMS:              request.AllowSMS,
		AllowEmail:            request.AllowEmail,
	}

	if err := controller.userService.UpdateNotificationSettings(userID, settings); err != nil {
		// do something with the error, maybe error is the base response...
	}

}

func (controller *UserController) updatePushID(userID int64, pushID string) {
	// set the new push ID
}
