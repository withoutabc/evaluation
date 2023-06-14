package table

import (
	"gorm.io/gorm"
	"rpc/utils"
	"time"
)

const TableNameFile = "files"

type File struct {
	ID             int64     `json:"id"  gorm:"primarykey;type:bigint"` //id
	Level          int32     `json:"level"`                             //等级4/6
	Year           int32     `json:"year"`                              //年份
	Month          int32     `json:"month"`                             //月份
	Set            int32     `json:"set"`                               //第几套
	Download       int64     `json:"download"`                          //原文件下载次数
	ResultDownload int64     `json:"result_download"`                   //结果下载次数
	CreateTime     time.Time `json:"create_time"`                       //创建时间
	UpdateTime     time.Time `json:"update_time"`                       //更新时间
}

func (*File) TableName() string {
	return TableNameFile
}

// BeforeCreate 创建文件时生成id
func (f *File) BeforeCreate(_ *gorm.DB) (err error) {
	f.ID = utils.GenerateId()
	f.Download = 0
	f.CreateTime = time.Now()
	f.UpdateTime = time.Now()
	return nil
}
