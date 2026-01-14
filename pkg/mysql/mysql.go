package mysql

import (
	"log"
	"os"
	"sync"
	"time"
	"vben/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var once sync.Once

// 定义数据库连接结构体
type connect struct {
	primaryDB *gorm.DB
	currentDB *gorm.DB
}

// 设置一个全局变量保存数据库连接
var _connect *connect

// 连接 MySQL 数据库
func connectMysql() {
	// 启用打印日志
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // 日志等级: Silent、Error、Warn、Info
			Colorful:      false,       // 禁用彩色打印
		},
	)

	// 从配置文件获取主库的 DSN
	primary := config.Database.Primary.UserName + ":" + config.Database.Primary.Password + "@tcp(" + config.Database.Primary.Host + ":" + config.Database.Primary.Port + ")/" + config.Database.Primary.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	// 打开主库连接
	primaryDB, err := gorm.Open(mysql.Open(primary), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Printf("主数据库连接失败:", err)
		return
	}

	sqlDB, err := primaryDB.DB()
	if err != nil {
		log.Printf("获取从数据库SQL DB失败: %v", err)
		return
	}
	// 设置连接池参数
	sqlDB.SetMaxIdleConns(config.Database.MaxIdleConn)
	sqlDB.SetMaxOpenConns(config.Database.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 打开从库连接
	// 从配置文件获取从库的 DSN
	var secondaryDB *gorm.DB
	if config.Database.Secondary.Enabled {
		secondary := config.Database.Secondary.UserName + ":" + config.Database.Secondary.Password + "@tcp(" + config.Database.Secondary.Host + ":" + config.Database.Secondary.Port + ")/" + config.Database.Secondary.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"

		var err error
		secondaryDB, err = gorm.Open(mysql.Open(secondary), &gorm.Config{
			Logger: newLogger,
		})
		if err != nil {
			log.Printf("从数据库连接失败: %v", err)
			return
		}

		// 可以在这里添加连接池配置
		sqlDB, err := secondaryDB.DB()
		if err != nil {
			log.Printf("获取从数据库SQL DB失败: %v", err)
			return
		}
		// 设置连接池参数
		sqlDB.SetMaxIdleConns(config.Database.MaxIdleConn)
		sqlDB.SetMaxOpenConns(config.Database.MaxOpenConn)
		sqlDB.SetConnMaxLifetime(time.Hour)

	}

	// 初始化数据库连接结构体，默认使用主数据库
	_connect = &connect{
		primaryDB: primaryDB,
		currentDB: secondaryDB, // 默认连接到主数据库
	}
}

// SetPrimary 切换到主库
func SetPrimary() {
	if _connect != nil {
		_connect.currentDB = _connect.primaryDB
	}
}

/*使用
mysql.MysqlDb(true) 使用从数据库
mysql.MysqlDb() 使用主数据库
*/
// MysqlDb 获取当前使用的数据库连接
func MysqlDb(param ...bool) *gorm.DB {
	if _connect == nil {
		once.Do(func() {
			connectMysql()
		})
	}
	if len(param) == 0 {
		return _connect.primaryDB
	}
	if len(param) > 0 {
		return _connect.currentDB
	}
	return _connect.primaryDB
}
