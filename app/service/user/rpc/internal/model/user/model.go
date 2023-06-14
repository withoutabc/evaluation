package user

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"log"
	"rpc/app/service/user/errs"
	"rpc/app/service/user/rpc/internal/model/dao/table"
	"rpc/app/service/user/rpc/pb"
	"rpc/utils"
	"rpc/utils/secret"
)

type (
	Model interface {
		Register(ctx context.Context, username, password string) int32
		Login(ctx context.Context, username, password string) (*pb.LoginResp_Data, int32)
		Refresh(ctx context.Context, refreshToken string) (string, string, int32)
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

func (m *DefaultModel) Register(ctx context.Context, username, password string) int32 {
	//用户名是否存在
	var user table.User
	result := m.db.WithContext(ctx).Where(&table.User{Username: username}).First(&user)
	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			log.Println(result.Error)
			return errs.InternalServer
		}
	} else {
		return errs.RepeatedUsername
	}
	//生成盐值
	salt, err := secret.GenerateSalt()
	if err != nil {
		log.Println(err)
		return errs.InternalServer
	}
	//加密
	hashedPassword := secret.HashWithSalt(password, salt)
	//用户信息写入数据库
	result = m.db.Create(&table.User{
		Username: username,
		Password: string(hashedPassword),
		Salt:     salt,
		Role:     "user",
	})
	if result.Error != nil {
		log.Println(result.Error)
		return errs.InternalServer
	}
	return errs.No
}

func (m *DefaultModel) Login(ctx context.Context, username, password string) (*pb.LoginResp_Data, int32) {
	//用户名是否存在
	var mysqlUser table.User
	result := m.db.WithContext(ctx).Where(&table.User{Username: username}).First(&mysqlUser)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println(err)
			return &pb.LoginResp_Data{}, errs.NoRecord
		}
		log.Println(err)
		return &pb.LoginResp_Data{}, errs.InternalServer
	}
	//验证密码是否正确
	if string(secret.HashWithSalt(password, mysqlUser.Salt)) != mysqlUser.Password {
		return &pb.LoginResp_Data{}, errs.WrongPassword
	}
	//生成token
	aToken, rToken, _, err := secret.GenToken(mysqlUser.ID, mysqlUser.Role)
	if err != nil {
		log.Println(err)
		return &pb.LoginResp_Data{}, errs.InternalServer
	}
	return &pb.LoginResp_Data{
		Id:           mysqlUser.ID,
		Role:         mysqlUser.Role,
		LoginTime:    utils.Time(),
		AccessToken:  aToken,
		RefreshToken: rToken,
	}, errs.No
}

func (m *DefaultModel) Refresh(ctx context.Context, refreshToken string) (string, string, int32) {
	accessToken, refreshToken, _, err := secret.RefreshToken(refreshToken)
	if err != nil {
		if err.Error() == "错误的类型" {
			return "", "", errs.WrongTokenType
		}
		if err.Error() == "invalid refresh token signature" {
			return "", "", errs.WrongSignature
		}
		log.Printf("refresh token err:%v\n", err)
		return "", "", errs.InternalServer
	}
	return accessToken, refreshToken, errs.No

}
