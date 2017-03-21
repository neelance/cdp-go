// (experimental)
package applicationcache

import (
	"encoding/json"
	"log"

	"github.com/neelance/cdp-go/rpc"
)

// (experimental)
type Domain struct {
	Client *rpc.Client
}

// Detailed application cache resource information.
type ApplicationCacheResource interface{}

// Detailed application cache information.
type ApplicationCache interface{}

// Frame identifier - manifest URL pair.
type FrameWithManifest interface{}

type GetFramesWithManifestsResult struct {
	// Array of frame identifiers with manifest urls for each frame containing a document associated with some application cache.
	FrameIds []FrameWithManifest `json:"frameIds"`
}

// Returns array of frame identifiers with manifest urls for each frame containing a document associated with some application cache.
func (d *Domain) GetFramesWithManifests() (*GetFramesWithManifestsResult, error) {
	var result GetFramesWithManifestsResult
	err := d.Client.Call("ApplicationCache.getFramesWithManifests", nil, &result)
	return &result, err
}

// Enables application cache domain notifications.
func (d *Domain) Enable() error {
	return d.Client.Call("ApplicationCache.enable", nil, nil)
}

type GetManifestForFrameOpts struct {
	// Identifier of the frame containing document whose manifest is retrieved.
	FrameId interface{} `json:"frameId"`
}

type GetManifestForFrameResult struct {
	// Manifest URL for document in the given frame.
	ManifestURL string `json:"manifestURL"`
}

// Returns manifest URL for document in the given frame.
func (d *Domain) GetManifestForFrame(opts *GetManifestForFrameOpts) (*GetManifestForFrameResult, error) {
	var result GetManifestForFrameResult
	err := d.Client.Call("ApplicationCache.getManifestForFrame", opts, &result)
	return &result, err
}

type GetApplicationCacheForFrameOpts struct {
	// Identifier of the frame containing document whose application cache is retrieved.
	FrameId interface{} `json:"frameId"`
}

type GetApplicationCacheForFrameResult struct {
	// Relevant application cache data for the document in given frame.
	ApplicationCache ApplicationCache `json:"applicationCache"`
}

// Returns relevant application cache data for the document in given frame.
func (d *Domain) GetApplicationCacheForFrame(opts *GetApplicationCacheForFrameOpts) (*GetApplicationCacheForFrameResult, error) {
	var result GetApplicationCacheForFrameResult
	err := d.Client.Call("ApplicationCache.getApplicationCacheForFrame", opts, &result)
	return &result, err
}

type ApplicationCacheStatusUpdatedEvent struct {
	// Identifier of the frame containing document whose application cache updated status.
	FrameId interface{} `json:"frameId"`

	// Manifest URL.
	ManifestURL string `json:"manifestURL"`

	// Updated application cache status.
	Status int `json:"status"`
}

func (d *Domain) OnApplicationCacheStatusUpdated(listener func(*ApplicationCacheStatusUpdatedEvent)) {
	d.Client.AddListener("ApplicationCache.applicationCacheStatusUpdated", func(params json.RawMessage) {
		var event ApplicationCacheStatusUpdatedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type NetworkStateUpdatedEvent struct {
	IsNowOnline bool `json:"isNowOnline"`
}

func (d *Domain) OnNetworkStateUpdated(listener func(*NetworkStateUpdatedEvent)) {
	d.Client.AddListener("ApplicationCache.networkStateUpdated", func(params json.RawMessage) {
		var event NetworkStateUpdatedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}
