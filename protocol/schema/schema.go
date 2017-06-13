// Provides information about the protocol schema.
package schema

import (
	"github.com/neelance/cdp-go/rpc"
)

// Provides information about the protocol schema.
type Domain struct {
	Client *rpc.Client
}

// Description of the protocol domain.

type Domain struct {
	// Domain name.
	Name string `json:"name"`

	// Domain version.
	Version string `json:"version"`
}

// Returns supported domains.
type GetDomainsRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) GetDomains() *GetDomainsRequest {
	return &GetDomainsRequest{opts: make(map[string]interface{}), client: d.Client}
}

type GetDomainsResult struct {
	// List of supported domains.
	Domains []*Domain `json:"domains"`
}

func (r *GetDomainsRequest) Do() (*GetDomainsResult, error) {
	var result GetDomainsResult
	err := r.client.Call("Schema.getDomains", r.opts, &result)
	return &result, err
}

func init() {
}
