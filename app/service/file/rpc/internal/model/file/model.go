package file

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"rpc/app/service/file/errs"
	"rpc/app/service/file/rpc/internal/model/dao/table"
)

type (
	Model interface {
		Upload(ctx context.Context, level, year, month, set int32) int32
		Download(ctx context.Context, id int64, downloadType int32) int32
		GetFileList(ctx context.Context, level, year, month, set int32) ([]*table.File, int32)
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
		return errs.FileIsExist
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

func (d *DefaultModel) Download(ctx context.Context, id int64, downloadType int32) int32 {
	//查找文件是否存在
	result := d.db.WithContext(ctx).Where(&table.File{ID: id})
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errs.FileNotExist
		}
		return errs.InternalServer
	}

	//增加下载次数
	if downloadType == 1 {
		result = d.db.WithContext(ctx).Where(&table.File{ID: id}).UpdateColumn("download", gorm.Expr("download + ?", 1))
		if err := result.Error; err != nil {
			return errs.InternalServer
		}
	}
	if downloadType == 2 {
		//判断是否生成了结果
		var file table.File
		result = d.db.WithContext(ctx).Where(&table.File{ID: id}).First(&file)
		if err := result.Error; err != nil {
			return errs.InternalServer
		}
		if file.Result != 2 {
			return errs.ResultNotExist
		}
		//否则下载次数+1
		result = d.db.WithContext(ctx).Where(&table.File{ID: id}).UpdateColumn("result_download", gorm.Expr("result_download + ?", 1))
		if err := result.Error; err != nil {
			return errs.InternalServer
		}
	}
	return errs.No
}

func (d *DefaultModel) GetFileList(ctx context.Context, level, year, month, set int32) ([]*table.File, int32) {
	var files []*table.File
	cond := table.File{}
	if level != 0 {
		cond.Level = level
	}
	if year != 0 {
		cond.Year = year
	}
	if month != 0 {
		cond.Month = month
	}
	if set != 0 {
		cond.Set = set
	}
	result := d.db.WithContext(ctx).Where(&cond).Find(files)
	if err := result.Error; err != nil {
		return nil, errs.InternalServer
	}
	return files, errs.No
}
