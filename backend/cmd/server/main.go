package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/iflow/monitor/internal/config"
	"github.com/iflow/monitor/internal/handler"
	"github.com/iflow/monitor/internal/middleware"
	"github.com/iflow/monitor/internal/model"
)

func main() {
	cfg := config.Init()
	db := model.InitDB(cfg)
	if db == nil {
		log.Fatal("数据库连接失败")
	}
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	model.AutoMigrate(db)
	model.InitRedis(cfg)

	r := gin.Default()
	r.Use(middleware.Cors())

	api := r.Group("/api")
	{
		api.POST("/login", handler.Login)
		api.POST("/register", handler.Register)

		auth := api.Group("")
		auth.Use(middleware.JwtAuth())
		{
			auth.GET("/user/info", handler.GetUserInfo)
			auth.PUT("/user/info", handler.UpdateUserInfo)
			auth.GET("/users", handler.GetUsers)
			auth.GET("/clusters", handler.GetClusters)
			auth.POST("/clusters", handler.CreateCluster)
			auth.GET("/datasources", handler.GetDataSources)
			auth.GET("/dashboard", handler.GetDashboard)
		}
	}

	log.Printf("服务器启动在端口 %s", cfg.Server.Port)
	r.Run(":" + cfg.Server.Port)
}
