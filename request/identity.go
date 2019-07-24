package request

type (
	LoginRequest struct {
		Email    string `json:"email" binding:"required,email,max=80"`
		Password string `json:"password" binding:"required"`
	}
)
