package enum

// UserEnum -
type UserEnum int

const (
	// CheckUser -
	CheckUser = iota
	// UserNotExist -
	UserNotExist
)

func (action UserEnum) String() string {
	return [...]string{
		"查看帳號",
		"用戶不存在",
	}[action]
}
