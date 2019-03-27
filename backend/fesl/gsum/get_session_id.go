package gsum

import (
	"github.com/Synaxis/battlefieldHeroesBackend/backend/network"
)

type ansGetSessionID struct {
	Txn string `fesl:"TXN"`
	// Games  []Game  `fesl:"games"`
	// Events []Event `fesl:"events"`
}

// GetSessionID handles gsum.GetSessionID command
func (gsum *GameSummary) GetSessionID(ev network.EventClientCommand) {
	gsum.answer(ev.Client, 0, ansGetSessionID{
		Txn: gsumGetSessionID,
	})
}
