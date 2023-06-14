package errs

import "errors"

const (
	No               int32 = iota
	Unknown                //未知错误
	RepeatedUsername       //用户名重复
	InternalServer         //服务器内部错误
	NoRecord               //没有记录
	WrongPassword          //密码错误

	WrongTokenType //token类型错误
	WrongSignature //签名认证错误
)

var (
	NoErrs     = errors.New("没有错")
	UnknownErr = errors.New("未知错误")

	RepeatedUsernameErr = errors.New("用户名重复")
	InternalServerErr   = errors.New("服务器内部错误")
	NoRecordErr         = errors.New("没有记录")
	WrongPasswordErr    = errors.New("密码错误")

	WrongTokenTypeErr = errors.New("token类型错误")
	WrongSignatureErr = errors.New("签名认证错误")
)

var ErrorsMap = map[int32]error{
	No:               NoErrs,
	Unknown:          UnknownErr,
	RepeatedUsername: RepeatedUsernameErr,
	InternalServer:   InternalServerErr,
	NoRecord:         NoRecordErr,
	WrongPassword:    WrongPasswordErr,

	WrongTokenType: WrongTokenTypeErr,
	WrongSignature: WrongSignatureErr,
}
