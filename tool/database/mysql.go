package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlConfig struct {
	Host     string
	Username string
	Password string
	DbName   string
	Charset  string
	Pool     int
}

type MysqlClient struct {
	dbWrite *sql.DB
	dbRead  *sql.DB
}

// defer call this function on main program
func (m *MysqlClient) CloseConnection() {
	m.dbWrite.Close()
	m.dbRead.Close()
}

func NewMysqlMasterOnly(cfg MysqlConfig) *MysqlClient {
	db, err := buildConnection(cfg)
	if err != nil {
		panic(err)
	}
	return &MysqlClient{
		dbRead:  db,
		dbWrite: db,
	}
}

func NewMysqlMasterSlave(writeConfig, readConfig MysqlConfig) *MysqlClient {
	dbWrite, err := buildConnection(writeConfig)
	if err != nil {
		panic(err)
	}

	dbRead, err := buildConnection(readConfig)
	if err != nil {
		panic(err)
	}
	return &MysqlClient{
		dbRead:  dbRead,
		dbWrite: dbWrite,
	}
}

func buildConnection(cfg MysqlConfig) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&enableCircuitBreaker=true",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.DbName,
		cfg.Charset),
	)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(cfg.Pool)
	db.SetConnMaxLifetime(300 * time.Second)
	err = db.Ping()
	return db, err
}
