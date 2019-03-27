package theater

import (
	"github.com/sirupsen/logrus"

	"github.com/Synaxis/battlefieldHeroesBackend/backend/network"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/network/codec"
)

type reqGDAT struct {
	// TID=3
	TID int `fesl:"TID"`

	// LID=0
	LobbyID string `fesl:"LID"`
	// GID=1
	GameID int `fesl:"GID"`
}

type ansGDAT struct {
	TID string `fesl:"TID"`

	EloRank             string `fesl:"B-U-elo_rank"`
	AvgAllyRank         string `fesl:"B-U-avg_ally_rank"`
	AvgAxisRank         string `fesl:"B-U-avg_axis_rank"`
	ArmyDistribution    string `fesl:"B-U-army_distribution"`
	ArmyBalance         string `fesl:"B-U-army_balance"`
	PercentFull         string `fesl:"B-U-percent_full"`
	AvailSlotsNational  string `fesl:"B-U-avail_slots_national"`
	AvailSlotsRoyal     string `fesl:"B-U-avail_slots_royal"`
	AvailableVipsNation string `fesl:"B-U-avail_vips_national"`
	AvailableVipsRoyal  string `fesl:"B-U-avail_vips_royal"`
	IsRanked            string `fesl:"B-U-ranked"`
	Easyzone            string `fesl:"B-U-easyzone"`
	ServerType          string `fesl:"B-U-servertype"`
	ServerState         string `fesl:"B-U-server_state"`
	MapName             string `fesl:"B-U-map_name"`
	PunkBuster          string `fesl:"B-U-punkb"`
	StdDevLevel         string `fesl:"B-U-lvl_sdv"`
	AvgLevel            string `fesl:"B-U-lvl_avg"`
	IntIp               string `fesl:"INT-IP"`
	IntPort             string `fesl:"INT-PORT"`
	IP                  string `fesl:"IP"`
	MaxPlayers          string `fesl:"MAX-PLAYERS"`
	Port                string `fesl:"PORT"`
	ugid                string `fesl:"UGID"`
	stype               string `fesl:"TYPE"`
	hxfr                string `fesl:"HXFR"`
	httype              string `fesl:"HTTYPE"`
	disablequeue        string `fesl:"DISABLE-AUTO-DEQUEUE"`
	rt                  string `fesl:"RT"`
	qlen                string `fesl:"QLEN"`
	queuelength         string `fesl:"QUEUE-LENGTH"`
	reservehost         string `fesl:"RESERVE-HOST"`
	secret              string `fesl:"SECRET"`
	joinmode            string `fesl:"J"`
	umap                string `fesl:"B-U-map"`
	version             string `fesl:"B-version"`
	servertype          string `fesl:"B-U-servertype"`
	datacenter          string `fesl:"B-U-data_center"`
	comName             string `fesl:"B-U-community_name"`

	GameID     string `fesl:"GID"`
	Join       string `fesl:"JOIN"`
	ServerName string `fesl:"NAME"`

	AP      string `fesl:"AP"` // KitCount, int
	LobbyID int    `fesl:"LID"`

	// GameType          string `fesl:"TYPE"` // = "P"
	// IpAddr string `fesl:"I,omitempty"` //string
	// Port string `fesl:"P,omitempty"` // int
	// Password string `fesl:"PW,omitempty"` // string
	// MaxPlayersCount string `fesl:"MP,omitempty"` // KitCapacity, int
	// ActualPlayersCount    string `fesl:"AP"` // KitCount, int
	// PlayerMaxObservers string `fesl:"B-maxObservers,omitempty"` // int
	// PlayerActualObservers string `fesl:"B-numObservers,omitempty"` // int
	// JoiningPlayersCount string `fesl:"JP,omitempty"` // int
	// QueuedPlayersCount string `fesl:"QP,omitempty"` // int
	// HostPlayerName string `fesl:"HN,omitempty"` // string
	// HostPlayerUserID string `fesl:"HU,omitempty"` // int
	// Version string `fesl:"V,omitempty"`
	// GameProtocolVersion string `fesl:"B-version,omitempty"`
	// IsFavorite string `fesl:"F,omitempty"` // int
	// FavPlayerCount string `fesl:"NF,omitempty"` // int
	// Platform string `fesl:"PL,omitempty"` // if PL=XBOX then XUID (XboxUserID must be specified)
	// JoinMode string `fesl:"J,omitempty"` // O / W / C (default=O, but if there will some random string not eaual to W and C it will also work)
}

