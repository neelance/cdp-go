// This domain allows to control rendering of the page. (experimental)
package rendering

import (
	"github.com/neelance/cdp-go/rpc"
)

// This domain allows to control rendering of the page. (experimental)
type Domain struct {
	Client *rpc.Client
}

type SetShowPaintRectsOpts struct {
	// True for showing paint rectangles
	Result bool `json:"result"`
}

// Requests that backend shows paint rectangles
func (d *Domain) SetShowPaintRects(opts *SetShowPaintRectsOpts) error {
	return d.Client.Call("Rendering.setShowPaintRects", opts, nil)
}

type SetShowDebugBordersOpts struct {
	// True for showing debug borders
	Show bool `json:"show"`
}

// Requests that backend shows debug borders on layers
func (d *Domain) SetShowDebugBorders(opts *SetShowDebugBordersOpts) error {
	return d.Client.Call("Rendering.setShowDebugBorders", opts, nil)
}

type SetShowFPSCounterOpts struct {
	// True for showing the FPS counter
	Show bool `json:"show"`
}

// Requests that backend shows the FPS counter
func (d *Domain) SetShowFPSCounter(opts *SetShowFPSCounterOpts) error {
	return d.Client.Call("Rendering.setShowFPSCounter", opts, nil)
}

type SetShowScrollBottleneckRectsOpts struct {
	// True for showing scroll bottleneck rects
	Show bool `json:"show"`
}

// Requests that backend shows scroll bottleneck rects
func (d *Domain) SetShowScrollBottleneckRects(opts *SetShowScrollBottleneckRectsOpts) error {
	return d.Client.Call("Rendering.setShowScrollBottleneckRects", opts, nil)
}

type SetShowViewportSizeOnResizeOpts struct {
	// Whether to paint size or not.
	Show bool `json:"show"`
}

// Paints viewport size upon main frame resize.
func (d *Domain) SetShowViewportSizeOnResize(opts *SetShowViewportSizeOnResizeOpts) error {
	return d.Client.Call("Rendering.setShowViewportSizeOnResize", opts, nil)
}
