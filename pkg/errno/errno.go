package errno

var MsgFlags = map[int]string{
	Success: "ok",

	InvalidParams: "请求参数错误",

	Error:                       "fail",
	ErrorDatabaseOperate:        "数据库操作错误",
	ErrorDatabaseRecordNotFound: "数据库返回空值",
	ErrorUserNotExist:           "用户不存在",
	ErrorDeviceNotExist:         "设备不存在",
	ErrorUserHaveExisted:        "用户已存在",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[Error]
}
