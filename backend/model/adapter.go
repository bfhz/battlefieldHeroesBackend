package model

import (
	"github.com/gocraft/dbr"
)

type Queries struct {
	Conn *dbr.Connection
}

type QueriesAdapter interface {
	GetServerByID(sess *dbr.Session, serverID int) (Server, error)
	GetServerByName(sess *dbr.Session, soldierName string) (Server, error)
	GetServerLogin(sess *dbr.Session, accountName string) (Server, error)

	GetPlayerByToken(sess *dbr.Session, sessionID string) (Player, error)
	GetPlayerByID(sess *dbr.Session, playerID int) (Player, error)

	GetHeroesByPID(sess *dbr.Session, playerID int) ([]Hero, error)
	GetHeroByName(sess *dbr.Session, heroName string) (Hero, error)
	GetHeroStats(sess *dbr.Session, heroID int) (HeroStats, error)

	UpdateStats(tx *dbr.Tx, heroID int, pr *HeroStats) error
}
