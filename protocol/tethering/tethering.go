// The Tethering domain defines methods and events for browser port binding. (experimental)
package tethering

import (
	"encoding/json"
	"log"

	"github.com/neelance/cdp-go/rpc"
)

// The Tethering domain defines methods and events for browser port binding. (experimental)
type Domain struct {
	Client *rpc.Client
}

type BindOpts struct {
	// Port number to bind.
	Port int `json:"port"`
}

// Request browser port binding.
func (d *Domain) Bind(opts *BindOpts) error {
	return d.Client.Call("Tethering.bind", opts, nil)
}

type UnbindOpts struct {
	// Port number to unbind.
	Port int `json:"port"`
}

// Request browser port unbinding.
func (d *Domain) Unbind(opts *UnbindOpts) error {
	return d.Client.Call("Tethering.unbind", opts, nil)
}

type AcceptedEvent struct {
	// Port number that was successfully bound.
	Port int `json:"port"`

	// Connection id to be used.
	ConnectionId string `json:"connectionId"`
}

// Informs that port was successfully bound and got a specified connection id.
func (d *Domain) OnAccepted(listener func(*AcceptedEvent)) {
	d.Client.AddListener("Tethering.accepted", func(params json.RawMessage) {
		var event AcceptedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}
