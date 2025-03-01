package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// 用户模块的错误 code = 1000...
	ERROR_USERNAME_USED       = 1001
	ERROR_PASSWORD_WRONG      = 1002
	ERROR_USER_NOT_EXIST      = 1003
	ERROR_TOKEN_NOT_EXIST     = 1004
	ERROR_TOKEN_RUNTIME       = 1005
	ERROR_TOKEN_WRONG         = 1006
	ERROR_TOKEN_TYPE_WRONG    = 1007
	ERROR_USER_NOT_PERMISSION = 1008
	ERROR_PASSWORD_EMPTY      = 1009

	// 文章模块的错误 code = 2000...
	ERROR_ART_NOT_EXIST = 2001

	// 分类模块的错误 code = 3000...
	ERROR_CATENAME_USER  = 3001
	ERROR_CATE_NOT_EXIST = 3002
)

var codeMsg = map[int]string{
	SUCCESS:                   "OK",
	ERROR:                     "FAIL",
	ERROR_USERNAME_USED:       "用户名已使用",
	ERROR_PASSWORD_WRONG:      "密码错误",
	ERROR_USER_NOT_EXIST:      "用户不存在",
	ERROR_TOKEN_NOT_EXIST:     "TOKEN不存在，请重新登录",
	ERROR_TOKEN_RUNTIME:       "TOKEN已过期，请重新登录",
	ERROR_TOKEN_WRONG:         "TOKEN不正确，请重新登录",
	ERROR_TOKEN_TYPE_WRONG:    "TOKEN格式错误，请重新登录",
	ERROR_USER_NOT_PERMISSION: "该用户无权限",
	ERROR_ART_NOT_EXIST:       "文章不存在",
	ERROR_CATENAME_USER:       "该分类已存在",
	ERROR_CATE_NOT_EXIST:      "该分类不存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
