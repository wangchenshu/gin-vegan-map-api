package enum

// StatusEnum -
type StatusEnum int

const (
	// Success -
	Success = iota
	// Failed -
	Failed
	// GetSuccess -
	GetSuccess
	// GetFailed -
	GetFailed
	// AddSuccess -
	AddSuccess
	// AddFailed -
	AddFailed
	// UpdateSuccess -
	UpdateSuccess
	// UpdateFailed -
	UpdateFailed
	// DeleteSuccess -
	DeleteSuccess
	// DeleteFailed -
	DeleteFailed
	// InvalidParams -
	InvalidParams
	// NoPermission -
	NoPermission
)

func (action StatusEnum) String() string {
	return [...]string{
		"成功",
		"失敗",
		"查詢成功",
		"查詢失敗",
		"新增成功",
		"新增失敗",
		"修改成功",
		"修改失敗",
		"刪除成功",
		"刪除失敗",
		"參數錯誤",
		"無權限訪問",
	}[action]
}
