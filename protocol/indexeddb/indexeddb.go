// (experimental)
package indexeddb

import (
	"github.com/neelance/cdp-go/rpc"
)

// (experimental)
type Domain struct {
	Client *rpc.Client
}

// Database with an array of object stores.
type DatabaseWithObjectStores interface{}

// Object store.
type ObjectStore interface{}

// Object store index.
type ObjectStoreIndex interface{}

// Key.
type Key interface{}

// Key range.
type KeyRange interface{}

// Data entry.
type DataEntry interface{}

// Key path.
type KeyPath interface{}

// Enables events from backend.
func (d *Domain) Enable() error {
	return d.Client.Call("IndexedDB.enable", nil, nil)
}

// Disables events from backend.
func (d *Domain) Disable() error {
	return d.Client.Call("IndexedDB.disable", nil, nil)
}

type RequestDatabaseNamesOpts struct {
	// Security origin.
	SecurityOrigin string `json:"securityOrigin"`
}

type RequestDatabaseNamesResult struct {
	// Database names for origin.
	DatabaseNames []string `json:"databaseNames"`
}

// Requests database names for given security origin.
func (d *Domain) RequestDatabaseNames(opts *RequestDatabaseNamesOpts) (*RequestDatabaseNamesResult, error) {
	var result RequestDatabaseNamesResult
	err := d.Client.Call("IndexedDB.requestDatabaseNames", opts, &result)
	return &result, err
}

type RequestDatabaseOpts struct {
	// Security origin.
	SecurityOrigin string `json:"securityOrigin"`

	// Database name.
	DatabaseName string `json:"databaseName"`
}

type RequestDatabaseResult struct {
	// Database with an array of object stores.
	DatabaseWithObjectStores DatabaseWithObjectStores `json:"databaseWithObjectStores"`
}

// Requests database with given name in given frame.
func (d *Domain) RequestDatabase(opts *RequestDatabaseOpts) (*RequestDatabaseResult, error) {
	var result RequestDatabaseResult
	err := d.Client.Call("IndexedDB.requestDatabase", opts, &result)
	return &result, err
}

type RequestDataOpts struct {
	// Security origin.
	SecurityOrigin string `json:"securityOrigin"`

	// Database name.
	DatabaseName string `json:"databaseName"`

	// Object store name.
	ObjectStoreName string `json:"objectStoreName"`

	// Index name, empty string for object store data requests.
	IndexName string `json:"indexName"`

	// Number of records to skip.
	SkipCount int `json:"skipCount"`

	// Number of records to fetch.
	PageSize int `json:"pageSize"`

	// Key range. (optional)
	KeyRange KeyRange `json:"keyRange,omitempty"`
}

type RequestDataResult struct {
	// Array of object store data entries.
	ObjectStoreDataEntries []DataEntry `json:"objectStoreDataEntries"`

	// If true, there are more entries to fetch in the given range.
	HasMore bool `json:"hasMore"`
}

// Requests data from object store or index.
func (d *Domain) RequestData(opts *RequestDataOpts) (*RequestDataResult, error) {
	var result RequestDataResult
	err := d.Client.Call("IndexedDB.requestData", opts, &result)
	return &result, err
}

type ClearObjectStoreOpts struct {
	// Security origin.
	SecurityOrigin string `json:"securityOrigin"`

	// Database name.
	DatabaseName string `json:"databaseName"`

	// Object store name.
	ObjectStoreName string `json:"objectStoreName"`
}

// Clears all entries from an object store.
func (d *Domain) ClearObjectStore(opts *ClearObjectStoreOpts) error {
	return d.Client.Call("IndexedDB.clearObjectStore", opts, nil)
}

type DeleteDatabaseOpts struct {
	// Security origin.
	SecurityOrigin string `json:"securityOrigin"`

	// Database name.
	DatabaseName string `json:"databaseName"`
}

// Deletes a database.
func (d *Domain) DeleteDatabase(opts *DeleteDatabaseOpts) error {
	return d.Client.Call("IndexedDB.deleteDatabase", opts, nil)
}
