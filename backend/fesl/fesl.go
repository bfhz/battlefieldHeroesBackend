package fesl

import (
	"context"
	"time"

	"github.com/Synaxis/battlefieldHeroesBackend/backend/fesl/acct"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/fesl/fsys"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/fesl/gsum"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/fesl/pnow"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/fesl/rank"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/matchmaking"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/network"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/storage/database"

	"github.com/sirupsen/logrus"
)

// Fesl - handles incoming and outgoing FESL data
type Fesl struct {
	socket *network.Socket
	cmds   network.CmdRegistry
}

// New create new Fesl
func New(bind string, server bool, db database.Adapter, mm *matchmaking.Pool) *Fesl {
	socket, err := network.NewSocketTLS(bind)
	if err != nil {
		logrus.Fatal(err)
		return nil
	}

	const (
		numberOfRegisteredCommands = 14
	)
	r := make(network.CmdRegistry, numberOfRegisteredCommands)
	{
		acct := &acct.Account{DB: db}
		fsys := &fsys.ConnectSystem{ServerMode: server}
		gsum := &gsum.GameSummary{}
		pnow := &pnow.PlayNow{MM: mm}
		rank := &rank.Ranking{DB: db}

		r.Register("newClient", func(ev network.EventClientCommand) {
			// TLS only
			fsys.MemCheck(ev.Client)

			ev.Client.HeartTicker = time.NewTicker(55 * time.Second)
			go func() {
				for ev.Client.IsActive {
					select {
					case <-ev.Client.HeartTicker.C:
						if !ev.Client.IsActive {
							return
						}
						fsys.MemCheck(ev.Client)
					}
				}
			}()

			logrus.Debug("New client has connected")
		})
		r.Register("client.command.Hello", func(ev network.EventClientCommand) {
			if !server {
				gsum.GetSessionID(ev)
			}
			fsys.Hello(ev)
		})
		r.Register("client.command.NuLogin", acct.NuLogin)
		r.Register("client.command.NuGetPersonas", acct.NuGetPersonas)
		r.Register("client.command.NuGetAccount", acct.NuGetAccount)
		r.Register("client.command.NuLoginPersona", acct.NuLoginPersona)
		r.Register("client.command.GetStatsForOwners", rank.GetStatsForOwners)
		r.Register("client.command.GetStats", rank.GetStats)
		r.Register("client.command.NuLookupUserInfo", acct.NuLookupUserInfo)
		r.Register("client.command.GetPingSites", fsys.GetPingSites)
		r.Register("client.command.UpdateStats", rank.UpdateStats)
		r.Register("client.command.GetTelemetryToken", acct.GetTelemetryToken)
		r.Register("client.command.Start", func(ev network.EventClientCommand) {
			pnow.Start(ev)
			pnow.Status(ev)
		})
		r.Register("client.command.MemCheck", func(ev network.EventClientCommand) {
			// By now: Nothing interesting here, we can skip it.
			// TODO: Use MemCheck response for telemetry.
		})
	}

	return &Fesl{socket, r}
}

// ListenAndServe starts listening socket in goroutine
func (fm *Fesl) ListenAndServe(ctx context.Context) {
	go fm.Run(ctx)
}

// Run starts listening the socket for events and handles them upon receiving
// a message
func (fm *Fesl) Run(ctx context.Context) {
	for {
		select {
		case event := <-fm.socket.EventChan:
			fm.Handle(event)
		case <-ctx.Done():
			return
		}
	}
}

// Handle takes care of handling a single event
func (fm *Fesl) Handle(event network.SocketEvent) {
	ev, ok := event.Data.(network.EventClientCommand)
	if !ok {
		logrus.Error("Logic error: Cannot cast event to network.EventClientCommand")
		return
	}

	if !ev.Client.IsActive {
		logrus.WithField("command", ev.Command).Warn("Inactive client")
		return
	}

	fn, ok := fm.cmds.Find(event.Name)
	if !ok {
		logrus.
			WithFields(logrus.Fields{
				"event":   event.Name,
				"payload": ev.Command.Message,
				"query":   ev.Command.Query,
			}).
			Warn("fesl.UnhandledRequest")
		return
	}

	fn(ev)
}
