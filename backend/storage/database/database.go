package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"github.com/sirupsen/logrus"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/config"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/model"
)

var (
	logger = logrus.WithField("prefix", "database")
)

type DB struct {
	*model.Queries
}

// New connects and wraps connection into wrapper
func New() (Adapter, error) {
	conn, err := NewConn()
	if err != nil {
		return nil, err
	}
	return &DB{&model.Queries{conn}}, nil
}

func (db *DB) NewSession() *dbr.Session {
	return db.Queries.Conn.NewSession(NewLogger())
}

func (db *DB) Begin(sess *dbr.Session) (*dbr.Tx, error) {
	return sess.Begin()
}

func (db *DB) Commit(tx *dbr.Tx) error {
	err := tx.Commit()
	if err != nil {
		logger.
			WithError(err).
			Error("Problem with commiting the transaction")
		return err
	}
	return nil
}

func (db *DB) Rollback(tx *dbr.Tx) error {
	err := tx.Rollback()
	if err != nil {
		logger.
			WithError(err).
			Error("Problem with rollbacking the transaction")
		return err
	}
	return nil
}

// NewConn tries to establish connection with database
func NewConn() (*dbr.Connection, error) {
	return dbConnect(newMySQL())
}

// dbConnect establishes connection to the database, and then pings it
// to verify if it responds.
func dbConnect(driver, dsnAddr string) (*dbr.Connection, error) {
	logger := NewLogger()

	conn, err := dbr.Open(driver, dsnAddr, logger)
	if err != nil {
		logrus.Fatal(err)
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		logrus.Fatal(err)
		return nil, err
	}

	conn.SetMaxIdleConns(config.Database.DatabaseMaxIdleConns)
	conn.SetMaxOpenConns(config.Database.DatabaseMaxOpenConns)

	return conn, nil
}
