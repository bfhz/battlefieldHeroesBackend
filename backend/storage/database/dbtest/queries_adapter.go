package dbtest

import (
	"github.com/gocraft/dbr"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/model"
)

type QueriesAdapter struct {
	Server                     model.Server
	ErrGetServerByID          error
	ErrGetServerByName error
	ErrGetServerLogin error

	Player               model.Player
	ErrGetPlayerByToken error
	ErrGetPlayerByID    error

	Heroes                  []model.Hero
	ErrGetHeroesByPID error
	Hero                    model.Hero
	ErrGetHeroByName       error
	HeroStats               model.HeroStats
	ErrGetHeroStats        error
	ErrUpdateStats      error
}

func (a *QueriesAdapter) GetServerByID(sess *dbr.Session, serverID int) (model.Server, error) {
	return a.Server, a.ErrGetServerByID
}
func (a *QueriesAdapter) GetServerByName(sess *dbr.Session, soldierName string) (model.Server, error) {
	return a.Server, a.ErrGetServerByName
}
func (a *QueriesAdapter) GetServerLogin(sess *dbr.Session, accountName string) (model.Server, error) {
	return a.Server, a.ErrGetServerLogin
}

func (a *QueriesAdapter) GetPlayerByToken(sess *dbr.Session, token string) (model.Player, error) {
	return a.Player, a.ErrGetPlayerByToken
}
func (a *QueriesAdapter) GetPlayerByID(sess *dbr.Session, playerID int) (model.Player, error) {
	return a.Player, a.ErrGetPlayerByID
}

func (a *QueriesAdapter) GetHeroesByPID(sess *dbr.Session, playerID int) ([]model.Hero, error) {
	return a.Heroes, a.ErrGetHeroesByPID
}
func (a *QueriesAdapter) GetHeroByName(sess *dbr.Session, heroName string) (model.Hero, error) {
	return a.Hero, a.ErrGetHeroByName
}
func (a *QueriesAdapter) GetHeroStats(sess *dbr.Session, heroID int) (model.HeroStats, error) {
	return a.HeroStats, a.ErrGetHeroStats
}
func (a *QueriesAdapter) UpdateStats(tx *dbr.Tx, heroID int, pr *model.HeroStats) error {
	return a.ErrUpdateStats
}
