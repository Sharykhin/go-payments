package model

type (
	User struct {
		id    int64
		email string
	}

	UserInterface interface {
		GetID() int64
		GetEmail() string
	}
)

func NewUser(userID int64, email string) UserInterface {
	return User{
		id:    userID,
		email: email,
	}
}

func (u User) GetID() int64 {
	return u.id
}

func (u User) GetEmail() string {
	return u.email
}
