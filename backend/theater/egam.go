package theater

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/Synaxis/battlefieldHeroesBackend/backend/network"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/network/codec"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/ranking"
)

// EGAM is sent to Game-Client
type reqEGAM struct {
	// GID=1
	GameID int `fesl:"GID"`
	// LID=1
	LobbyID int `fesl:"LID"`
	// PORT=54671
	Port int `fesl:"PORT"`
	// PTYPE=P
	PlatformType int `fesl:"PTYPE"`
	// R-INT-IP=192.168.0.101
	RemoteIP string `fesl:"R-INT-IP"`
	// R-INT-PORT=54671
	RemotePort int `fesl:"R-INT-PORT"`
	// R-U-accid=2
	AccountID int `fesl:"R-U-accid"` // TODO: Hero or PlayerID? PlayerID :(
	// R-U-category=3
	Category int `fesl:"R-U-category"` // TODO: What exactly it is?
	// R-U-dataCenter=iad
	Region string `fesl:"R-U-dataCenter"`
	// R-U-elo=1000
	StatsElo int `fesl:"R-U-elo"`
	// R-U-externalIp=127.0.0.1
	ExternalIP string `fesl:"R-U-externalIp"`
	// R-U-kit=0
	StatsKit int `fesl:"R-U-kit"`
	// R-U-lvl=1
	StatsLevel int `fesl:"R-U-lvl"`
	// R-U-team=1
	StatsTeam int `fesl:"R-U-team"`
	// TID=4
	TID int `fesl:"TID"`
}

type ansEGAM struct {
	TID     string `fesl:"TID"`
	LobbyID string `fesl:"LID"`
	GameID  int    `fesl:"GID"`
}

// a EGAM - CLIENT called when a client wants to join a gameserver
func (tm *Theater) EnterGame(event network.EventClientCommand) {
	gameID, err := event.Command.Message.IntVal("GID")
	if err != nil {
		logrus.WithError(err).Warn("Cannot parse value of GID in theater.EGAM")
		return
	}

	game, err := tm.mm.GetGame(gameID)
	if err != nil {
		logrus.
			WithError(err).
			WithField("gameID", gameID).
			Warn("Not found any server when joining game")
		return
	}

	event.Client.WriteEncode(&codec.Answer{
		Type: codec.ThtrEnterGame,
		Payload: ansEGAM{
			event.Command.Message["TID"],
			event.Command.Message["LID"],
			gameID,
		},
	})

	heroStats, err := tm.db.GetHeroStats(tm.db.NewSession(), event.Client.PlayerData.HeroID)
	if err != nil {
		logrus.
			WithError(err).
			WithField("heroID", event.Client.PlayerData.HeroID).
			Warn("Cannot fetch stats for hero when entering a game")
		return
	}

	// TODO: Validate if passed values are correct - protect against -1 level
	// stats := map[string]string{
	// 	"c_kit":  event.Command.Message["R-U-kit"],
	// 	"c_team": event.Command.Message["R-U-team"],
	// 	"elo":    event.Command.Message["R-U-elo"],
	// 	"level":  event.Command.Message["R-U-lvl"],
	// }

	stats, err := ranking.GetStats(&heroStats, "c_kit", "c_team", "elo", "level")
	if err != nil {
		logrus.
			WithError(err).
			WithField("heroID", event.Client.PlayerData.HeroID).
			Warn("Cannot get stats for hero when entering a game")
		return
	}

	ticket := generateTicket()

	gr := GameRequest{
		// PlayerID: 2,
		PlayerID: event.Client.PlayerData.PlayerID,
		// HeroID: 2,
		HeroID: event.Client.PlayerData.HeroID,
		// HeroName: "FirstHero",
		HeroName: event.Client.PlayerData.HeroName,
		GameID:   game.ID,
		LobbyID:  event.Command.Message["LID"],
		Ticket:   ticket,
		Stats:    stats,
	}

	tm.EnterGameRequest(&event, game.GameServer, gr)
	tm.EGEG(&event, game.GameServer, gr)
}

type GameRequest struct {
	PlayerID int
	HeroID   int
	HeroName string
	GameID   int
	LobbyID  string
	Stats    ranking.Stats
	Ticket   string
}

var (
	lastTicket = 10005
	ticketMu   sync.Mutex
)

func generateTicket() string {
	defer ticketMu.Unlock()
	ticketMu.Lock()
	lastTicket++
	return fmt.Sprintf("%d", lastTicket)
}
