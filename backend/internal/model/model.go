package model

import (
	"github.com/go-redis/redis/v8"
	"github.com/iflow/monitor/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var rdb *redis.Client

func InitDB(cfg *config.Config) *gorm.DB {
	dsn := cfg.Database.User + ":" + cfg.Database.Password + "@tcp(" + cfg.Database.Host + ":3306)/" + cfg.Database.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}
	db = d
	return db
}

func InitRedis(cfg *config.Config) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host + ":" + "6379",
		Password: cfg.Redis.Password,
		DB:       0,
	})
}

func GetDB() *gorm.DB {
	return db
}

func GetRedis() *redis.Client {
	return rdb
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate()
}
