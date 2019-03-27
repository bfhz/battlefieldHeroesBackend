package fsys

import (
	"github.com/Synaxis/battlefieldHeroesBackend/backend/network"
)

type ansGetPingSites struct {
	// TXN stands for Taxon, sub-query name of the command.
	// Should be equal: GetPingSites.
	TXN string `fesl:"TXN"`

	// MinPings defines at least how many endpoints should pinged to calculate
	// the ping correctly.
	MinPings int `fesl:"minPingSitesToPing"`

	// PingSites defines a list of endpoints, which should be pinged,
	// accordiningly to minPingSitesToPing setting.
	PingSites []pingSite `fesl:"pingSites"`
}

type pingSite struct {
	Addr string `fesl:"addr"`
	Name string `fesl:"name"`
	Type int    `fesl:"type"`
}

var pingSites = []pingSite{
	{"26.49.251.103", network.RegionAsia, 0},
	{"26.49.251.103", network.RegionEurope, 0},
	{"26.49.251.103", network.RegionEastCoast, 0},
	{"26.49.251.103", network.RegionWestCoast, 0},
}

// GetPingSites handles fsys.GetPingSites command.
func (fsys *ConnectSystem) GetPingSites(event network.EventClientCommand) {
	fsys.answer(
		event.Client,
		event.Command.PayloadID,
		ansGetPingSites{
			TXN:       fsysGetPingSites,
			MinPings:  0,
			PingSites: pingSites,
		},
	)
}
