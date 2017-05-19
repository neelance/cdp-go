// This domain emulates different environments for the page.
package emulation

import (
	"encoding/json"
	"log"

	"github.com/neelance/cdp-go/rpc"
)

// This domain emulates different environments for the page.
type Domain struct {
	Client *rpc.Client
}

// Screen orientation.

type ScreenOrientation struct {
	// Orientation type.
	Type string `json:"type"`

	// Orientation angle.
	Angle int `json:"angle"`
}

// advance: If the scheduler runs out of immediate work, the virtual time base may fast forward to allow the next delayed task (if any) to run; pause: The virtual time base may not advance; pauseIfNetworkFetchesPending: The virtual time base may not advance if there are any pending resource fetches. (experimental)

type VirtualTimePolicy string

type SetDeviceMetricsOverrideOpts struct {
	// Overriding width value in pixels (minimum 0, maximum 10000000). 0 disables the override.
	Width int `json:"width"`

	// Overriding height value in pixels (minimum 0, maximum 10000000). 0 disables the override.
	Height int `json:"height"`

	// Overriding device scale factor value. 0 disables the override.
	DeviceScaleFactor float64 `json:"deviceScaleFactor"`

	// Whether to emulate mobile device. This includes viewport meta tag, overlay scrollbars, text autosizing and more.
	Mobile bool `json:"mobile"`

	// Whether a view that exceeds the available browser window area should be scaled down to fit.
	FitWindow bool `json:"fitWindow"`

	// Scale to apply to resulting view image. Ignored in |fitWindow| mode. (optional, experimental)
	Scale float64 `json:"scale,omitempty"`

	// Not used. (optional, experimental)
	OffsetX float64 `json:"offsetX,omitempty"`

	// Not used. (optional, experimental)
	OffsetY float64 `json:"offsetY,omitempty"`

	// Overriding screen width value in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|. (optional, experimental)
	ScreenWidth int `json:"screenWidth,omitempty"`

	// Overriding screen height value in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|. (optional, experimental)
	ScreenHeight int `json:"screenHeight,omitempty"`

	// Overriding view X position on screen in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|. (optional, experimental)
	PositionX int `json:"positionX,omitempty"`

	// Overriding view Y position on screen in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|. (optional, experimental)
	PositionY int `json:"positionY,omitempty"`

	// Screen orientation override. (optional)
	ScreenOrientation *ScreenOrientation `json:"screenOrientation,omitempty"`
}

// Overrides the values of device screen dimensions (window.screen.width, window.screen.height, window.innerWidth, window.innerHeight, and "device-width"/"device-height"-related CSS media query results).
func (d *Domain) SetDeviceMetricsOverride(opts *SetDeviceMetricsOverrideOpts) error {
	return d.Client.Call("Emulation.setDeviceMetricsOverride", opts, nil)
}

// Clears the overriden device metrics.
func (d *Domain) ClearDeviceMetricsOverride() error {
	return d.Client.Call("Emulation.clearDeviceMetricsOverride", nil, nil)
}

type ForceViewportOpts struct {
	// X coordinate of top-left corner of the area (CSS pixels).
	X float64 `json:"x"`

	// Y coordinate of top-left corner of the area (CSS pixels).
	Y float64 `json:"y"`

	// Scale to apply to the area (relative to a page scale of 1.0).
	Scale float64 `json:"scale"`
}

// Overrides the visible area of the page. The change is hidden from the page, i.e. the observable scroll position and page scale does not change. In effect, the command moves the specified area of the page into the top-left corner of the frame. (experimental)
func (d *Domain) ForceViewport(opts *ForceViewportOpts) error {
	return d.Client.Call("Emulation.forceViewport", opts, nil)
}

// Resets the visible area of the page to the original viewport, undoing any effects of the <code>forceViewport</code> command. (experimental)
func (d *Domain) ResetViewport() error {
	return d.Client.Call("Emulation.resetViewport", nil, nil)
}

// Requests that page scale factor is reset to initial values. (experimental)
func (d *Domain) ResetPageScaleFactor() error {
	return d.Client.Call("Emulation.resetPageScaleFactor", nil, nil)
}

type SetPageScaleFactorOpts struct {
	// Page scale factor.
	PageScaleFactor float64 `json:"pageScaleFactor"`
}

// Sets a specified page scale factor. (experimental)
func (d *Domain) SetPageScaleFactor(opts *SetPageScaleFactorOpts) error {
	return d.Client.Call("Emulation.setPageScaleFactor", opts, nil)
}

type SetVisibleSizeOpts struct {
	// Frame width (DIP).
	Width int `json:"width"`

	// Frame height (DIP).
	Height int `json:"height"`
}

