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

type GetWindowForTargetOpts struct {
	// Devtools agent host id.
	TargetId interface{} `json:"targetId"`
}

type GetWindowForTargetResult struct {
	// Browser window id.
	WindowId WindowID `json:"windowId"`

	// Bounds information of the window. When window state is 'minimized', the restored window position and size are returned.
	Bounds *Bounds `json:"bounds"`
}

// Get the browser window that contains the devtools target.
func (d *Domain) GetWindowForTarget(opts *GetWindowForTargetOpts) (*GetWindowForTargetResult, error) {
	var result GetWindowForTargetResult
	err := d.Client.Call("Browser.getWindowForTarget", opts, &result)
	return &result, err
}

type SetWindowBoundsOpts struct {
	// Browser window id.
	WindowId WindowID `json:"windowId"`

	// New window bounds. The 'minimized', 'maximized' and 'fullscreen' states cannot be combined with 'left', 'top', 'width' or 'height'. Leaves unspecified fields unchanged.
	Bounds *Bounds `json:"bounds"`
}

// Set position and/or size of the browser window.
func (d *Domain) SetWindowBounds(opts *SetWindowBoundsOpts) error {
	return d.Client.Call("Browser.setWindowBounds", opts, nil)
}

type GetWindowBoundsOpts struct {
	// Browser window id.
	WindowId WindowID `json:"windowId"`
}

type GetWindowBoundsResult struct {
	// Bounds information of the window. When window state is 'minimized', the restored window position and size are returned.
	Bounds *Bounds `json:"bounds"`
}

// Get position and size of the browser window.
func (d *Domain) GetWindowBounds(opts *GetWindowBoundsOpts) (*GetWindowBoundsResult, error) {
	var result GetWindowBoundsResult
	err := d.Client.Call("Browser.getWindowBounds", opts, &result)
	return &result, err
}
