package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type MongoConfig struct {
	Host           []string
	Username       string
	Password       string
	DbName         string
	ReplicaSetName string
	Timeout        time.Duration
	MaxPool        uint64
	MinPool        uint64
}

type MongoClient struct {
	Client   *mongo.Client
	Database *mongo.Database
}

// Construct mongo connection(s)
func NewMongoConn(mongoCfg MongoConfig) *MongoClient {
	credential := options.Credential{
		Username:   mongoCfg.Username,
		Password:   mongoCfg.Password,
		AuthSource: mongoCfg.DbName,
	}
	clientOpt := &options.ClientOptions{}
	clientOpt.SetHosts(mongoCfg.Host)
	clientOpt.SetReplicaSet(mongoCfg.ReplicaSetName)
	clientOpt.SetReadPreference(readpref.SecondaryPreferred())
	clientOpt.SetAuth(credential)
	clientOpt.SetMinPoolSize(mongoCfg.MinPool)
	clientOpt.SetMaxPoolSize(mongoCfg.MaxPool)
	clientOpt.SetMaxConnIdleTime(mongoCfg.Timeout)
	client, err := mongo.Connect(context.Background(), clientOpt)
	if err != nil {
		panic(err)
	}
	db := client.Database(mongoCfg.DbName)

	return &MongoClient{
		Client:   client,
		Database: db,
	}
}

// defer call this function on main program
func (m *MongoClient) CloseConnection() {
	_ = m.Client.Disconnect(context.Background())
}
