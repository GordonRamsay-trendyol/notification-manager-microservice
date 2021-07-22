package dto

type UpdateNotificationRequest struct {
	AllowPushNotification bool `json:"allowPushNotification"`
	AllowSMS              bool `json:"allowSms"`
	AllowEmail            bool `json:"allowEmail"`
}
