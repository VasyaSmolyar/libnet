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
	conn := &DBConnect{config: cfg}
	db, err := conn.connect()
	if err != nil {
		return nil, err
	}
	conn.Conn = db
	return conn, nil
}

type DBConnect struct {
	config *viper.Viper
	Conn   *sql.DB
}

func (db *DBConnect) connect() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		db.config.GetString("db.user"),
		db.config.GetString("db.pass"),
		db.config.GetString("db.host"),
		db.config.GetString("db.port"),
		db.config.GetString("db.name"),
	)

	conn, err := sql.Open(db.config.GetString("db.driver"), dsn)
	if err != nil {
		return nil, err
	}

	conn.SetConnMaxLifetime(time.Minute * time.Duration(db.config.GetInt("db.maxLifetime")))
	conn.SetMaxOpenConns(db.config.GetInt("db.maxOpenConns"))
	conn.SetMaxIdleConns(db.config.GetInt("db.maxIdleConns"))

	if err = conn.Ping(); err != nil {
		return nil, err
	} else {
		log.Println("DB connection established")
	}

	return conn, nil
}
