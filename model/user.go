package model

type User struct {
	ID int64 `json:"id"`

	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`

	// Email will be used for email notifications
	Email string `json:"email"`

	// PushID will be used for mobile push notifications
	PushID string `json:"pushId"`

	// PhoneNumber will be used for sms notifications
	PhoneNumber string `json:"phoneNumber"`

	NotificationSettings
}

type NotificationSettings struct {
	AllowPushNotification bool `json:"allowPushNotification"`
	AllowSMS              bool `json:"allowSms"`
	AllowEmail            bool `json:"allowEmail"`
}
