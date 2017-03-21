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
type StorageType interface{}

type ClearDataForOriginOpts struct {
	// Security origin.
	Origin string `json:"origin"`

	// Comma separated origin names.
	StorageTypes string `json:"storageTypes"`
}

// Clears storage for origin.
func (d *Domain) ClearDataForOrigin(opts *ClearDataForOriginOpts) error {
	return d.Client.Call("Storage.clearDataForOrigin", opts, nil)
}
