// Provides access to log entries. (experimental)
package log

import (
	"encoding/json"
	"log"

	"github.com/neelance/cdp-go/rpc"
)

// Provides access to log entries. (experimental)
type Domain struct {
	Client *rpc.Client
}

// Log entry.
type LogEntry interface{}

// Violation configuration setting.
type ViolationSetting interface{}

// Enables log domain, sends the entries collected so far to the client by means of the <code>entryAdded</code> notification.
func (d *Domain) Enable() error {
	return d.Client.Call("Log.enable", nil, nil)
}

// Disables log domain, prevents further log entries from being reported to the client.
func (d *Domain) Disable() error {
	return d.Client.Call("Log.disable", nil, nil)
}

// Clears the log.
func (d *Domain) Clear() error {
	return d.Client.Call("Log.clear", nil, nil)
}

type StartViolationsReportOpts struct {
	// Configuration for violations.
	Config []ViolationSetting `json:"config"`
}

// start violation reporting.
func (d *Domain) StartViolationsReport(opts *StartViolationsReportOpts) error {
	return d.Client.Call("Log.startViolationsReport", opts, nil)
}

// Stop violation reporting.
func (d *Domain) StopViolationsReport() error {
	return d.Client.Call("Log.stopViolationsReport", nil, nil)
}

type EntryAddedEvent struct {
	// The entry.
	Entry LogEntry `json:"entry"`
}

// Issued when new message was logged.
func (d *Domain) OnEntryAdded(listener func(*EntryAddedEvent)) {
	d.Client.AddListener("Log.entryAdded", func(params json.RawMessage) {
		var event EntryAddedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}