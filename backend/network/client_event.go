package network

import (
	"github.com/Synaxis/battlefieldHeroesBackend/backend/network/codec"

	"github.com/sirupsen/logrus"
)

// ClientEvent is the generic struct for events
// by this Client
type ClientEvent struct {
	Name string
	Data interface{}
}

type EventClientCommand struct {
	Client *Client
	// If TLS (theater then we ignore payloadID - it is always 0x0)
	Command *codec.Command
}

func (c *Client) FireClientClose(event ClientEvent) SocketEvent {
	return SocketEvent{
		Name: "client.close",
		Data: EventClientCommand{Client: c},
	}
}

func (c *Client) FireClose() ClientEvent {
	return ClientEvent{Name: "close", Data: c}
}

func (c *Client) FireClientCommand(event ClientEvent) SocketEvent {
	return SocketEvent{
		Name: "client." + event.Name,
		Data: EventClientCommand{
			Client:  c,
			Command: event.Data.(*codec.Command),
		},
	}
}

// CmdFunc is a definition of the command handler
type CmdFunc func(EventClientCommand)

// CmdRegistry takes care of routing command handlers
type CmdRegistry map[string]CmdFunc

// Register add command to registry
func (r CmdRegistry) Register(cmdName string, fn CmdFunc) {
	if _, ok := r[cmdName]; ok {
		logrus.Errorf("Attempting to replace existing command: %s", cmdName)
		return
	}
	r[cmdName] = fn
}

// Find attempts to return command with given name
func (r CmdRegistry) Find(cmdName string) (CmdFunc, bool) {
	cmd, ok := r[cmdName]
	return cmd, ok
}
