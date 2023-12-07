package errno

const (
	Success = 200

	InvalidParams = 400

	Error                       = 500
	ErrorDatabaseOperate        = 50001
	ErrorDatabaseRecordNotFound = 50002
	ErrorUserNotExist           = 50003
	ErrorDeviceNotExist         = 50004
	ErrorUserHaveExisted        = 50005
)
