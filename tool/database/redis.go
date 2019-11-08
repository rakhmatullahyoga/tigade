package database

import (
	"github.com/go-redis/redis"
)

type RedisConfig struct {
	Host        []string
	UseSentinel bool
	MasterName  string
}

type RedisClient struct {
	Client *redis.Client
}

func NewRedisConn(opt RedisConfig) *RedisClient {
	var client *redis.Client
	if opt.UseSentinel {
		client = redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName:    opt.MasterName,
			SentinelAddrs: opt.Host,
		})
	} else {
		client = redis.NewClient(&redis.Options{
			Addr: opt.Host[0],
		})
	}

	// health check
	_, err := client.Ping().Result()
	if err != nil {
		panic(err.Error())
	}

	return &RedisClient{Client: client}
}

func (r *RedisClient) CloseConnection() {
	_ = r.Client.Close()
}
