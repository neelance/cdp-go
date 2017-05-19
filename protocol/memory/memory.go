// (experimental)
package memory

import (
	"github.com/neelance/cdp-go/rpc"
)

// (experimental)
type Domain struct {
	Client *rpc.Client
}

// Memory pressure level.

type PressureLevel string

type GetDOMCountersResult struct {
	Documents int `json:"documents"`

	Nodes int `json:"nodes"`

	JsEventListeners int `json:"jsEventListeners"`
}

func (d *Domain) GetDOMCounters() (*GetDOMCountersResult, error) {
	var result GetDOMCountersResult
	err := d.Client.Call("Memory.getDOMCounters", nil, &result)
	return &result, err
}

type SetPressureNotificationsSuppressedOpts struct {
	// If true, memory pressure notifications will be suppressed.
	Suppressed bool `json:"suppressed"`
}

// Enable/disable suppressing memory pressure notifications in all processes.
func (d *Domain) SetPressureNotificationsSuppressed(opts *SetPressureNotificationsSuppressedOpts) error {
	return d.Client.Call("Memory.setPressureNotificationsSuppressed", opts, nil)
}

type SimulatePressureNotificationOpts struct {
	// Memory pressure level of the notification.
	Level PressureLevel `json:"level"`
}

// Simulate a memory pressure notification in all processes.
func (d *Domain) SimulatePressureNotification(opts *SimulatePressureNotificationOpts) error {
	return d.Client.Call("Memory.simulatePressureNotification", opts, nil)
}
