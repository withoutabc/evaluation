package file

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"rpc/app/service/file/errs"
	"rpc/app/service/file/rpc/internal/model/dao/table"
	"rpc/app/service/file/rpc/pb"
)

type (
	Model interface {
		Upload(ctx context.Context, level, year, month, set int32) int32
		Download(ctx context.Context, id int64) int32
		GetFileList(ctx context.Context) (*pb.File, int32)
	}
	DefaultModel struct {
		db  *gorm.DB
		rdb *redis.Client
	}
)

func NewModel(db *gorm.DB, rdb *redis.Client) *DefaultModel {
	return &DefaultModel{
		db:  db,
		rdb: rdb,
	}
}

func (d *DefaultModel) Upload(ctx context.Context, level, year, month, set int32) int32 {
	var file table.File
	result := d.db.WithContext(ctx).Where(&table.File{
		Level: level,
		Year:  year,
		Month: month,
		Set:   set,
	}).First(&file)
	//文件不存在或查找错误
	if err := result.Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return errs.InternalServer
		}
	} else {
		return errs.FileNotExist
	}
	//创建
	result = d.db.WithContext(ctx).Create(&table.File{
		Level: level,
		Year:  year,
		Month: month,
		Set:   set,
	})
	if err := result.Error; err != nil {
		return errs.InternalServer
	}
	return errs.No
}

func (d *DefaultModel) Download(ctx context.Context, id int64) int32 {
	//查找文件是否存在
	panic("implement me")
}

func (d *DefaultModel) GetFileList(ctx context.Context) (*pb.File, int32) {
	//TODO implement me
	panic("implement me")
}
