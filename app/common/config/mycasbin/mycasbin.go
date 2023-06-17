package mycasbin

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func InitEnforcer() *casbin.Enforcer {

	a, err := gormadapter.NewAdapter("mysql", "root:224488@tcp(127.0.0.1:3306)/evaluation", true)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	e, err := casbin.NewEnforcer("E:\\我的菜鸡项目实战\\evaluation\\model.conf", a)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	return e
}
