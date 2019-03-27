package theater

import (
	"github.com/Synaxis/battlefieldHeroesBackend/backend/network"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/network/codec"
)

type reqEGEG struct {
	reqEGAM
}

type ansEGEG struct {
	TID           string `fesl:"TID"`
	Platform      string `fesl:"PL"`
	Ticket        string `fesl:"TICKET"`
	PlayerID      int    `fesl:"PID"`
	IP            string `fesl:"I"`
	Port          string `fesl:"P"`
	HostUserID    int    `fesl:"HUID"`
	EncryptionKey string `fesl:"EKEY"`
	// Alternatively to EKEY it is possible to use NOENCYRPTIONKEY
	NoEcryptionKey string `fesl:"NOENCYRPTIONKEY,omitempty"`
	IntIP          string `fesl:"INT-IP"`
	IntPort        string `fesl:"INT-PORT"`
	Secret         string `fesl:"SECRET,omitempty"`
	// Alternatively to SECRET it is possible to use NOSECRET
	NoSecret string `fesl:"NOSECRET,omitempty"`
	Ugid     string `fesl:"UGID,omitempty"`
	// Alternatively to UGID it is possible to use NOGUID
	NoGUID  string `fesl:"NOGUID,omitempty"`
	LobbyID int `fesl:"LID"`
	GameID  int    `fesl:"GID"`
}

// EGEG is sent Client to receive last confirmation before joining game
func (tm *Theater) EGEG(event *network.EventClientCommand, gameServer *network.Client, gr GameRequest) {
	event.Client.WriteEncode(&codec.Answer{
		Type: codec.ThtrEnterGameEntitleGame,
		Payload: ansEGEG{
			TID:           event.Command.Message["TID"],
			Platform:      "pc",
			Ticket:        gr.Ticket,
			PlayerID:      gr.PlayerID,
			IP:            gameServer.ServerData.Get("IP"),
			Port:          gameServer.ServerData.Get("PORT"),
			HostUserID:    gameServer.PlayerData.ServerID,
			EncryptionKey: "O65zZ2D2A58mNrZw1hmuJw%3d%3d",
			// NoEcryptionKey: "1",
			IntIP:   gameServer.ServerData.Get("INT-IP"),
			IntPort: gameServer.ServerData.Get("INT-PORT"),
			Secret:  "MargeSimpson",
			Ugid:    gameServer.ServerData.Get("UGID"),
			LobbyID: 1,
			GameID:  gr.GameID,
		},
	})
}
