// Query and modify DOM storage. (experimental)
package domstorage

import (
	"encoding/json"
	"log"

	"github.com/neelance/cdp-go/rpc"
)

// Query and modify DOM storage. (experimental)
type Domain struct {
	Client *rpc.Client
}

// DOM Storage identifier. (experimental)

type StorageId struct {
	// Security origin for the storage.
	SecurityOrigin string `json:"securityOrigin"`

	// Whether the storage is local storage (not session storage).
	IsLocalStorage bool `json:"isLocalStorage"`
}

// DOM Storage item. (experimental)

type Item []string

// Enables storage tracking, storage events will now be delivered to the client.
func (d *Domain) Enable() error {
	return d.Client.Call("DOMStorage.enable", nil, nil)
}

// Disables storage tracking, prevents storage events from being sent to the client.
func (d *Domain) Disable() error {
	return d.Client.Call("DOMStorage.disable", nil, nil)
}

type ClearOpts struct {
	StorageId *StorageId `json:"storageId"`
}

func (d *Domain) Clear(opts *ClearOpts) error {
	return d.Client.Call("DOMStorage.clear", opts, nil)
}

type GetDOMStorageItemsOpts struct {
	StorageId *StorageId `json:"storageId"`
}

type GetDOMStorageItemsResult struct {
	Entries []*Item `json:"entries"`
}

func (d *Domain) GetDOMStorageItems(opts *GetDOMStorageItemsOpts) (*GetDOMStorageItemsResult, error) {
	var result GetDOMStorageItemsResult
	err := d.Client.Call("DOMStorage.getDOMStorageItems", opts, &result)
	return &result, err
}

type SetDOMStorageItemOpts struct {
	StorageId *StorageId `json:"storageId"`

	Key string `json:"key"`

	Value string `json:"value"`
}

func (d *Domain) SetDOMStorageItem(opts *SetDOMStorageItemOpts) error {
	return d.Client.Call("DOMStorage.setDOMStorageItem", opts, nil)
}

type RemoveDOMStorageItemOpts struct {
	StorageId *StorageId `json:"storageId"`

	Key string `json:"key"`
}

func (d *Domain) RemoveDOMStorageItem(opts *RemoveDOMStorageItemOpts) error {
	return d.Client.Call("DOMStorage.removeDOMStorageItem", opts, nil)
}

type DomStorageItemsClearedEvent struct {
	StorageId *StorageId `json:"storageId"`
}

func (d *Domain) OnDomStorageItemsCleared(listener func(*DomStorageItemsClearedEvent)) {
	d.Client.AddListener("DOMStorage.domStorageItemsCleared", func(params json.RawMessage) {
		var event DomStorageItemsClearedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type DomStorageItemRemovedEvent struct {
	StorageId *StorageId `json:"storageId"`

	Key string `json:"key"`
}

func (d *Domain) OnDomStorageItemRemoved(listener func(*DomStorageItemRemovedEvent)) {
	d.Client.AddListener("DOMStorage.domStorageItemRemoved", func(params json.RawMessage) {
		var event DomStorageItemRemovedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type DomStorageItemAddedEvent struct {
	StorageId *StorageId `json:"storageId"`

	Key string `json:"key"`

	NewValue string `json:"newValue"`
}

func (d *Domain) OnDomStorageItemAdded(listener func(*DomStorageItemAddedEvent)) {
	d.Client.AddListener("DOMStorage.domStorageItemAdded", func(params json.RawMessage) {
		var event DomStorageItemAddedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type DomStorageItemUpdatedEvent struct {
	StorageId *StorageId `json:"storageId"`

	Key string `json:"key"`

	OldValue string `json:"oldValue"`

	NewValue string `json:"newValue"`
}

func (d *Domain) OnDomStorageItemUpdated(listener func(*DomStorageItemUpdatedEvent)) {
	d.Client.AddListener("DOMStorage.domStorageItemUpdated", func(params json.RawMessage) {
		var event DomStorageItemUpdatedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}
