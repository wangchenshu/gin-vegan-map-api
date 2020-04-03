package enum

// AuthEnum -
type AuthEnum int

const (
	// AuthSuccess -
	AuthSuccess = iota
	// AuthFail -
	AuthFail
	// CheckTokenSuccess -
	CheckTokenSuccess
	// CheckTokenFail -
	CheckTokenFail
	// TokenExpired -
	TokenExpired
	// GenerateTokenFail -
	GenerateTokenFail
)

func (action AuthEnum) String() string {
	return [...]string{
		"驗證成功",
		"驗證失敗",
		"驗證Token成功",
		"驗證Token失敗",
		"Token逾期",
		"建立Token失敗",
	}[action]
}
