package constant

type ErrorCode struct {
	Code    int
	Message string
}

var (
	UNDEFINED       = ErrorCode{-1, "unsettled"}
	SUCCESS         = ErrorCode{0, "ok"}
	PARAMS_ERROR    = ErrorCode{40000, "请求参数错误"}
	NOT_LOGIN_ERROR = ErrorCode{40100, "未登录"}
	NO_AUTH_ERROR   = ErrorCode{40101, "无权限"}
	NOT_FOUND_ERROR = ErrorCode{40400, "请求数据不存在"}
	FORBIDDEN_ERROR = ErrorCode{40300, "禁止访问"}
	SYSTEM_ERROR    = ErrorCode{50000, "系统内部异常"}
	OPERATION_ERROR = ErrorCode{50001, "操作失败"}
)
