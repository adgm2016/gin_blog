package model

import (
	"fmt"
	"gin_blog/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var (
	db       *gorm.DB
	err      error
	dbString = fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
)

func InitDb() {
	//dbString = "root:@/gin_blog?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open(utils.Db, dbString)
	fmt.Println(dbString)
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数", err)
	}

	// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
	db.SingularTable(true)
	db.AutoMigrate(&User{}, &Article{}, &Category{})
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	db.DB().SetConnMaxLifetime(10 * time.Second)

	//关闭数据库
	//db.Close()
}
