package tigade

import (
	"tigade/config"
	"tigade/tool"
	"tigade/tool/database"
)

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

type CoreService struct {
	connections *connection
	Logger      tool.Logger
}

func NewCoreService() *CoreService {
	cfg := config.GetInstance()
	conns := NewConnections()

	mysql, err := database.NewMysqlConn(&cfg.MysqlWrite, cfg.MysqlRead)
	CheckError(err)
	conns.Add(mysql)

	elastic, err := database.NewElasticConn(cfg.Elastic)
	CheckError(err)
	conns.Add(elastic)

	mongo, err := database.NewMongoConn(cfg.Mongo)
	CheckError(err)
	conns.Add(mongo)

	redis, err := database.NewRedisConn(cfg.Redis)
	CheckError(err)
	conns.Add(redis)

	logger := tool.NewLogger(cfg.Environment)

	return &CoreService{conns, logger}
}

func (cs *CoreService) Shutdown() {
	cs.connections.CloseConnection()
}
