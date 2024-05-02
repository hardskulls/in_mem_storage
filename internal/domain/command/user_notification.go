package command

// UserNotification is a value object representing a user notification.
type UserNotification struct {
	msg string
}

func NewUserNotification(msg string) UserNotification {
	return UserNotification{msg: msg}
}

func (u UserNotification) String() string {
	return u.msg
}
