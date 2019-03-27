package model

import (
	"github.com/gocraft/dbr"
	"github.com/guregu/null"
)

// Player
type Player struct {
	ID           int      `db:"user_id"`
	Username     string   `db:"heroName"`
	heroID       int      `db:"heroID"`
	Password     string   `db:"sessionID"`
	sessionID    string   `db:"sessionID"`
	SelectedHero null.Int `db:"heroID"`
}

// GetPlayerByToken returns a player associated with given sessionID
func (q *Queries) GetPlayerByToken(sess *dbr.Session, sessionID string) (player Player, err error) {
	err = sess.
		Select(
			"user_id",
			"heroName",
			"sessionID",
			"sessionID",
			"heroID",
		).
		From(tableHeroes).
		Where("sessionID = ?", sessionID).
		LoadOne(&player)
	return player, err
}

// GetPlayerByID returns a player associated with given playerID
func (q *Queries) GetPlayerByID(sess *dbr.Session, playerID int) (player Player, err error) {
	err = sess.
		Select(
			"user_id",
			"heroName",
			"sessionID",
			"sessionID",
			"heroID",
		).
		From(tableHeroes).
		Where("user_id = ?", playerID). //from network.cli_data
		LoadOne(&player)
	return player, err
}
