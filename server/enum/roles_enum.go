package enum

// RolesEnum -
type RolesEnum int

const (
	// RoleAdmin -
	RoleAdmin = iota
	// RoleManagement -
	RoleManagement
)

func (action RolesEnum) String() string {
	return [...]string{
		"ROLE_ADMIN",
		"ROLE_MANAGEMENT",
	}[action]
}