// Resizes the frame/viewport of the page. Note that this does not affect the frame's container (e.g. browser window). Can be used to produce screenshots of the specified size. Not supported on Android. (experimental)
func (d *Domain) SetVisibleSize(opts *SetVisibleSizeOpts) error {
	return d.Client.Call("Emulation.setVisibleSize", opts, nil)
}

type SetScriptExecutionDisabledOpts struct {
	// Whether script execution should be disabled in the page.
	Value bool `json:"value"`
}

// Switches script execution in the page. (experimental)
func (d *Domain) SetScriptExecutionDisabled(opts *SetScriptExecutionDisabledOpts) error {
	return d.Client.Call("Emulation.setScriptExecutionDisabled", opts, nil)
}

type SetGeolocationOverrideOpts struct {
	// Mock latitude (optional)
	Latitude float64 `json:"latitude,omitempty"`

	// Mock longitude (optional)
	Longitude float64 `json:"longitude,omitempty"`

	// Mock accuracy (optional)
	Accuracy float64 `json:"accuracy,omitempty"`
}

// Overrides the Geolocation Position or Error. Omitting any of the parameters emulates position unavailable. (experimental)
func (d *Domain) SetGeolocationOverride(opts *SetGeolocationOverrideOpts) error {
	return d.Client.Call("Emulation.setGeolocationOverride", opts, nil)
}

// Clears the overriden Geolocation Position and Error. (experimental)
func (d *Domain) ClearGeolocationOverride() error {
	return d.Client.Call("Emulation.clearGeolocationOverride", nil, nil)
}

type SetTouchEmulationEnabledOpts struct {
	// Whether the touch event emulation should be enabled.
	Enabled bool `json:"enabled"`

	// Touch/gesture events configuration. Default: current platform. (optional)
	Configuration string `json:"configuration,omitempty"`
}

// Toggles mouse event-based touch event emulation.
func (d *Domain) SetTouchEmulationEnabled(opts *SetTouchEmulationEnabledOpts) error {
	return d.Client.Call("Emulation.setTouchEmulationEnabled", opts, nil)
}

type SetEmulatedMediaOpts struct {
	// Media type to emulate. Empty string disables the override.
	Media string `json:"media"`
}

// Emulates the given media for CSS media queries.
func (d *Domain) SetEmulatedMedia(opts *SetEmulatedMediaOpts) error {
	return d.Client.Call("Emulation.setEmulatedMedia", opts, nil)
}

type SetCPUThrottlingRateOpts struct {
	// Throttling rate as a slowdown factor (1 is no throttle, 2 is 2x slowdown, etc).
	Rate float64 `json:"rate"`
}

// Enables CPU throttling to emulate slow CPUs. (experimental)
func (d *Domain) SetCPUThrottlingRate(opts *SetCPUThrottlingRateOpts) error {
	return d.Client.Call("Emulation.setCPUThrottlingRate", opts, nil)
}

type CanEmulateResult struct {
	// True if emulation is supported.
	Result bool `json:"result"`
}

// Tells whether emulation is supported. (experimental)
func (d *Domain) CanEmulate() (*CanEmulateResult, error) {
	var result CanEmulateResult
	err := d.Client.Call("Emulation.canEmulate", nil, &result)
	return &result, err
}

type SetVirtualTimePolicyOpts struct {
	Policy *VirtualTimePolicy `json:"policy"`

	// If set, after this many virtual milliseconds have elapsed virtual time will be paused and a virtualTimeBudgetExpired event is sent. (optional)
	Budget int `json:"budget,omitempty"`
}

// Turns on virtual time for all frames (replacing real-time with a synthetic time source) and sets the current virtual time policy.  Note this supersedes any previous time budget. (experimental)
func (d *Domain) SetVirtualTimePolicy(opts *SetVirtualTimePolicyOpts) error {
	return d.Client.Call("Emulation.setVirtualTimePolicy", opts, nil)
}

type SetDefaultBackgroundColorOverrideOpts struct {
	// RGBA of the default background color. If not specified, any existing override will be cleared. (optional)
	Color interface{} `json:"color,omitempty"`
}

// Sets or clears an override of the default background color of the frame. This override is used if the content does not specify one. (experimental)
func (d *Domain) SetDefaultBackgroundColorOverride(opts *SetDefaultBackgroundColorOverrideOpts) error {
	return d.Client.Call("Emulation.setDefaultBackgroundColorOverride", opts, nil)
}

type VirtualTimeBudgetExpiredEvent struct {
}

// Notification sent after the virual time budget for the current VirtualTimePolicy has run out. (experimental)
func (d *Domain) OnVirtualTimeBudgetExpired(listener func(*VirtualTimeBudgetExpiredEvent)) {
	d.Client.AddListener("Emulation.virtualTimeBudgetExpired", func(params json.RawMessage) {
		var event VirtualTimeBudgetExpiredEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}