// GDAT - CLIENT called to get data about the server
func (tm *Theater) GameData(event network.EventClientCommand) {
	gameID, err := event.Command.Message.IntVal("GID")
	if err != nil {
		logrus.WithError(err).Warn("Cannot parse GID in theater.GDAT")
		return
	}

	game, err := tm.mm.GetGame(gameID)
	if err != nil {
		logrus.
			WithError(err).
			WithField("gameID", gameID).
			Warn("Cannot find Game in matchmaking pool")
		return
	}
	gSData := game.GameServer.ServerData

	event.Client.WriteEncode(&codec.Answer{
		Type: codec.ThtrGamesData,
		Payload: ansGDAT{
			//LobbyID:             game.LobbyID,
			LobbyID: 1,
			AP:      gSData.Get("AP"),
			TID:     event.Command.Message["TID"],
			//GameID:              gSData.Get("GID"), //was gr.game.ID
			GameID:              gSData.Get("GID"),
			IntIp:               gSData.Get("INT-IP"),
			IntPort:             gSData.Get("INT-PORT"),
			version:             gSData.Get("B-version"),
			umap:                gSData.Get("B-U-map"),
			secret:              gSData.Get("SECRET"),
			comName:             gSData.Get("B-U-community_name"),
			datacenter:          gSData.Get("B-U-data_center"),
			servertype:          gSData.Get("B-U-servertype"),
			reservehost:         gSData.Get("RESERVE-HOST"),
			queuelength:         gSData.Get("QUEUE-LENGTH"),
			qlen:                gSData.Get("QLEN"),
			rt:                  gSData.Get("RT"),
			disablequeue:        gSData.Get("DISABLE-AUTO-DEQUEUE"),
			httype:              gSData.Get("HTTYPE"),
			hxfr:                gSData.Get("HXFR"),
			IP:                  gSData.Get("IP"),
			MaxPlayers:          gSData.Get("MAX-PLAYERS"),
			Port:                gSData.Get("PORT"),
			Join:                gSData.Get("JOIN"),
			ServerName:          gSData.Get("NAME"),
			stype:               gSData.Get("TYPE"),
			ugid:                gSData.Get("UGID"),
			EloRank:             gSData.Get("B-U-elo_rank"),
			AvgAllyRank:         gSData.Get("B-U-avg_ally_rank"),
			AvgAxisRank:         gSData.Get("B-U-avg_axis_rank"),
			ArmyDistribution:    gSData.Get("B-U-army_distribution"),
			ArmyBalance:         gSData.Get("B-U-army_balance"),
			PercentFull:         gSData.Get("B-U-percent_full"),
			AvailSlotsNational:  gSData.Get("B-U-avail_slots_national"),
			AvailSlotsRoyal:     gSData.Get("B-U-avail_slots_royal"),
			AvailableVipsNation: gSData.Get("B-U-avail_vips_national"),
			AvailableVipsRoyal:  gSData.Get("B-U-avail_vips_royal"),
			IsRanked:            gSData.Get("B-U-ranked"),
			Easyzone:            gSData.Get("B-U-easyzone"),
			ServerType:          gSData.Get("B-U-servertype"),
			ServerState:         gSData.Get("B-U-server_state"),
			PunkBuster:          gSData.Get("B-U-punkb"),
			MapName:             gSData.Get("B-U-map_name"),
			AvgLevel:            gSData.Get("B-U-lvl_avg"),
			StdDevLevel:         gSData.Get("B-U-lvl_sdv"),
		},
	})
}
