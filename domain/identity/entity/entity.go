package entity

// Role represents role as integer
// so we can use mask to check more difficult permissions
type Role int64

const (
	RoleAnon = Role(iota) + 100
	RoleUser
	RoleAdmin
	RoleSuperAdmin
)

type (
	// UserContext represents minimum information
	// about authenticated user so other domains can use it
	// to apply authorization and so on.
	UserContext struct {
		ID    int64
		Roles []Role
	}
)
