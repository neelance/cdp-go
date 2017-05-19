// (experimental)
package database

import (
	"github.com/neelance/cdp-go/rpc"
)

// (experimental)
type Domain struct {
	Client *rpc.Client
}

// Unique identifier of Database object. (experimental)

type DatabaseId string

// Database object. (experimental)

type Database struct {
	// Database ID.
	Id DatabaseId `json:"id"`

	// Database domain.
	Domain string `json:"domain"`

	// Database name.
	Name string `json:"name"`

	// Database version.
	Version string `json:"version"`
}

// Database error.

type Error struct {
	// Error message.
	Message string `json:"message"`

	// Error code.
	Code int `json:"code"`
}

// Enables database tracking, database events will now be delivered to the client.
type EnableRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) Enable() *EnableRequest {
	return &EnableRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Enables database tracking, database events will now be delivered to the client.
func (r *EnableRequest) Do() error {
	return r.client.Call("Database.enable", r.opts, nil)
}

// Disables database tracking, prevents database events from being sent to the client.
type DisableRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) Disable() *DisableRequest {
	return &DisableRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Disables database tracking, prevents database events from being sent to the client.
func (r *DisableRequest) Do() error {
	return r.client.Call("Database.disable", r.opts, nil)
}

type GetDatabaseTableNamesRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) GetDatabaseTableNames() *GetDatabaseTableNamesRequest {
	return &GetDatabaseTableNamesRequest{opts: make(map[string]interface{}), client: d.Client}
}

func (r *GetDatabaseTableNamesRequest) DatabaseId(v DatabaseId) *GetDatabaseTableNamesRequest {
	r.opts["databaseId"] = v
	return r
}

type GetDatabaseTableNamesResult struct {
	TableNames []string `json:"tableNames"`
}

func (r *GetDatabaseTableNamesRequest) Do() (*GetDatabaseTableNamesResult, error) {
	var result GetDatabaseTableNamesResult
	err := r.client.Call("Database.getDatabaseTableNames", r.opts, &result)
	return &result, err
}

type ExecuteSQLRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) ExecuteSQL() *ExecuteSQLRequest {
	return &ExecuteSQLRequest{opts: make(map[string]interface{}), client: d.Client}
}

func (r *ExecuteSQLRequest) DatabaseId(v DatabaseId) *ExecuteSQLRequest {
	r.opts["databaseId"] = v
	return r
}

func (r *ExecuteSQLRequest) Query(v string) *ExecuteSQLRequest {
	r.opts["query"] = v
	return r
}

type ExecuteSQLResult struct {
	// (optional)
	ColumnNames []string `json:"columnNames"`

	// (optional)
	Values []interface{} `json:"values"`

	// (optional)
	SqlError *Error `json:"sqlError"`
}

func (r *ExecuteSQLRequest) Do() (*ExecuteSQLResult, error) {
	var result ExecuteSQLResult
	err := r.client.Call("Database.executeSQL", r.opts, &result)
	return &result, err
}

func init() {
	rpc.EventTypes["Database.addDatabase"] = func() interface{} { return new(AddDatabaseEvent) }
}

type AddDatabaseEvent struct {
	Database *Database `json:"database"`
}
