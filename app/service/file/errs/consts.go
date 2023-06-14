package errs

import "errors"

const (
	No             int32 = iota
	Unknown              //未知错误
	InternalServer       //服务器内部错误
	NoRecord             //没有记录

	WrongTokenType //token类型错误
	WrongSignature //签名认证错误

	FileTooBig         //文件太大
	FileGetWrong       //文件获取失败
	FileSystemInternal //文件系统出错
	FileRepeated       //文件重复上传
	FileNameWrong      //文件等级错误

	FileNotExist //下载时文件不存在
)

var (
	NoErrs     = errors.New("没有错")
	UnknownErr = errors.New("未知错误")

	InternalServerErr = errors.New("服务器内部错误")
	NoRecordErr       = errors.New("没有记录")

	WrongTokenTypeErr = errors.New("token类型错误")
	WrongSignatureErr = errors.New("签名认证错误")

	FileTooBigErr         = errors.New("文件太大")
	FileGetWrongErr       = errors.New("文件获取失败")
	FileSystemInternalErr = errors.New("文件系统出错")
	FileRepeatedErr       = errors.New("文件重复上传")
	FileNameWrongErr      = errors.New("文件属性错误")

	FileNotExistErr = errors.New("下载时文件不存在")
)

var ErrorsMap = map[int32]error{
	No:             NoErrs,
	Unknown:        UnknownErr,
	InternalServer: InternalServerErr,
	NoRecord:       NoRecordErr,

	WrongTokenType: WrongTokenTypeErr,
	WrongSignature: WrongSignatureErr,

	FileTooBig:         FileTooBigErr,
	FileGetWrong:       FileGetWrongErr,
	FileSystemInternal: FileSystemInternalErr,
	FileRepeated:       FileRepeatedErr,
	FileNameWrong:      FileNameWrongErr,
	FileNotExist:       FileNotExistErr,
}
