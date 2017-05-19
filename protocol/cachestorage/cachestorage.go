// (experimental)
package cachestorage

import (
	"github.com/neelance/cdp-go/rpc"
)

// (experimental)
type Domain struct {
	Client *rpc.Client
}

// Unique identifier of the Cache object.

type CacheId string

// Data entry.

type DataEntry struct {
	// Request url spec.
	Request string `json:"request"`

	// Response stataus text.
	Response string `json:"response"`
}

// Cache identifier.

type Cache struct {
	// An opaque unique id of the cache.
	CacheId CacheId `json:"cacheId"`

	// Security origin of the cache.
	SecurityOrigin string `json:"securityOrigin"`

	// The name of the cache.
	CacheName string `json:"cacheName"`
}

type RequestCacheNamesOpts struct {
	// Security origin.
	SecurityOrigin string `json:"securityOrigin"`
}

type RequestCacheNamesResult struct {
	// Caches for the security origin.
	Caches []*Cache `json:"caches"`
}

// Requests cache names.
func (d *Domain) RequestCacheNames(opts *RequestCacheNamesOpts) (*RequestCacheNamesResult, error) {
	var result RequestCacheNamesResult
	err := d.Client.Call("CacheStorage.requestCacheNames", opts, &result)
	return &result, err
}

type RequestEntriesOpts struct {
	// ID of cache to get entries from.
	CacheId CacheId `json:"cacheId"`

	// Number of records to skip.
	SkipCount int `json:"skipCount"`

	// Number of records to fetch.
	PageSize int `json:"pageSize"`
}

type RequestEntriesResult struct {
	// Array of object store data entries.
	CacheDataEntries []*DataEntry `json:"cacheDataEntries"`

	// If true, there are more entries to fetch in the given range.
	HasMore bool `json:"hasMore"`
}

// Requests data from cache.
func (d *Domain) RequestEntries(opts *RequestEntriesOpts) (*RequestEntriesResult, error) {
	var result RequestEntriesResult
	err := d.Client.Call("CacheStorage.requestEntries", opts, &result)
	return &result, err
}

type DeleteCacheOpts struct {
	// Id of cache for deletion.
	CacheId CacheId `json:"cacheId"`
}

// Deletes a cache.
func (d *Domain) DeleteCache(opts *DeleteCacheOpts) error {
	return d.Client.Call("CacheStorage.deleteCache", opts, nil)
}

type DeleteEntryOpts struct {
	// Id of cache where the entry will be deleted.
	CacheId CacheId `json:"cacheId"`

	// URL spec of the request.
	Request string `json:"request"`
}

// Deletes a cache entry.
func (d *Domain) DeleteEntry(opts *DeleteEntryOpts) error {
	return d.Client.Call("CacheStorage.deleteEntry", opts, nil)
}
