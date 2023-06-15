// Code generated by goctl. DO NOT EDIT.
package types

type UploadReq struct {
	Level int32 `form:"level"`
	Year  int32 `form:"year"`
	Month int32 `form:"month"`
	Set   int32 `form:"set"`
}

type UploadResp struct {
	Status int32  `json:"status"`
	Msg    string `json:"msg"`
}

type DownloadReq struct {
	Type  int32 `form:"type"`
	Id    int64 `form:"id"`
	Level int32 `form:"level"`
	Year  int32 `form:"year"`
	Month int32 `form:"month"`
	Set   int32 `form:"set"`
}

type DownloadResp struct {
	Status int32  `json:"status"`
	Msg    string `json:"msg"`
}

type GetFileListReq struct {
	Level int32 `form:"level,optional"`
	Year  int32 `form:"year,optional"`
	Month int32 `form:"month,optional"`
	Set   int32 `form:"set,optional"`
}

type GetFileListResp struct {
	Status int32       `json:"status"`
	Msg    string      `json:"msg"`
	List   interface{} `json:"list,omitempty"`
}

type RemoveFileReq struct {
	Id int64 `form:"id"`
}

type RemoveFileResp struct {
	Status int32  `json:"status"`
	Msg    string `json:"msg"`
}
