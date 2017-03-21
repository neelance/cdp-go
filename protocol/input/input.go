package input

import (
	"github.com/neelance/cdp-go/rpc"
)

type Domain struct {
	Client *rpc.Client
}

// (experimental)
type TouchPoint interface{}

// (experimental)
type GestureSourceType interface{}

type DispatchKeyEventOpts struct {
	// Type of the key event.
	Type string `json:"type"`

	// Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0). (optional)
	Modifiers int `json:"modifiers,omitempty"`

	// Time at which the event occurred. Measured in UTC time in seconds since January 1, 1970 (default: current time). (optional)
	Timestamp float64 `json:"timestamp,omitempty"`

	// Text as generated by processing a virtual key code with a keyboard layout. Not needed for for <code>keyUp</code> and <code>rawKeyDown</code> events (default: "") (optional)
	Text string `json:"text,omitempty"`

	// Text that would have been generated by the keyboard if no modifiers were pressed (except for shift). Useful for shortcut (accelerator) key handling (default: ""). (optional)
	UnmodifiedText string `json:"unmodifiedText,omitempty"`

	// Unique key identifier (e.g., 'U+0041') (default: ""). (optional)
	KeyIdentifier string `json:"keyIdentifier,omitempty"`

	// Unique DOM defined string value for each physical key (e.g., 'KeyA') (default: ""). (optional)
	Code string `json:"code,omitempty"`

	// Unique DOM defined string value describing the meaning of the key in the context of active modifiers, keyboard layout, etc (e.g., 'AltGr') (default: ""). (optional)
	Key string `json:"key,omitempty"`

	// Windows virtual key code (default: 0). (optional)
	WindowsVirtualKeyCode int `json:"windowsVirtualKeyCode,omitempty"`

	// Native virtual key code (default: 0). (optional)
	NativeVirtualKeyCode int `json:"nativeVirtualKeyCode,omitempty"`

	// Whether the event was generated from auto repeat (default: false). (optional)
	AutoRepeat bool `json:"autoRepeat,omitempty"`

	// Whether the event was generated from the keypad (default: false). (optional)
	IsKeypad bool `json:"isKeypad,omitempty"`

	// Whether the event was a system key event (default: false). (optional)
	IsSystemKey bool `json:"isSystemKey,omitempty"`
}

// Dispatches a key event to the page.
func (d *Domain) DispatchKeyEvent(opts *DispatchKeyEventOpts) error {
	return d.Client.Call("Input.dispatchKeyEvent", opts, nil)
}

type DispatchMouseEventOpts struct {
	// Type of the mouse event.
	Type string `json:"type"`

	// X coordinate of the event relative to the main frame's viewport.
	X int `json:"x"`

	// Y coordinate of the event relative to the main frame's viewport. 0 refers to the top of the viewport and Y increases as it proceeds towards the bottom of the viewport.
	Y int `json:"y"`

	// Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0). (optional)
	Modifiers int `json:"modifiers,omitempty"`

	// Time at which the event occurred. Measured in UTC time in seconds since January 1, 1970 (default: current time). (optional)
	Timestamp float64 `json:"timestamp,omitempty"`

	// Mouse button (default: "none"). (optional)
	Button string `json:"button,omitempty"`

	// Number of times the mouse button was clicked (default: 0). (optional)
	ClickCount int `json:"clickCount,omitempty"`
}

// Dispatches a mouse event to the page.
func (d *Domain) DispatchMouseEvent(opts *DispatchMouseEventOpts) error {
	return d.Client.Call("Input.dispatchMouseEvent", opts, nil)
}

type DispatchTouchEventOpts struct {
	// Type of the touch event.
	Type string `json:"type"`

	// Touch points.
	TouchPoints []TouchPoint `json:"touchPoints"`

	// Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0). (optional)
	Modifiers int `json:"modifiers,omitempty"`

	// Time at which the event occurred. Measured in UTC time in seconds since January 1, 1970 (default: current time). (optional)
	Timestamp float64 `json:"timestamp,omitempty"`
}

// Dispatches a touch event to the page. (experimental)
func (d *Domain) DispatchTouchEvent(opts *DispatchTouchEventOpts) error {
	return d.Client.Call("Input.dispatchTouchEvent", opts, nil)
}

