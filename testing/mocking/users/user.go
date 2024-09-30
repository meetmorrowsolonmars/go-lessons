package users

type User struct {
	ID   string
	Name string
	Role Role
}

type Role string

func (r Role) CanRead() bool {
	switch r {
	case RoleGuest, RoleUser, RoleAdmin:
		return true
	}
	return false
}

func (r Role) CanWrite() bool {
	switch r {
	case RoleGuest:
		return false
	case RoleUser, RoleAdmin:
		return true
	}
	return false
}

func (r Role) String() string {
	return string(r)
}

const (
	RoleGuest Role = "Guest"
	RoleUser  Role = "User"
	RoleAdmin Role = "Admin"
)
