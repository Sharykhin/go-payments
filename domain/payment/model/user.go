package model

type (
	// User is a concrete implementation of UserInterface
	// that is used across payment domain
	User struct {
		id    int64
		email string
	}

	// UserInterface describes user representation in a payment domain
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