type EmulateTouchFromMouseEventOpts struct {
	// Type of the mouse event.
	Type string `json:"type"`

	// X coordinate of the mouse pointer in DIP.
	X int `json:"x"`

	// Y coordinate of the mouse pointer in DIP.
	Y int `json:"y"`

	// Time at which the event occurred. Measured in UTC time in seconds since January 1, 1970.
	Timestamp float64 `json:"timestamp"`

	// Mouse button.
	Button string `json:"button"`

	// X delta in DIP for mouse wheel event (default: 0). (optional)
	DeltaX float64 `json:"deltaX,omitempty"`

	// Y delta in DIP for mouse wheel event (default: 0). (optional)
	DeltaY float64 `json:"deltaY,omitempty"`

	// Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0). (optional)
	Modifiers int `json:"modifiers,omitempty"`

	// Number of times the mouse button was clicked (default: 0). (optional)
	ClickCount int `json:"clickCount,omitempty"`
}

// Emulates touch event from the mouse event parameters. (experimental)
func (d *Domain) EmulateTouchFromMouseEvent(opts *EmulateTouchFromMouseEventOpts) error {
	return d.Client.Call("Input.emulateTouchFromMouseEvent", opts, nil)
}

type SynthesizePinchGestureOpts struct {
	// X coordinate of the start of the gesture in CSS pixels.
	X int `json:"x"`

	// Y coordinate of the start of the gesture in CSS pixels.
	Y int `json:"y"`

	// Relative scale factor after zooming (>1.0 zooms in, <1.0 zooms out).
	ScaleFactor float64 `json:"scaleFactor"`

	// Relative pointer speed in pixels per second (default: 800). (optional)
	RelativeSpeed int `json:"relativeSpeed,omitempty"`

	// Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type). (optional)
	GestureSourceType GestureSourceType `json:"gestureSourceType,omitempty"`
}

// Synthesizes a pinch gesture over a time period by issuing appropriate touch events. (experimental)
func (d *Domain) SynthesizePinchGesture(opts *SynthesizePinchGestureOpts) error {
	return d.Client.Call("Input.synthesizePinchGesture", opts, nil)
}

type SynthesizeScrollGestureOpts struct {
	// X coordinate of the start of the gesture in CSS pixels.
	X int `json:"x"`

	// Y coordinate of the start of the gesture in CSS pixels.
	Y int `json:"y"`

	// The distance to scroll along the X axis (positive to scroll left). (optional)
	XDistance int `json:"xDistance,omitempty"`

	// The distance to scroll along the Y axis (positive to scroll up). (optional)
	YDistance int `json:"yDistance,omitempty"`

	// The number of additional pixels to scroll back along the X axis, in addition to the given distance. (optional)
	XOverscroll int `json:"xOverscroll,omitempty"`

	// The number of additional pixels to scroll back along the Y axis, in addition to the given distance. (optional)
	YOverscroll int `json:"yOverscroll,omitempty"`

	// Prevent fling (default: true). (optional)
	PreventFling bool `json:"preventFling,omitempty"`

	// Swipe speed in pixels per second (default: 800). (optional)
	Speed int `json:"speed,omitempty"`

	// Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type). (optional)
	GestureSourceType GestureSourceType `json:"gestureSourceType,omitempty"`

	// The number of times to repeat the gesture (default: 0). (optional)
	RepeatCount int `json:"repeatCount,omitempty"`

	// The number of milliseconds delay between each repeat. (default: 250). (optional)
	RepeatDelayMs int `json:"repeatDelayMs,omitempty"`

	// The name of the interaction markers to generate, if not empty (default: ""). (optional)
	InteractionMarkerName string `json:"interactionMarkerName,omitempty"`
}

// Synthesizes a scroll gesture over a time period by issuing appropriate touch events. (experimental)
func (d *Domain) SynthesizeScrollGesture(opts *SynthesizeScrollGestureOpts) error {
	return d.Client.Call("Input.synthesizeScrollGesture", opts, nil)
}

type SynthesizeTapGestureOpts struct {
	// X coordinate of the start of the gesture in CSS pixels.
	X int `json:"x"`

	// Y coordinate of the start of the gesture in CSS pixels.
	Y int `json:"y"`

	// Duration between touchdown and touchup events in ms (default: 50). (optional)
	Duration int `json:"duration,omitempty"`

	// Number of times to perform the tap (e.g. 2 for double tap, default: 1). (optional)
	TapCount int `json:"tapCount,omitempty"`

	// Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type). (optional)
	GestureSourceType GestureSourceType `json:"gestureSourceType,omitempty"`
}

// Synthesizes a tap gesture over a time period by issuing appropriate touch events. (experimental)
func (d *Domain) SynthesizeTapGesture(opts *SynthesizeTapGestureOpts) error {
	return d.Client.Call("Input.synthesizeTapGesture", opts, nil)
}
