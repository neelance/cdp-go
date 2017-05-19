// (experimental)
package storage

import (
	"github.com/neelance/cdp-go/rpc"
)

// (experimental)
type Domain struct {
	Client *rpc.Client
}

// Enum of possible storage types.

type StorageType string

// Clears storage for origin.
type ClearDataForOriginRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) ClearDataForOrigin() *ClearDataForOriginRequest {
	return &ClearDataForOriginRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Security origin.
func (r *ClearDataForOriginRequest) Origin(v string) *ClearDataForOriginRequest {
	r.opts["origin"] = v
	return r
}

// Comma separated origin names.
func (r *ClearDataForOriginRequest) StorageTypes(v string) *ClearDataForOriginRequest {
	r.opts["storageTypes"] = v
	return r
}

// Clears storage for origin.
func (r *ClearDataForOriginRequest) Do() error {
	return r.client.Call("Storage.clearDataForOrigin", r.opts, nil)
}

func init() {
}
