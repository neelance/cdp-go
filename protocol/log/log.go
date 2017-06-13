// Provides access to log entries. (experimental)
package log

import (
	"github.com/neelance/cdp-go/rpc"
)

// Provides access to log entries. (experimental)
type Client struct {
	*rpc.Client
}

// Log entry.

type LogEntry struct {
	// Log entry source.
	Source string `json:"source"`

	// Log entry severity.
	Level string `json:"level"`

	// Logged text.
	Text string `json:"text"`

	// Timestamp when this entry was added.
	Timestamp interface{} `json:"timestamp"`

	// URL of the resource if known. (optional)
	URL string `json:"url,omitempty"`

	// Line number in the resource. (optional)
	LineNumber int `json:"lineNumber,omitempty"`

	// JavaScript stack trace. (optional)
	StackTrace interface{} `json:"stackTrace,omitempty"`

	// Identifier of the network request associated with this entry. (optional)
	NetworkRequestId interface{} `json:"networkRequestId,omitempty"`

	// Identifier of the worker associated with this entry. (optional)
	WorkerId string `json:"workerId,omitempty"`
}

// Violation configuration setting.

type ViolationSetting struct {
	// Violation type.
	Name string `json:"name"`

	// Time threshold to trigger upon.
	Threshold float64 `json:"threshold"`
}

// Enables log domain, sends the entries collected so far to the client by means of the <code>entryAdded</code> notification.
type EnableRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Client) Enable() *EnableRequest {
	return &EnableRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Enables log domain, sends the entries collected so far to the client by means of the <code>entryAdded</code> notification.
func (r *EnableRequest) Do() error {
	return r.client.Call("Log.enable", r.opts, nil)
}

// Disables log domain, prevents further log entries from being reported to the client.
type DisableRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Client) Disable() *DisableRequest {
	return &DisableRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Disables log domain, prevents further log entries from being reported to the client.
func (r *DisableRequest) Do() error {
	return r.client.Call("Log.disable", r.opts, nil)
}

// Clears the log.
type ClearRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Client) Clear() *ClearRequest {
	return &ClearRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Clears the log.
func (r *ClearRequest) Do() error {
	return r.client.Call("Log.clear", r.opts, nil)
}

// start violation reporting.
type StartViolationsReportRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Client) StartViolationsReport() *StartViolationsReportRequest {
	return &StartViolationsReportRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Configuration for violations.
func (r *StartViolationsReportRequest) Config(v []*ViolationSetting) *StartViolationsReportRequest {
	r.opts["config"] = v
	return r
}

// start violation reporting.
func (r *StartViolationsReportRequest) Do() error {
	return r.client.Call("Log.startViolationsReport", r.opts, nil)
}

// Stop violation reporting.
type StopViolationsReportRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Client) StopViolationsReport() *StopViolationsReportRequest {
	return &StopViolationsReportRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Stop violation reporting.
func (r *StopViolationsReportRequest) Do() error {
	return r.client.Call("Log.stopViolationsReport", r.opts, nil)
}

func init() {
	rpc.EventTypes["Log.entryAdded"] = func() interface{} { return new(EntryAddedEvent) }
}

// Issued when new message was logged.
type EntryAddedEvent struct {
	// The entry.
	Entry *LogEntry `json:"entry"`
}
