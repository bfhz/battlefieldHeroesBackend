package rank

import (
	"github.com/sirupsen/logrus"

	"github.com/Synaxis/battlefieldHeroesBackend/backend/network"
)

type reqGetStats struct {
	Owner      string   `fesl:"owner"`      // owner=5
	OwnerType  string   `fesl:"ownerType"`  // ownerType=1
	PeriodID   string   `fesl:"periodId"`   // periodId=0
	PeriodPast string   `fesl:"periodPast"` // periodPast=0
	Keys       []string `fesl:"keys"`       // keys.0=c_apr, keys.1=level ...
}

type ansGetStats struct {
	Txn       string      `fesl:"TXN"`
	OwnerID   int         `fesl:"ownerId"`
	OwnerType int         `fesl:"ownerType"`
	Stats     []statsPair `fesl:"stats"`
}

type statsPair struct {
	Key   string `fesl:"key"`
	Text  string `fesl:"text,omitempty"`
	Value string `fesl:"value"`
}

// GetStats - Get basic stats about a soldier/owner (account holder)
func (r *Ranking) GetStats(event network.EventClientCommand) {
	switch event.Client.GetClientType() {
	case "server":
		r.serverGetStats(&event)
	default:
		r.clientGetStats(&event)
	}
}

func (r *Ranking) serverGetStats(event *network.EventClientCommand) {
	r.getStats(event)
}

func (r *Ranking) clientGetStats(event *network.EventClientCommand) {
	r.getStats(event)
}

func (r *Ranking) getStats(event *network.EventClientCommand) {
	if event.Command.Message["owner"] == "Current" {
		// In the tutorial "Current" is reserved name for the hero.
		return
	}

	ownerID, err := event.Command.Message.IntVal("owner")
	if err != nil {
		logrus.
			WithField("owner", event.Command.Message["owner"]).
			WithField("cmd", "rank.GetStats").
			Warnf("Cannot parse ownerID")
		return
	}

	heroID := event.Client.PlayerData.HeroID
	if heroID == 0 {
		if event.Client.Type == "server" {
			heroID = ownerID
			// Server uses only heroID to identify owners
			logrus.Warnf("GetStats (server), replacing heroID with ownerID")
		} else {
			// Whether client is not yet logged in it requires data about the
			// tutorial completion, sadly we need to look for a player's master
			// hero
			player, err := r.DB.GetPlayerByID(r.DB.NewSession(), ownerID)
			if err != nil {
				logrus.
					WithField("heroID", heroID).
					WithField("owner", event.Command.Message["owner"]).
					WithField("cmd", "rank.GetStats").
					Warnf("Cannot find player of given ownerID")
				return
			}

			if player.SelectedHero.Valid == false {
				logrus.
					WithField("heroID", heroID).
					WithField("owner", event.Command.Message["owner"]).
					WithField("player", player.ID).
					WithField("SelectedHero", player.SelectedHero).
					WithField("cmd", "rank.GetStats").
					Warnf("Player has unset selected hero")
				return
			}

			logrus.Warnf("GetStats (client), replacing ownerID with heroID with %v", player.SelectedHero.Int64)
			heroID = int(player.SelectedHero.Int64)
		}
	}

	logrus.
		WithField("heroID", heroID).
		WithField("ownerID", ownerID).
		WithField("clientType", event.Client.Type).
		Print("Requesting rank.GetStats")

	heroStats, err := r.DB.GetHeroStats(r.DB.NewSession(), heroID)
	if err != nil {
		logrus.WithError(err).Errorf("Cannot retrieve stats of hero %d", heroID)
		return
	}

	keys := event.Command.Message.ArrayStrings("keys")
	stats, err := r.fetchStats(&heroStats, keys)
	if err != nil {
		logrus.WithError(err).Error("Cannot retrieve given stats")
		return
	}

	r.answer(
		event.Client,
		event.Command.PayloadID,
		ansGetStats{
			Txn:       rankGetStats,
			OwnerID:   heroID,
			OwnerType: 1,
			Stats:     stats,
		},
	)
}
