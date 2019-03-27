package acct

import (
	"github.com/sirupsen/logrus"

	"github.com/Synaxis/battlefieldHeroesBackend/backend/network"
)

type reqNuLoginPersona struct {
	// TXN=NuLoginPersona
	Txn  string `fesl:"TXN"`
	Name string `fesl:"name"` // Value specified in +soldierName
}

type ansNuLoginPersona struct {
	Txn       string `fesl:"TXN"`
	ProfileID int    `fesl:"profileId"`
	UserID    int    `fesl:"userId"`
	LobbyKey  string `fesl:"lkey"`
}

// NuLoginPersona handles acct.NuLoginPersona command
func (acct *Account) NuLoginPersona(event network.EventClientCommand) {
	switch event.Client.GetClientType() {
	case clientTypeServer:
		acct.serverNuLoginPersona(event)
	default:
		acct.clientNuLoginPersona(event)
	}
}

// clientNuLoginPersona used when user selects hero from game client
func (acct *Account) clientNuLoginPersona(event network.EventClientCommand) {
	heroName := event.Command.Message["name"]

	h, err := acct.DB.GetHeroByName(acct.DB.NewSession(), heroName)
	if err != nil {
		logrus.WithError(err).Warnf("Cannot find hero for name '%s'", heroName)
		return
	}

	event.Client.PlayerData.HeroID = h.ID
	event.Client.PlayerData.PlayerID = h.PlayerID
	event.Client.PlayerData.HeroName = h.HeroName

	acct.answer(
		event.Client,
		event.Command.PayloadID,
		ansNuLoginPersona{
			Txn:       acctNuLoginPersona,
			ProfileID: event.Client.PlayerData.PlayerID,
			UserID:    event.Client.PlayerData.HeroID,
			LobbyKey:  event.Client.PlayerData.LobbyKey,
		},
	)
}

// NuLoginPersonaServer - soldier login command
func (acct *Account) serverNuLoginPersona(event network.EventClientCommand) {
	a, err := acct.DB.GetServerByName(
		acct.DB.NewSession(),
		event.Command.Message["name"],
	)
	if err != nil {
		logrus.WithError(err).Warn("Cannot find server")
		return
	}

	acct.answer(
		event.Client,
		event.Command.PayloadID,
		ansNuLoginPersona{
			Txn:       acctNuLoginPersona,
			ProfileID: a.ID,
			UserID:    a.ID,
			LobbyKey:  event.Client.PlayerData.LobbyKey,
		},
	)
}
