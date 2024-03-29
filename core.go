package tigade

import (
	"github.com/rakhmatullahyoga/tigade/config"
	"github.com/rakhmatullahyoga/tigade/tool"
	"github.com/rakhmatullahyoga/tigade/tool/database"
	"time"
)

// Dependencies adapter definition
type Logger interface {
	LogInfo(trackId string, tags []string, message string)
	LogError(trackId string, tags []string, message string, error error, code int)
	LogPanic(trackId string, tags []string, message string, error error, code int)
}

// Worker is an interface for background job processor
type Worker interface {
	EnqueuePerform(jobName string, payload map[string]interface{})
	DelayPerform(jobName string, payload map[string]interface{}, delay time.Duration)
	SchedulePerform(jobName string, payload map[string]interface{}, time time.Time)
}

type PersistenceConnection interface {
	CloseConnection()
}

type connection struct {
	clients []PersistenceConnection
}

func NewConnections() *connection {
	return &connection{make([]PersistenceConnection, 0)}
}

func (c *connection) CloseConnection() {
	for _, client := range c.clients {
		client.CloseConnection()
	}
}

func (c *connection) Add(client PersistenceConnection) {
	c.clients = append(c.clients, client)
}

// Core of the service: the application use cases and its dependencies
type CoreService struct {
	connections *connection
	Logger      Logger
}

func NewCoreService() *CoreService {
	cfg := config.GetInstance()
	conns := NewConnections()

	var mysql *database.MysqlClient
	if cfg.MysqlMasterSlave {
		mysql = database.NewMysqlMasterSlave(cfg.MysqlWrite, cfg.MysqlRead)
	} else {
		mysql = database.NewMysqlMasterOnly(cfg.MysqlWrite)
	}
	conns.Add(mysql)

	elastic := database.NewElasticConn(cfg.Elastic)
	conns.Add(elastic)

	mongo := database.NewMongoConn(cfg.Mongo)
	conns.Add(mongo)

	redis := database.NewRedisConn(cfg.Redis)
	conns.Add(redis)

	logger := tool.NewLogger(cfg.Environment)

	return &CoreService{conns, logger}
}

func (cs *CoreService) Shutdown() {
	cs.connections.CloseConnection()
}
