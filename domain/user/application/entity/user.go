package entity

type (
	User struct {
		ID       int64    `json:"id"`
		Identity Identity `json:"identity"`
	}

	Identity struct {
		Password string `json:"-"`
	}
)
