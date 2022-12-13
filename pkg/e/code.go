package e

const (
	SUCCESS               = 200
	UpdatePasswordSuccess = 201
	NotExistInentifier    = 202
	ERROR                 = 500
	InvalidParams         = 400

	ErrorExistUser      = 10001
	ErrorNotExistUser   = 10002
	ErrorFailEncryption = 10003
	ErrorNotCompare     = 10004

	ErrorAuthToken             = 20001
	ErrorAuthCheckTokenFail    = 20002
	ErrorAuthCheckTokenTimeout = 20003

	ErrorDatabase = 40001

	ErrorUploadFile = 50001
)
