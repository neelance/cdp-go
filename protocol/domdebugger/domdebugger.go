// DOM debugging allows setting breakpoints on particular DOM operations and events. JavaScript execution will stop on these operations as if there was a regular breakpoint set.
package domdebugger

import (
	"github.com/neelance/cdp-go/rpc"
)

// DOM debugging allows setting breakpoints on particular DOM operations and events. JavaScript execution will stop on these operations as if there was a regular breakpoint set.
type Domain struct {
	Client *rpc.Client
}

// DOM breakpoint type.

type DOMBreakpointType string

// Object event listener. (experimental)

type EventListener struct {
	// <code>EventListener</code>'s type.
	Type string `json:"type"`

	// <code>EventListener</code>'s useCapture.
	UseCapture bool `json:"useCapture"`

	// <code>EventListener</code>'s passive flag.
	Passive bool `json:"passive"`

	// <code>EventListener</code>'s once flag.
	Once bool `json:"once"`

	// Script id of the handler code.
	ScriptId interface{} `json:"scriptId"`

	// Line number in the script (0-based).
	LineNumber int `json:"lineNumber"`

	// Column number in the script (0-based).
	ColumnNumber int `json:"columnNumber"`

	// Event handler function value. (optional)
	Handler interface{} `json:"handler,omitempty"`

	// Event original handler function value. (optional)
	OriginalHandler interface{} `json:"originalHandler,omitempty"`

	// Node the listener is added to (if any). (optional)
	BackendNodeId interface{} `json:"backendNodeId,omitempty"`
}

type SetDOMBreakpointOpts struct {
	// Identifier of the node to set breakpoint on.
	NodeId interface{} `json:"nodeId"`

	// Type of the operation to stop upon.
	Type DOMBreakpointType `json:"type"`
}

// Sets breakpoint on particular operation with DOM.
func (d *Domain) SetDOMBreakpoint(opts *SetDOMBreakpointOpts) error {
	return d.Client.Call("DOMDebugger.setDOMBreakpoint", opts, nil)
}

type RemoveDOMBreakpointOpts struct {
	// Identifier of the node to remove breakpoint from.
	NodeId interface{} `json:"nodeId"`

	// Type of the breakpoint to remove.
	Type DOMBreakpointType `json:"type"`
}

// Removes DOM breakpoint that was set using <code>setDOMBreakpoint</code>.
func (d *Domain) RemoveDOMBreakpoint(opts *RemoveDOMBreakpointOpts) error {
	return d.Client.Call("DOMDebugger.removeDOMBreakpoint", opts, nil)
}

type SetEventListenerBreakpointOpts struct {
	// DOM Event name to stop on (any DOM event will do).
	EventName string `json:"eventName"`

	// EventTarget interface name to stop on. If equal to <code>"*"</code> or not provided, will stop on any EventTarget. (optional, experimental)
	TargetName string `json:"targetName,omitempty"`
}

// Sets breakpoint on particular DOM event.
func (d *Domain) SetEventListenerBreakpoint(opts *SetEventListenerBreakpointOpts) error {
	return d.Client.Call("DOMDebugger.setEventListenerBreakpoint", opts, nil)
}

type RemoveEventListenerBreakpointOpts struct {
	// Event name.
	EventName string `json:"eventName"`

	// EventTarget interface name. (optional, experimental)
	TargetName string `json:"targetName,omitempty"`
}

// Removes breakpoint on particular DOM event.
func (d *Domain) RemoveEventListenerBreakpoint(opts *RemoveEventListenerBreakpointOpts) error {
	return d.Client.Call("DOMDebugger.removeEventListenerBreakpoint", opts, nil)
}

type SetInstrumentationBreakpointOpts struct {
	// Instrumentation name to stop on.
	EventName string `json:"eventName"`
}

// Sets breakpoint on particular native event. (experimental)
func (d *Domain) SetInstrumentationBreakpoint(opts *SetInstrumentationBreakpointOpts) error {
	return d.Client.Call("DOMDebugger.setInstrumentationBreakpoint", opts, nil)
}

type RemoveInstrumentationBreakpointOpts struct {
	// Instrumentation name to stop on.
	EventName string `json:"eventName"`
}

// Removes breakpoint on particular native event. (experimental)
func (d *Domain) RemoveInstrumentationBreakpoint(opts *RemoveInstrumentationBreakpointOpts) error {
	return d.Client.Call("DOMDebugger.removeInstrumentationBreakpoint", opts, nil)
}

type SetXHRBreakpointOpts struct {
	// Resource URL substring. All XHRs having this substring in the URL will get stopped upon.
	URL string `json:"url"`
}

// Sets breakpoint on XMLHttpRequest.
func (d *Domain) SetXHRBreakpoint(opts *SetXHRBreakpointOpts) error {
	return d.Client.Call("DOMDebugger.setXHRBreakpoint", opts, nil)
}

type RemoveXHRBreakpointOpts struct {
	// Resource URL substring.
	URL string `json:"url"`
}

// Removes breakpoint from XMLHttpRequest.
func (d *Domain) RemoveXHRBreakpoint(opts *RemoveXHRBreakpointOpts) error {
	return d.Client.Call("DOMDebugger.removeXHRBreakpoint", opts, nil)
}

type GetEventListenersOpts struct {
	// Identifier of the object to return listeners for.
	ObjectId interface{} `json:"objectId"`

	// The maximum depth at which Node children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0. (optional, experimental)
	Depth int `json:"depth,omitempty"`

	// Whether or not iframes and shadow roots should be traversed when returning the subtree (default is false). Reports listeners for all contexts if pierce is enabled. (optional, experimental)
	Pierce bool `json:"pierce,omitempty"`
}

type GetEventListenersResult struct {
	// Array of relevant listeners.
	Listeners []*EventListener `json:"listeners"`
}

// Returns event listeners of the given object. (experimental)
func (d *Domain) GetEventListeners(opts *GetEventListenersOpts) (*GetEventListenersResult, error) {
	var result GetEventListenersResult
	err := d.Client.Call("DOMDebugger.getEventListeners", opts, &result)
	return &result, err
}
