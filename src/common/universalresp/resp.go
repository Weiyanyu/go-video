package universalresp

var (
	//common code (0 ~ 99)
	RespCodeSuccess      = ResponceCode{0, "success"}
	RespCodeInteralError = ResponceCode{1, "interal error"}

	RespCodeBindReParamError  = ResponceCode{2, "bind request parameters error"}
	RespCodeUnauthorizedError = ResponceCode{3, "user unauthorized"}

	//user service code (200 ~ 299)
	RespCodeUserRegisterError     = ResponceCode{200, "user register error"}
	RespCodeUserNotFound          = ResponceCode{201, "user not found error"}
	RespCodeUserLoginError        = ResponceCode{202, "user login error"}
	RespCodeUserAlreadyRegistered = ResponceCode{203, "user already registered"}

	//auth service code(300 ~ 399)
	RespCodeGenTokenError      = ResponceCode{301, "gen token error"}
	RespCodeValidateTokenError = ResponceCode{302, "validate token error"}
	RespCodeDeleteTokenError   = ResponceCode{303, "delete token error"}
)

//ResponceCode include code and message
type ResponceCode struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

//ServiceResponse to user
type ServiceResponse struct {
	ResponceCode
	Data interface{} `json:"data"`
}

//NewServiceResponse return a ServiceResponse object with some param
func NewServiceResponse(respCode ResponceCode, data ...interface{}) *ServiceResponse {
	return &ServiceResponse{
		ResponceCode: respCode,
		Data:         data,
	}
}
