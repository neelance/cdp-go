// (experimental)
package database

import (
	"encoding/json"
	"log"

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
func (d *Domain) Enable() error {
	return d.Client.Call("Database.enable", nil, nil)
}

// Disables database tracking, prevents database events from being sent to the client.
func (d *Domain) Disable() error {
	return d.Client.Call("Database.disable", nil, nil)
}

type GetDatabaseTableNamesOpts struct {
	DatabaseId DatabaseId `json:"databaseId"`
}

type GetDatabaseTableNamesResult struct {
	TableNames []string `json:"tableNames"`
}

func (d *Domain) GetDatabaseTableNames(opts *GetDatabaseTableNamesOpts) (*GetDatabaseTableNamesResult, error) {
	var result GetDatabaseTableNamesResult
	err := d.Client.Call("Database.getDatabaseTableNames", opts, &result)
	return &result, err
}

type ExecuteSQLOpts struct {
	DatabaseId DatabaseId `json:"databaseId"`

	Query string `json:"query"`
}

type ExecuteSQLResult struct {
	// (optional)
	ColumnNames []string `json:"columnNames"`

	// (optional)
	Values []interface{} `json:"values"`

	// (optional)
	SqlError *Error `json:"sqlError"`
}

func (d *Domain) ExecuteSQL(opts *ExecuteSQLOpts) (*ExecuteSQLResult, error) {
	var result ExecuteSQLResult
	err := d.Client.Call("Database.executeSQL", opts, &result)
	return &result, err
}

type AddDatabaseEvent struct {
	Database *Database `json:"database"`
}

func (d *Domain) OnAddDatabase(listener func(*AddDatabaseEvent)) {
	d.Client.AddListener("Database.addDatabase", func(params json.RawMessage) {
		var event AddDatabaseEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}
