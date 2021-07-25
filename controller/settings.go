package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/GordonRamsay-trendyol/notification-manager-microservice/controller/dto"
	"github.com/GordonRamsay-trendyol/notification-manager-microservice/model"
	"github.com/GordonRamsay-trendyol/notification-manager-microservice/service"
)

type UserController interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateNotificationSettings(w http.ResponseWriter, r *http.Request)
	UpdatePushID(w http.ResponseWriter, r *http.Request)
}

func NewUserController(userService service.UserService) UserController {
	return &userControllerImpl{
		userService: userService,
	}
}

type userControllerImpl struct {
	userService service.UserService
}

func (c *userControllerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		WriteErrResponse(w, 400, NewHttpError("400", "Body could not read to bytes."))
		return
	}

	user := &model.User{}
	if err = json.Unmarshal(bytes, user); err != nil {
		WriteErrResponse(w, 400, NewHttpError("400", "Json unmarshall failed could not converted to user object."))
		return
	}

	c.userService.CreateUser(user)
	WriteResponse(w, 200, []byte("SUCCESS"))
}

func (c *userControllerImpl) UpdateNotificationSettings(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	headerUserID := r.Header.Get("USER_ID")

	if len(headerUserID) == 0 {
		WriteErrResponse(w, http.StatusBadRequest, NewHttpError("400", "User ID didn't specified in headers."))
		return
	}

	userID, err := strconv.ParseInt(headerUserID, 10, 2)

	if err != nil {
		WriteErrResponse(w, http.StatusBadRequest, NewHttpError("400", "Request body could not read."))
		return
	}

	req := dto.UpdateNotificationRequest{}
	if err = json.Unmarshal(bytes, &req); err != nil {
		WriteErrResponse(w, http.StatusBadRequest, NewHttpError("400", "Data could not unmarshalled"))
		return
	}

	if err = c.userService.UpdateNotificationSettings(userID, model.NotificationSettings{
		AllowPushNotification: req.AllowPushNotification,
		AllowSMS:              req.AllowSMS,
		AllowEmail:            req.AllowEmail,
	}); err != nil {
		WriteErrResponse(w, 500, NewHttpError("500", "Internal server error!"))
		return
	}

	WriteResponse(w, 200, []byte("SUCCESS"))
}

func (c *userControllerImpl) UpdatePushID(w http.ResponseWriter, r *http.Request) {
	pushId := r.URL.Query().Get("pushId")
	headerUserID := r.Header.Get("USER_ID")

	if len(headerUserID) == 0 {
		WriteErrResponse(w, http.StatusBadRequest, NewHttpError("400", "User ID didn't specified in headers."))
		return
	}

	if len(pushId) == 0 {
		WriteErrResponse(w, 400, NewHttpError("400", "PushID query parameter should be given."))
		return
	}

	userID, err := strconv.ParseInt(headerUserID, 10, 2)

	if err != nil {
		WriteErrResponse(w, 400, NewHttpError("400", "UserID should be integer"))
		return
	}

	if err := c.userService.UpdatePushID(userID, pushId); err != nil {
		WriteErrResponse(w, 500, NewHttpError("500", "Internal Server Error"))
		return
	}

	WriteResponse(w, 200, []byte("SUCCESS"))
}
