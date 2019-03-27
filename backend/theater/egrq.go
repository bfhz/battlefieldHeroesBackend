package theater

import (
	"net"
	"strconv"

	"github.com/Synaxis/battlefieldHeroesBackend/backend/network"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/network/codec"
)

type reqEGRQ struct {
	reqEGAM
}

type ansEGRQ struct {
	TID      string `fesl:"TID"`
	Name     string `fesl:"NAME"`
	UserID   int    `fesl:"UID"`
	PlayerID int    `fesl:"PID"`
	Ticket   string `fesl:"TICKET"`
	IP       string `fesl:"IP"`
	Port     string `fesl:"PORT"`
	IntIP    string `fesl:"INT-IP"`
	IntPort  string `fesl:"INT-PORT"`
	// PTPE can be O or P
	Ptype        string `fesl:"PTYPE"`
	RUser        string `fesl:"R-USER"`
	RUid         int    `fesl:"R-UID"`
	RUAccid      int    `fesl:"R-U-accid"`
	RUElo        string `fesl:"R-U-elo"`
	RUTeam       string `fesl:"R-U-team"`
	RUKit        string `fesl:"R-U-kit"`
	RULvl        string `fesl:"R-U-lvl"`
	RUDataCenter string `fesl:"R-U-dataCenter"`
	RUExternalIP string `fesl:"R-U-externalIp"`
	RUInternalIP string `fesl:"R-U-internalIp"`
	RUCategory   string `fesl:"R-U-category"`
	RIntIP       string `fesl:"R-INT-IP"`
	RIntPort     string `fesl:"R-INT-PORT"`
	Xuid         string `fesl:"XUID"`
	RXuid        string `fesl:"R-XUID"`
	LobbyID      string `fesl:"LID"`
	GameID       int    `fesl:"GID"`
}

// EnterGameRequest (EGRQ) is sent to Server to inform about the player
// who wants join server
func (tm *Theater) EnterGameRequest(event *network.EventClientCommand, gameServer *network.Client, gr GameRequest) {
	externalIP := event.Client.IpAddr.(*net.TCPAddr).IP.String()

	gameServer.WriteEncode(&codec.Answer{
		Type: codec.ThtrEnterGameRequest,
		Payload: ansEGRQ{
			TID:          "0", //when we set this to 0 it "resets it"
			Name:         gr.HeroName,
			UserID:       gr.HeroID,
			PlayerID:     gr.PlayerID,
			Ticket:       gr.Ticket,
			IP:           externalIP,
			Port:         strconv.Itoa(event.Client.IpAddr.(*net.TCPAddr).Port),
			IntIP:        event.Command.Message["R-INT-IP"],
			IntPort:      event.Command.Message["R-INT-PORT"],
			Ptype:        "P",
			RUser:        gr.HeroName,
			RUid:         gr.HeroID,
			RUAccid:      gr.HeroID, //this need to be the same
			RUElo:        gr.Stats["elo"],
			RUTeam:       gr.Stats["c_team"],
			RUKit:        gr.Stats["c_kit"],
			RULvl:        gr.Stats["level"],
			RUDataCenter: network.RegionEastCoast,
			RUExternalIP: externalIP,
			RUInternalIP: event.Command.Message["R-INT-IP"],
			RUCategory:   event.Command.Message["R-U-category"],
			RIntIP:       event.Command.Message["R-INT-IP"],
			RIntPort:     event.Command.Message["R-INT-PORT"],
			Xuid:         "24",
			RXuid:        "24",
			LobbyID:      "1",
			GameID:       gr.GameID,
		},
	})
}
