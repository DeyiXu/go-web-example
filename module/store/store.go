package store

import (
	"github.com/bwmarrin/snowflake"
	"github.com/nilorg/pkg/logger"

	"github.com/jinzhu/gorm"

	// mysql驱动
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var (
	// DB 数据库连接
	DB *gorm.DB
	// SnowflakeNode 雪花ID节点
	snowflakeNode *snowflake.Node
)

func initDB() {
	// 初始化数据库
	db, err := gorm.Open("mysql", viper.GetString("mysql.address"))
	if err != nil {
		logger.Fatalf("open mysql error: %s \n", err)
	}
	err = db.DB().Ping()
	if err != nil {
		logger.Fatalf("ping mysql error: %s \n", err)
	}

	db.LogMode(viper.GetBool("mysql.log"))

	db.DB().SetMaxOpenConns(viper.GetInt("mysql.max_open"))
	db.DB().SetMaxIdleConns(viper.GetInt("mysql.max_idle"))
	// db.DB().SetConnMaxLifetime(time.Hour)

	DB = db
}

func initID() {
	// 设置雪花ID节点
	node, err := snowflake.NewNode(viper.GetInt64("snowflake.node"))
	if err != nil {
		logger.Fatalf("snowflake:%v\n", err)
	}
	snowflakeNode = node
}

// Start 启动存储
func Start() {
	initDB()
	initID()
}

// Close 关闭
func Close() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			logger.Errorf("mysql db error:%v\n", err)
		}
	}
}

// NewSnowflakeID 获取雪花ID
func NewSnowflakeID() snowflake.ID {
	return snowflakeNode.Generate()
}
