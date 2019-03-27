package acct

import (
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"

	"github.com/Synaxis/battlefieldHeroesBackend/backend/network"
)

type reqNuLogin struct {
	// TXN=NuLogin
	TXN string `fesl:"TXN"`
	// returnEncryptedInfo=0
	ReturnEncryptedInfo int `fesl:"returnEncryptedInfo"`
	// macAddr=$0a0027000000
	MacAddr string `fesl:"macAddr"`
}

type reqNuLoginServer struct {
	reqNuLogin

	AccountName     string `fesl:"nuid"`     // Value specified in +eaAccountName
	AccountPassword string `fesl:"password"` // Value specified in +eaAccountPassword
}

type reqNuLoginClient struct {
	reqNuLogin

	EncryptedInfo string `fesl:"encryptedInfo"` // Value specified in +sessionId
}

type ansNuLogin struct {
	Txn       string `fesl:"TXN"`
	ProfileID int    `fesl:"profileId"`
	UserID    int    `fesl:"userId"`
	NucleusID int    `fesl:"nuid"`
	LobbyKey  string `fesl:"lkey"`
}

type ansNuLoginErr struct {
	Txn     string                `fesl:"TXN"`
	Message string                `fesl:"localizedMessage"`
	Errors  []nuLoginContainerErr `fesl:"errorContainer"`
	Code    int                   `fesl:"errorCode"`
}

type nuLoginContainerErr struct {
	Value      string `fesl:"value"`
	FieldError string `fesl:"fieldError"`
	FieldName  string `fesl:"fieldName"`
}

//if you think this is difficult go learn c++
func randStr() (rand string, retErr error) {
	newRandom := uuid.NewV4().String()
	rand = newRandom
	return rand, nil
}


// NuLogin handles acct.NuLogin command
func (acct *Account) NuLogin(event network.EventClientCommand) {
	lkey, _ := randStr()

	if lkey ==  "" {
		logrus.Println("error with lkey convert")
	}
	
	event.Client.PlayerData.LobbyKey = lkey
	err := network.Lobby.Add(lkey, event.Client.PlayerData)
	if err != nil {
		logrus.WithError(err).Warn("Cannot add ClientData in acct.NuLogin")
		return
	}

	switch event.Client.GetClientType() {
	case clientTypeServer:
		acct.serverNuLogin(event)
	default:
		acct.clientNuLogin(event)
	}
}

func (acct *Account) clientNuLogin(event network.EventClientCommand) {
	player, err := acct.DB.GetPlayerByToken(
		acct.DB.NewSession(),
		event.Command.Message["encryptedInfo"],
	)
	if err != nil {
		logrus.WithError(err).Warn("Client cannot sign in the acct.NuLogin")
		acct.clientNuLoginNotAuthorized(&event)
		return
	}

	event.Client.PlayerData.PlayerID = player.ID

	acct.answer(
		event.Client,
		event.Command.PayloadID,
		ansNuLogin{
			Txn:       acctNuLogin,
			UserID:    event.Client.PlayerData.PlayerID,
			ProfileID: event.Client.PlayerData.PlayerID,
			NucleusID: event.Client.PlayerData.PlayerID,
			LobbyKey:  event.Client.PlayerData.LobbyKey,
		},
	)
}

func (acct *Account) clientNuLoginNotAuthorized(event *network.EventClientCommand) {
	acct.answer(
		event.Client,
		event.Command.PayloadID,
		ansNuLoginErr{
			Txn:     acctNuLogin,
			Message: `"The user is not entitled to access this game"`,
			Code:    120,
		},
	)
}

// acctNuLoginServer - login command for servers
func (acct *Account) serverNuLogin(event network.EventClientCommand) {
	srv, err := acct.DB.GetServerLogin(
		acct.DB.NewSession(),
		event.Command.Message["nuid"],
	)
	if err != nil {
		logrus.WithError(err).Warn("Server cannot sign in the acct.NuLogin")
		acct.serverNuLoginNotAuthorized(&event)
		return
	}

	// TODO: Raw passwords are really, really insecure
	// 1. Validate credentials using some bcrypt or argon2
	// 2. DO NOT STORE raw passwords
	if srv.AccountPassword != event.Command.Message["password"] {
		acct.serverNuLoginNotAuthorized(&event)
		return
	}

	event.Client.PlayerData.ServerID = srv.ID
	event.Client.PlayerData.ServerSoldierName = srv.SoldierName
	event.Client.PlayerData.ServerUserName = srv.AccountUsername

	acct.answer(
		event.Client,
		event.Command.PayloadID,
		ansNuLogin{
			Txn:       acctNuLogin,
			ProfileID: srv.ID,
			UserID:    srv.ID,
			NucleusID: srv.ID,
			LobbyKey:  event.Client.PlayerData.LobbyKey,
		},
	)
}

func (acct *Account) serverNuLoginNotAuthorized(event *network.EventClientCommand) {
	acct.answer(
		event.Client,
		event.Command.PayloadID,
		ansNuLoginErr{
			Txn:     acctNuLogin,
			Message: `"The password the user specified is incorrect"`,
			Code:    122,
		},
	)
}
