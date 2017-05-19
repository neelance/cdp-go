// The Browser domain defines methods and events for browser managing. (experimental)
package browser

import (
	"github.com/neelance/cdp-go/rpc"
)

// The Browser domain defines methods and events for browser managing. (experimental)
type Domain struct {
	Client *rpc.Client
}

type WindowID int

// The state of the browser window.

type WindowState string

// Browser window bounds information

type Bounds struct {
	// The offset from the left edge of the screen to the window in pixels. (optional)
	Left int `json:"left,omitempty"`

	// The offset from the top edge of the screen to the window in pixels. (optional)
	Top int `json:"top,omitempty"`

	// The window width in pixels. (optional)
	Width int `json:"width,omitempty"`

	// The window height in pixels. (optional)
	Height int `json:"height,omitempty"`

	// The window state. Default to normal. (optional)
	WindowState WindowState `json:"windowState,omitempty"`
}

// Get the browser window that contains the devtools target.
type GetWindowForTargetRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) GetWindowForTarget() *GetWindowForTargetRequest {
	return &GetWindowForTargetRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Devtools agent host id.
func (r *GetWindowForTargetRequest) TargetId(v interface{}) *GetWindowForTargetRequest {
	r.opts["targetId"] = v
	return r
}

type GetWindowForTargetResult struct {
	// Browser window id.
	WindowId WindowID `json:"windowId"`

	// Bounds information of the window. When window state is 'minimized', the restored window position and size are returned.
	Bounds *Bounds `json:"bounds"`
}

func (r *GetWindowForTargetRequest) Do() (*GetWindowForTargetResult, error) {
	var result GetWindowForTargetResult
	err := r.client.Call("Browser.getWindowForTarget", r.opts, &result)
	return &result, err
}

// Set position and/or size of the browser window.
type SetWindowBoundsRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) SetWindowBounds() *SetWindowBoundsRequest {
	return &SetWindowBoundsRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Browser window id.
func (r *SetWindowBoundsRequest) WindowId(v WindowID) *SetWindowBoundsRequest {
	r.opts["windowId"] = v
	return r
}

// New window bounds. The 'minimized', 'maximized' and 'fullscreen' states cannot be combined with 'left', 'top', 'width' or 'height'. Leaves unspecified fields unchanged.
func (r *SetWindowBoundsRequest) Bounds(v *Bounds) *SetWindowBoundsRequest {
	r.opts["bounds"] = v
	return r
}

// Set position and/or size of the browser window.
func (r *SetWindowBoundsRequest) Do() error {
	return r.client.Call("Browser.setWindowBounds", r.opts, nil)
}

// Get position and size of the browser window.
type GetWindowBoundsRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) GetWindowBounds() *GetWindowBoundsRequest {
	return &GetWindowBoundsRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Browser window id.
func (r *GetWindowBoundsRequest) WindowId(v WindowID) *GetWindowBoundsRequest {
	r.opts["windowId"] = v
	return r
}

type GetWindowBoundsResult struct {
	// Bounds information of the window. When window state is 'minimized', the restored window position and size are returned.
	Bounds *Bounds `json:"bounds"`
}

func (r *GetWindowBoundsRequest) Do() (*GetWindowBoundsResult, error) {
	var result GetWindowBoundsResult
	err := r.client.Call("Browser.getWindowBounds", r.opts, &result)
	return &result, err
}

func init() {
}
