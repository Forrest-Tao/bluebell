package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInValidPassword
	CodeServerBusy

	CodeNeedLogin
	CodeInvalidToken
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户名存在",
	CodeUserNotExist:    "用户名不存在",
	CodeInValidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙",
	CodeNeedLogin:       "需要登录",
	CodeInvalidToken:    "无效的 token",
}

func (rc ResCode) Msg() string {
	if msg, ok := codeMsgMap[rc]; ok {
		return msg
	}
	//如果获取失败则返回服务繁忙的错误
	return codeMsgMap[CodeServerBusy]
}