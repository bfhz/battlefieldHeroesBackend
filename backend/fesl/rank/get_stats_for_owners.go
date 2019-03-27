package rank

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/Synaxis/battlefieldHeroesBackend/backend/network"
)

type reqGetStatsForOwners struct {
	Txn      string       `fesl:"TXN"`  // =GetStatsForOwners
	Keys     []string     `fesl:"keys"` // = elo,c_team,xp,level ...
	OwnerID  []statsOwner `fesl:"owners"`
	PeriodID int          `fesl:"periodId"` // =0
}

type statsOwner struct {
	OwnerID   int `fesl:"ownerId"`   // ID of hero
	OwnerType int `fesl:"ownerType"` // =1
}

type ansGetStatsForOwners struct {
	Txn   string           `fesl:"TXN"`
	Stats []statsContainer `fesl:"stats"`
}

type statsContainer struct {
	Stats     []statsPair `fesl:"stats"`
	OwnerID   int         `fesl:"ownerId"`
	OwnerType int         `fesl:"ownerType"`
}

func (r *Ranking) GetStatsForOwners(event network.EventClientCommand) {
	switch event.Client.GetClientType() {
	case "server":
		r.serverGetStatsForOwners(&event)
	default:
		r.clientGetStatsForOwners(&event)
	}
}

func (r *Ranking) serverGetStatsForOwners(event *network.EventClientCommand) {
	r.getStatsForOwners(event)
}

func (r *Ranking) clientGetStatsForOwners(event *network.EventClientCommand) {
	r.getStatsForOwners(event)
}

func (r *Ranking) getStatsForOwners(event *network.EventClientCommand) {
	stats := []statsContainer{}

	owners, err := event.Command.Message.IntVal("owners.[]")
	if err != nil {
		logrus.WithError(err).Warn("GetStatsForOwners")
		return
	}

	keys := event.Command.Message.ArrayStrings("keys")
	for i := 0; i < owners; i++ {
		ownerID, err := event.Command.Message.IntVal(fmt.Sprintf("owners.%d.ownerId", i))
		if err != nil {
			logrus.WithError(err).Warn("Cannot parse onwer in GetStatsForOwners")
		}

		heroStats, err := r.DB.GetHeroStats(r.DB.NewSession(), ownerID)
		if err != nil {
			logrus.WithError(err).Errorf("Cannot retrieve stats of hero %d", ownerID)
			return
		}

		statsPairs, err := r.fetchStats(&heroStats, keys)
		if err != nil {
			logrus.WithError(err).Warn("rank.GetStatsForOwners: Cannot fetch stats for hero of ID %d", ownerID)
			return
		}

		stats = append(stats, statsContainer{
			OwnerID:   ownerID,
			OwnerType: 1,
			Stats:     statsPairs,
		})
	}

	r.answer(event.Client, event.Command.PayloadID, ansGetStatsForOwners{
		Txn:   rankGetStatsForOwners,
		Stats: stats,
	})
}
