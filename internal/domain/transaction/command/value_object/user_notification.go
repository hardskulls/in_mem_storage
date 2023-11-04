package value_objects

type UserNotification struct {
	Msg string
}

func (u UserNotification) String() string {
	return u.Msg
}
