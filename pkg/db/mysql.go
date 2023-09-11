package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type Config interface {
	GetString(key string) string
	GetInt(key string) int
}

func Init(cfg Config) (*DBConnect, error) {
	db, err := connect(cfg)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	conn := &DBConnect{config: cfg, Conn: db}
	return conn, nil
}

type DBConnect struct {
	config Config
	Conn   *sql.DB
}

func connect(cfg Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.GetString("db.user"),
		cfg.GetString("db.pass"),
		cfg.GetString("db.host"),
		cfg.GetString("db.port"),
		cfg.GetString("db.name"),
	)

	conn, err := sql.Open(cfg.GetString("db.driver"), dsn)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	conn.SetConnMaxLifetime(time.Minute * time.Duration(cfg.GetInt("db.maxLifetime")))
	conn.SetMaxOpenConns(cfg.GetInt("db.maxOpenConns"))
	conn.SetMaxIdleConns(cfg.GetInt("db.maxIdleConns"))

	if err = conn.Ping(); err != nil {
		return nil, errors.WithStack(err)
	} else {
		log.Println("DB connection established")
	}

	return conn, nil
}
