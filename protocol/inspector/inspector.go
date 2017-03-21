// (experimental)
package inspector

import (
	"encoding/json"
	"log"

	"github.com/neelance/cdp-go/rpc"
)

// (experimental)
type Domain struct {
	Client *rpc.Client
}

// Enables inspector domain notifications.
func (d *Domain) Enable() error {
	return d.Client.Call("Inspector.enable", nil, nil)
}

// Disables inspector domain notifications.
func (d *Domain) Disable() error {
	return d.Client.Call("Inspector.disable", nil, nil)
}

type DetachedEvent struct {
	// The reason why connection has been terminated.
	Reason string `json:"reason"`
}

// Fired when remote debugging connection is about to be terminated. Contains detach reason.
func (d *Domain) OnDetached(listener func(*DetachedEvent)) {
	d.Client.AddListener("Inspector.detached", func(params json.RawMessage) {
		var event DetachedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type TargetCrashedEvent struct {
}

// Fired when debugging target has crashed
func (d *Domain) OnTargetCrashed(listener func(*TargetCrashedEvent)) {
	d.Client.AddListener("Inspector.targetCrashed", func(params json.RawMessage) {
		var event TargetCrashedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}
