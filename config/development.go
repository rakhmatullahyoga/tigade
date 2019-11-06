package config

import (
	"github.com/subosito/gotenv"
	"os"
	"strconv"
	"strings"
	"tigade/tool/database"
	"time"
)

func NewDevelopmentConfig() *Config {
	err := gotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	port, err := strconv.Atoi(os.Getenv("SERVICE_PORT"))
	if err != nil {
		port = 9999
	}

	dbPool, err := strconv.Atoi(os.Getenv("MYSQL_WRITE_POOL"))
	if err != nil {
		dbPool = 50
	}
	dbCfg := database.MysqlConfig{
		Username: os.Getenv("MYSQL_WRITE_USERNAME"),
		Password: os.Getenv("MYSQL_WRITE_PASSWORD"),
		Host:     os.Getenv("MYSQL_WRITE_HOST"),
		DbName:   os.Getenv("MYSQL_WRITE_DBNAME"),
		Charset:  os.Getenv("MYSQL_WRITE_CHARSET"),
		Pool:     dbPool,
	}

	esCfg := database.ElasticConfig{
		Host: os.Getenv("ELASTICSEARCH_URL"),
	}

	mongoHosts := strings.Split(os.Getenv("MONGO_HOST"), ",")
	maxPool, err := strconv.ParseUint(os.Getenv("MONGO_MAX_POOL_SIZE"), 10, 64)
	if err != nil {
		maxPool = 5
	}
	minPool, err := strconv.ParseUint(os.Getenv("MONGO_MIN_POOL_SIZE"), 10, 64)
	if err != nil {
		minPool = 1
	}
	mongoTimeout, err := strconv.Atoi(os.Getenv("MONGO_TIMEOUT_SECOND"))
	if err != nil {
		mongoTimeout = 1
	}
	mongoCfg := database.MongoConfig{
		Host:           mongoHosts,
		Username:       os.Getenv("MONGO_USERNAME"),
		Password:       os.Getenv("MONGO_PASSWORD"),
		DbName:         os.Getenv("MONGO_DBNAME"),
		ReplicaSetName: os.Getenv("MONGO_SET_NAME"),
		Timeout:        time.Second * time.Duration(mongoTimeout),
		MaxPool:        maxPool,
		MinPool:        minPool,
	}

	redisHosts := strings.Split(os.Getenv("REDIS_HOST"), ",")
	redisCfg := database.RedisConfig{
		Host:        redisHosts,
		UseSentinel: false,
	}

	return &Config{
		AppPort:    uint16(port),
		MysqlWrite: dbCfg,
		Elastic:    esCfg,
		Mongo:      mongoCfg,
		Redis:      redisCfg,
	}
}
