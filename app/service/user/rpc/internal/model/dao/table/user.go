package table

import (
	"gorm.io/gorm"
	"rpc/utils"
)

const TableNameUser = "users"

type User struct {
	ID       int64  `json:"id"  gorm:"primarykey;type:bigint" `
	Username string `json:"username"  gorm:"type:varchar(40);not null"`
	Password string `json:"password"  gorm:"not null;type:longblob"`
	Salt     []byte `json:"salt" gorm:"not null"`
	Role     string `json:"role" gorm:"not null"`
}

func (*User) TableName() string {
	return TableNameUser
}

// BeforeCreate 创建用户时生成id
func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	u.ID = utils.GenerateId(0)
	return nil
}
