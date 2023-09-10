package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/spf13/viper"
)

func Init(cfg *viper.Viper) (*DBConnect, error) {
	db, err := connect(cfg)
	if err != nil {
		return nil, err
	}

	conn := &DBConnect{config: cfg, Conn: db}
	return conn, nil
}

type DBConnect struct {
	config *viper.Viper
	Conn   *sql.DB
}

func connect(cfg *viper.Viper) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.GetString("db.user"),
		cfg.GetString("db.pass"),
		cfg.GetString("db.host"),
		cfg.GetString("db.port"),
		cfg.GetString("db.name"),
	)

	conn, err := sql.Open(cfg.GetString("db.driver"), dsn)
	if err != nil {
		return nil, err
	}

	conn.SetConnMaxLifetime(time.Minute * time.Duration(cfg.GetInt("db.maxLifetime")))
	conn.SetMaxOpenConns(cfg.GetInt("db.maxOpenConns"))
	conn.SetMaxIdleConns(cfg.GetInt("db.maxIdleConns"))

	if err = conn.Ping(); err != nil {
		return nil, err
	} else {
		log.Println("DB connection established")
	}

	return conn, nil
}
