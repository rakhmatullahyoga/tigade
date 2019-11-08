package database

import (
	"gopkg.in/olivere/elastic.v5"
)

type ElasticConfig struct {
	Host string
}

type ElasticClient struct {
	Client *elastic.Client
}

// Construct elastic connection
func NewElasticConn(elasticCfg ElasticConfig) *ElasticClient {
	esHost := elasticCfg.Host
	url := elastic.SetURL(esHost)
	client, err := elastic.NewClient(url)
	if err != nil {
		panic(err)
	}
	return &ElasticClient{client}
}

// defer call this function on main program
func (e *ElasticClient) CloseConnection() {
	e.Client.Stop()
}
