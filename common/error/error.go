package error

const (
	// common error codes, keep same with http status code
	ECodeOK             = 200
	ECodeBadParam       = 400
	ECodeNotFound       = 404
	ECodeUnAuthorized   = 401
	ECodeForbiden       = 403
	ECodeAlreadyExists  = 409
	ECodeNotExists      = 410
	ECodeInternalError  = 500
	ECodeUnImplement    = 501
	ECodeBadGateWay     = 502
	ECodeGateWayTimeOut = 504
	ECodeUnkown         = 520
	ECodeDbError        = 600
)

var EcodeMap = map[int32]string{
	ECodeOK:             "OK",
	ECodeBadParam:       "Bad Params",
	ECodeNotFound:       "Not Found",
	ECodeUnAuthorized:   "UnAuthorized",
	ECodeForbiden:       "Forbiden",
	ECodeAlreadyExists:  "Already Exists",
	ECodeNotExists:      "Not Exists",
	ECodeInternalError:  "Internal Error",
	ECodeUnImplement:    "UnImplement",
	ECodeBadGateWay:     "Bad GateWay",
	ECodeGateWayTimeOut: "Timeout",
	ECodeUnkown:         "Unkown",
	ECodeDbError:        "Db Error",
}
