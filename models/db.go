package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func init() {
	var err error
	//打开数据库，并复制给db变量，data/data.db数据文件路径
	db, err = gorm.Open("sqlite3", "data.db")

	if err != nil {
		panic("failed to connect database")
	}
	db.SetLogger(logs.GetLogger("orm"))
	db.LogMode(true)
	// Migrate the schema
	db.AutoMigrate(&Book{})
}
