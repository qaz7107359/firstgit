package db

import (
	"TARG_revenue_report_backend/db/model"
	"TARG_revenue_report_backend/utils"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strings"
	"time"
)

var _db *gorm.DB
var err error

func InitDB() {
	host := utils.DbHost
	port := utils.DbPort
	database := utils.DbName
	username := utils.DbUser
	password := utils.DbPass
	charset := utils.DbCharset

	dsn := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", database, "?charset=" + charset + "&parseTime=true"}, "")
	err := Database(dsn)
	if err != nil {
		fmt.Println(err)
	}
}

func Database(connString string) error {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connString, // DSN data source name
		DefaultStringSize:         256,        // string 类型字段的默认长度
		DisableDatetimePrecision:  true,       // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,      // 根据版本自动配置
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 是否开启单数表名
		},
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	_db = db
	migration(db)
	return err
}

func NewDB() *gorm.DB {
	return _db
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}

func migration(DB *gorm.DB) {
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			//&taskModel.AgvTask{}, //
			&model.User{},
			&model.Table{},
			&model.ModifyLog{},
			&model.RolePrivilege{},
			&model.RevenueData{},
		)
	if err != nil {
		return
	}
}
