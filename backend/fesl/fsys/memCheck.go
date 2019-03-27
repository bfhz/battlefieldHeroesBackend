package fsys

import (
	"github.com/Synaxis/battlefieldHeroesBackend/backend/network"
)

type reqMemCheck struct {
	// TXN stands for Taxon, sub-query name of the cmd.
	// Should be equal: MemCheck.
	TXN string `fesl:"TXN"`

	// FIXME: Result is usually an empty string
	Result string `fesl:"result"`
}

type ansMemCheck struct {
	// TXN stands for Taxon, sub-query name of the cmd.
	// Should be equal: MemCheck.
	TXN string `fesl:"TXN"`

	MemChecks []memCheck `fesl:"memcheck"`
	Salt      string     `fesl:"salt"`
	Type      int        `fesl:"type"`
}

type memCheck struct {
	Addr   string `fesl:"addr"`
	Length int    `fesl:"len"`
}

// MemCheck handles fsys.MemCheck cmd.
//
// defined in fsys.Hello - activityTimeoutSecs).
func (fsys *ConnectSystem) MemCheck(client *network.Client) {
	fsys.answer(client, 0xC0000000, ansMemCheck{
		TXN:  fsysMemCheck,
		Salt: "",
	})
}
