package res

type ErrorCode int

const (
	SettingError ErrorCode = 1001
)

var (
	ErrorMap = map[ErrorCode]string{
		SettingError: "系统错误",
	}
)