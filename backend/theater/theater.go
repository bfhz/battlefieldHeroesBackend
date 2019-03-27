package theater

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/Synaxis/battlefieldHeroesBackend/backend/matchmaking"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/network"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/storage/database"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/storage/kvstore"
)

var (
	// thtrGLST = "GLST"
	thtrKICK = "KICK"
	thtrPLVT = "PLVT"
	thtrUBRA = "UBRA"
)

// Theater Handles incoming and outgoing theater communication
type Theater struct {
	db        database.Adapter
	mm        *matchmaking.Pool
	socket    *network.Socket
	socketUDP *network.SocketUDP
	cmds      network.CmdRegistry
}

// New creates and starts a new TheaterManager
func New(bind string, db database.Adapter, kvs *kvstore.Storage, mm *matchmaking.Pool) *Theater {
	socket, err := network.NewSocketTCP(bind)
	if err != nil {
		logrus.Fatal(err)
		return nil
	}

	socketUDP, err := network.NewSocketUDP(bind)
	if err != nil {
		logrus.Fatal(err)
		return nil
	}

	tm := &Theater{
		mm:        mm,
		db:        db,
		socket:    socket,
		socketUDP: socketUDP,
	}

	const (
		numberOfRegisteredCommands = 15
	)
	r := make(network.CmdRegistry, numberOfRegisteredCommands)

	r.Register("newClient", func(ev network.EventClientCommand) {
		// Start Heartbeat
		ev.Client.HeartTicker = time.NewTicker(time.Second * 55)
		go func() {
			for ev.Client.IsActive {
				select {
				case <-ev.Client.HeartTicker.C:
					if !ev.Client.IsActive {
						return
					}
					tm.PING(ev.Client)
				}
			}
		}()
	})
	r.Register("client.command.CONN", tm.Connect)
	r.Register("client.command.USER", tm.Login)
	r.Register("client.command.LLST", func(ev network.EventClientCommand) {
		tm.GetLobbyList(ev)
		tm.LobbyData(ev)
	})
	r.Register("client.command.GDAT", tm.GameData)
	r.Register("client.command.EGAM", tm.EnterGame)
	r.Register("client.command.ECNL", tm.EnterConnectionLAN)
	r.Register("client.command.CGAM", tm.CreateGame)
	r.Register("client.command.UBRA", tm.UpdateBracket)
	r.Register("client.command.UGAM", tm.UpdateGameData)
	r.Register("client.command.EGRS", tm.EnterGameHostResponse)
	r.Register("client.command.PENT", tm.PlayerEntered)
	r.Register("client.command.PLVT", tm.PlayerExited)
	r.Register("client.command.UPLA", tm.UpdatePlayerData)
	r.Register("client.command.PING", func(ev network.EventClientCommand) {
		// TODO: Use metrics in the response and save it for later use.
		return
	})

	tm.cmds = r

	return tm
}

func (tm *Theater) ListenAndServe(ctx context.Context) {
	go tm.Run(ctx)
}

func (tm *Theater) Run(ctx context.Context) {
	for {
		select {
		case event := <-tm.socketUDP.EventChan:
			tm.handleUDP(event)
		case event := <-tm.socket.EventChan:
			tm.handleTLS(event)
		case <-ctx.Done():
			return
		}
	}
}

func (tm *Theater) handleUDP(event network.SocketUDPEvent) {
	switch event.Name {
	case "ECHO":
		tm.ECHO(event)
	default:
		logrus.
			WithFields(logrus.Fields{
				"event": event.Name,
				"data":  event.Data,
			}).
			Warn("theater.UnhandledRequest (UDP)")
	}
}

func (tm *Theater) handleTLS(event network.SocketEvent) {
	ev, ok := event.Data.(network.EventClientCommand)
	if !ok {
		logrus.Error("Logic error: Cannot cast event to network.EventClientCommand")
		return
	}

	if !ev.Client.IsActive {
		logrus.WithField("command", ev.Command).Warn("Inactive client")
		return
	}

	fn, ok := tm.cmds.Find(event.Name)
	if !ok {
		logrus.
			WithFields(logrus.Fields{
				"event":   event.Name,
				"query":   ev.Command.Query,
				"payload": ev.Command.Message,
			}).
			Warn("theater.UnhandledRequest")
		return
	}

	fn(ev)
}
