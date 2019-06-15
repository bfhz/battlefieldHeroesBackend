package internal_test

import (
	"bytes"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/Synaxis/battlefieldHeroesBackend/backend/network"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/network/codec"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/fesl/acct"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/fesl/fsys"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/fesl/gsum"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/fesl/rank"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/storage/database/dbtest"
)

type fakeConn struct {
	buf *bytes.Buffer
}

func (c *fakeConn) Read(b []byte) (n int, err error) { return 0, nil }
func (c *fakeConn) Write(b []byte) (n int, err error) {
	fmt.Printf("Wrote to client: %s\n", b)
	return len(b), nil
}
func (c *fakeConn) Close() error                       { return nil }
func (c *fakeConn) LocalAddr() net.Addr                { return nil }
func (c *fakeConn) RemoteAddr() net.Addr               { return nil }
func (c *fakeConn) SetDeadline(t time.Time) error      { return nil }
func (c *fakeConn) SetReadDeadline(t time.Time) error  { return nil }
func (c *fakeConn) SetWriteDeadline(t time.Time) error { return nil }

func createEventCommand(client *network.Client, payload codec.Fields) network.EClientCmd {
	return network.EClientCmd{
		Client: client,
		Cmd: &codec.Cmd{
			Msg: payload,
		},
	}
}

func createCommand(payload []byte) *codec.Cmd {
	return &codec.Cmd{Msg: codec.DecodeFESL(payload)}
}

func TestServerLogin(t *testing.T) {
	ipAddr, _ := net.ResolveIPAddr("", "127.0.0.1")
	client := &net.Client{
		IpAddr:   ipAddr,
		IsActive: true,
		Conn:     &fakeConn{buf: new(bytes.Buffer)},
	}

	db := dbtest.NewFakeDB()

	fsysCtrl := fsys.ConnectSystem{true}
	acctCtrl := acct.Account{db}

	fsysCtrl.Hello(client, &codec.Cmd{Msg: codec.Fields{
		"TXN":             `Hello`,
		"clientString":    `bfwest-pc`,
		"sku":             `125170`,
		"locale":          `en_US`,
		"clientPlatform":  `PC`,
		"clientVersion":   `1.42.217478.0`,
		"SDKVersion":      `5.0.0.0.0`,
		"protocolVersion": `2.0`,
		"fragmentSize":    `8096`,
		"clientType":      `server`,
	}})

	acctCtrl.NuLogin(createEventCommand(client, codec.Fields{
		"TXN":                 `NuLogin`,
		"returnEncryptedInfo": `0`,
		"nuid":                ``,
		"password":            `MachoVirgem`,
		"macAddr":             `$0a0027000000`,
	}))

	acctCtrl.NuGetPersonas(createEventCommand(client, codec.Fields{
		"TXN":       `NuGetPersonas`,
		"namespace": ``,
	}))

	acctCtrl.NuLoginPersona(createEventCommand(client, codec.Fields{
		"TXN":  `NuLoginPersona`,
		"name": `MachoVirgem`,
	}))

	acctCtrl.NuGetAccount(createEventCommand(client, codec.Fields{
		"TXN": `NuGetAccount`,
	}))

	acctCtrl.NuLookupUserInfo(createEventCommand(client, codec.Fields{
		"TXN":                 `NuLookupUserInfo`,
		"userInfo.[]":         `1`,
		"userInfo.0.userName": `Test-Server`,
	}))

}

