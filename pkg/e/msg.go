package e

var MsgFlags = map[int]string{
	SUCCESS:               "ok",
	UpdatePasswordSuccess: "修改密码成功",
	NotExistInentifier:    "该第三方账号未绑定",
	ERROR:                 "fail",
	InvalidParams:         "请求参数错误",

	ErrorExistUser:      "已存在该用户名",
	ErrorNotExistUser:   "该用户不存在",
	ErrorFailEncryption: "加密失败",
	ErrorNotCompare:     "密码不匹配",
	ErrorSendEmail:      "发送邮件失败",

	ErrorAuthToken:             "token生成失败",
	ErrorAuthCheckTokenFail:    "token鉴权失败",
	ErrorAuthCheckTokenTimeout: "token已超时",

	ErrorDatabase: "数据库操作出错,请重试",

	ErrorUploadFile: "文件上传失败",
}

// GetMsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
