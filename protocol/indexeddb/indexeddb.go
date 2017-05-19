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

type DatabaseWithObjectStores struct {
	// Database name.
	Name string `json:"name"`

	// Database version.
	Version int `json:"version"`

	// Object stores in this database.
	ObjectStores []*ObjectStore `json:"objectStores"`
}

// Object store.

type ObjectStore struct {
	// Object store name.
	Name string `json:"name"`

	// Object store key path.
	KeyPath *KeyPath `json:"keyPath"`

	// If true, object store has auto increment flag set.
	AutoIncrement bool `json:"autoIncrement"`

	// Indexes in this object store.
	Indexes []*ObjectStoreIndex `json:"indexes"`
}

// Object store index.

type ObjectStoreIndex struct {
	// Index name.
	Name string `json:"name"`

	// Index key path.
	KeyPath *KeyPath `json:"keyPath"`

	// If true, index is unique.
	Unique bool `json:"unique"`

	// If true, index allows multiple entries for a key.
	MultiEntry bool `json:"multiEntry"`
}

// Key.

type Key struct {
	// Key type.
	Type string `json:"type"`

	// Number value. (optional)
	Number float64 `json:"number,omitempty"`

	// String value. (optional)
	String string `json:"string,omitempty"`

	// Date value. (optional)
	Date float64 `json:"date,omitempty"`

	// Array value. (optional)
	Array []*Key `json:"array,omitempty"`
}

// Key range.

type KeyRange struct {
	// Lower bound. (optional)
	Lower *Key `json:"lower,omitempty"`

	// Upper bound. (optional)
	Upper *Key `json:"upper,omitempty"`

	// If true lower bound is open.
	LowerOpen bool `json:"lowerOpen"`

	// If true upper bound is open.
	UpperOpen bool `json:"upperOpen"`
}

// Data entry.

type DataEntry struct {
	// Key object.
	Key interface{} `json:"key"`

	// Primary key object.
	PrimaryKey interface{} `json:"primaryKey"`

	// Value object.
	Value interface{} `json:"value"`
}

// Key path.

type KeyPath struct {
	// Key path type.
	Type string `json:"type"`

	// String value. (optional)
	String string `json:"string,omitempty"`

	// Array value. (optional)
	Array []string `json:"array,omitempty"`
}

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
	DatabaseWithObjectStores *DatabaseWithObjectStores `json:"databaseWithObjectStores"`
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
	KeyRange *KeyRange `json:"keyRange,omitempty"`
}

type RequestDataResult struct {
	// Array of object store data entries.
	ObjectStoreDataEntries []*DataEntry `json:"objectStoreDataEntries"`

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
