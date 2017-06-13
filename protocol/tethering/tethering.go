// The Tethering domain defines methods and events for browser port binding. (experimental)
package tethering

import (
	"github.com/neelance/cdp-go/rpc"
)

// The Tethering domain defines methods and events for browser port binding. (experimental)
type Client struct {
	*rpc.Client
}

// Request browser port binding.
type BindRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Client) Bind() *BindRequest {
	return &BindRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Port number to bind.
func (r *BindRequest) Port(v int) *BindRequest {
	r.opts["port"] = v
	return r
}

// Request browser port binding.
func (r *BindRequest) Do() error {
	return r.client.Call("Tethering.bind", r.opts, nil)
}

// Request browser port unbinding.
type UnbindRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Client) Unbind() *UnbindRequest {
	return &UnbindRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Port number to unbind.
func (r *UnbindRequest) Port(v int) *UnbindRequest {
	r.opts["port"] = v
	return r
}

// Request browser port unbinding.
func (r *UnbindRequest) Do() error {
	return r.client.Call("Tethering.unbind", r.opts, nil)
}

func init() {
	rpc.EventTypes["Tethering.accepted"] = func() interface{} { return new(AcceptedEvent) }
}

// Informs that port was successfully bound and got a specified connection id.
type AcceptedEvent struct {
	// Port number that was successfully bound.
	Port int `json:"port"`

	// Connection id to be used.
	ConnectionId string `json:"connectionId"`
}
