package theater

import (
	"github.com/Synaxis/battlefieldHeroesBackend/backend/network"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/network/codec"
)

// Lobbies List
type ansLLST struct {
	TID        string `fesl:"TID"`
	NumLobbies int    `fesl:"NUM-LOBBIES"`
}

func (tm *Theater) GetLobbyList(event network.EventClientCommand) {
	event.Client.WriteEncode(&codec.Answer{
		Type:    codec.ThtrLobbyList,
		Payload: ansLLST{event.Command.Message["TID"], 1},
	})
}
