package config

import (
	"os"
	"sync"
	"tigade/tool/database"
)

//Config represent global config
type Config struct {
	Environment string
	AppPort     uint16
	Elastic     database.ElasticConfig
	MysqlRead   *database.MysqlConfig
	MysqlWrite  database.MysqlConfig
	Mongo       database.MongoConfig
	Redis       database.RedisConfig
}

// Config as singleton
var instance *Config
var once sync.Once

func GetInstance() *Config {
	once.Do(func() {
		env := os.Getenv("ENV")
		switch env {
		case "production":
			instance = NewProductionConfig()
		case "staging":
			instance = NewStagingConfig()
		default:
			instance = NewDevelopmentConfig()
		}
		instance.Environment = env
	})
	return instance
}