func TestClientLogin(t *testing.T) {
	ipAddr, _ := net.ResolveIPAddr("", "127.0.0.1")
	client := &net.Client{
		IpAddr:   ipAddr,
		IsActive: true,
		Conn:     &fakeConn{buf: new(bytes.Buffer)},
	}

	// kvs := storage.NewInMem()
	db := dbtest.NewFakeDB()

	fsysCtrl := fsys.ConnectSystem{false}
	gsumCtrl := gsum.GameSummary{}
	acctCtrl := acct.Account{db}
	rankCtrl := rank.Ranking{db}

	fsysHello := &codec.Cmd{Msg: codec.Fields{
		"TXN":             `Hello`,
		"clientString":    `bfwest-pc`,
		"sku":             `125170`,
		"locale":          `en_US`,
		"clientPlatform":  `PC`,
		"clientVersion":   `"1.42.217478.0 "`,
		"SDKVersion":      `5.0.0.0.0`,
		"protocolVersion": `2.0`,
		"fragmentSize":    `8096`,
		"clientType":      `client-noreg`,
	}}
	gsumCtrl.GetSessionID(client, fsysHello)
	fsysCtrl.Hello(client, fsysHello)

	// GET https://heroes-api/nucleus/authToken HTTP/1.1
	// GET https://heroes-api/en/products/offers HTTP/1.1

	acctCtrl.NuLogin(createEventCommand(
		client,
		codec.Fields{
			"TXN":                 `NuLogin`,
			"returnEncryptedInfo": `0`,
			"encryptedInfo":       `JailsonMendessenha`,
			"macAddr":             `$0a0027000000`,
		},
	))

	acctCtrl.NuGetPersonas(createEventCommand(
		client,
		codec.Fields{
			"TXN":       `NuGetPersonas`,
			"namespace": ``,
		},
	))

	// GET https://heroes-api/nucleus/check/user/{userID:2}

	acctCtrl.NuGetAccount(createEventCommand(
		client,
		codec.Fields{"TXN": "NuGetAccount"},
	))

	// owner=AuthPlayer
	rankCtrl.GetStats(createEventCommand(
		client,
		codec.Fields{
			"TXN":        `GetStats`,
			"owner":      `2`,
			"ownerType":  `1`,
			"periodId":   `0`,
			"periodPast": `0`,
			"keys.0":     `c_ltm`,
			"keys.1":     `c_slm`,
			"keys.2":     `c_tut`,
			"keys.[]":    `3`,
		},
	))

	// If there are any Heroes responded in NuGetPersonas,
	// then they will be queries here
	acctCtrl.NuLookupUserInfo(createEventCommand(
		client,
		codec.Fields{
			"TXN":                 `NuLookupUserInfo`,
			"userInfo.[]":         `1`,
			"userInfo.0.userName": `TestPlayer`,
		},
	))

	// Once again for owner=PlayerHeroID
	rankCtrl.GetStats(createEventCommand(
		client,
		codec.Fields{
			"TXN":        `GetStats`,
			"owner":      `5`,
			"ownerType":  `1`,
			"periodId":   `0`,
			"periodPast": `0`,
			"keys.0":     `c_apr`,
			"keys.1":     `c_fhrs`,
			"keys.2":     `c_ft`,
			"keys.3":     `c_hrc`,
			"keys.4":     `c_hrs`,
			"keys.5":     `c_kit`,
			"keys.6":     `c_skc`,
			"keys.7":     `c_team`,
			"keys.8":     `elo`,
			"keys.9":     `level`,
			"keys.10":    `xp`,
			"keys.[]":    `11`,
		},
	))

	// When player selects hero in GUI
	acctCtrl.NuLoginPersona(createEventCommand(
		client,
		codec.Fields{
			"TXN":  `NuLoginPersona`,
			"name": `TestPlayer`,
		},
	))

	fmt.Println(client.PlayerData)

	// fsysCtrl.GetPingSites()

	// map[PLAT:PC LOCALE:en_US SDKVERSION:5.0.0.0.0 TID:1 PROT:2 PROD:bfwest-pc VERS:\"1.42.217478.0 \"]
	// thtr.CONN()

	// map[NAME: TID:2 CID: MAC:$0a0027000004 SKU:125170 LKEY:]
	// thtr.USER()

	// ECHO map[TID:1 TYPE:1 UID:2]
}
