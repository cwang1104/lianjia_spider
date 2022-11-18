package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"

	"time"
)

const (
	DBUserName = "root"
	DBPassword = "4524"
	DBAddress  = "127.0.0.1"
	DBName     = "lj_spider"

	DBMaxOpenConns = 200
	DBMaxIdleConns = 100
	DBMaxLifeTime  = 120
)

var db *gorm.DB

// DbInit 数据库初始化
func DbInit() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DBUserName, DBPassword,
		DBAddress, DBName)
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                     dsn,
		DontSupportRenameColumn: true,
	}), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Println("gorm open failed", err)
		return err
	}

	sqlDb, err := db.DB()
	if err != nil {
		log.Println("db.DB failed", err)
		return err
	}
	sqlDb.SetMaxOpenConns(DBMaxOpenConns)
	sqlDb.SetMaxIdleConns(DBMaxIdleConns)
	sqlDb.SetConnMaxLifetime(time.Duration(DBMaxLifeTime) * time.Second)
	return nil
}

func init() {
	err := DbInit()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	log.Println("db init success")
}
