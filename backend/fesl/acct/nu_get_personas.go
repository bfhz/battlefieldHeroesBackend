package acct

import (
	"github.com/sirupsen/logrus"

	"github.com/Synaxis/battlefieldHeroesBackend/backend/network"
)

type reqServerNuGetPersonas struct {
	// TXN=NuGetPersonas
	TXN string `fesl:"TXN"`
	// namespace=
	Namespace string `fesl:"namespace"`
}

type ansNuGetPersonas struct {
	Txn      string   `fesl:"TXN"`
	Personas []string `fesl:"personas"`
}

// NuGetPersonas handles acct.NuGetPersonas command
// NuGetPersonas - Soldier data lookup call
func (acct *Account) NuGetPersonas(event network.EventClientCommand) {
	switch event.Client.GetClientType() {
	case clientTypeServer:
		acct.serverNuGetPersonas(event)
	default:
		acct.clientNuGetPersonas(event)
	}
}

func (acct *Account) clientNuGetPersonas(event network.EventClientCommand) {
	hs, err := acct.DB.GetHeroesByPID(
		acct.DB.NewSession(),
		event.Client.PlayerData.PlayerID,
	)
	if err != nil {
		logrus.
			WithField("playerID", event.Client.PlayerData.PlayerID).
			Error("Cannot fetch any heroes of player")
		return
	}

	ans := ansNuGetPersonas{Txn: acctNuGetPersonas, Personas: []string{}}

	for _, h := range hs {
		ans.Personas = append(ans.Personas, h.HeroName)
	}

	acct.answer(event.Client, event.Command.PayloadID, ans)
}

func (acct *Account) serverNuGetPersonas(event network.EventClientCommand) {
	srv, err := acct.DB.GetServerByID(
		acct.DB.NewSession(),
		event.Client.PlayerData.ServerID,
	)
	if err != nil {
		logrus.WithError(err).Errorf("acct: Cannot find server with ID '%d'", event.Client.PlayerData.ServerID)
		return
	}

	event.Client.PlayerData.ServerID = srv.ID
	event.Client.PlayerData.ServerUserName = srv.AccountUsername

	acct.answer(
		event.Client,
		event.Command.PayloadID,
		ansNuGetPersonas{
			Txn:      acctNuGetPersonas,
			Personas: []string{srv.AccountUsername},
		},
	)
}
