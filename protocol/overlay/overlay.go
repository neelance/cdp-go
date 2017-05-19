// This domain provides various functionality related to drawing atop the inspected page. (experimental)
package overlay

import (
	"encoding/json"
	"log"

	"github.com/neelance/cdp-go/rpc"
)

// This domain provides various functionality related to drawing atop the inspected page. (experimental)
type Domain struct {
	Client *rpc.Client
}

// Configuration data for the highlighting of page elements.

type HighlightConfig struct {
	// Whether the node info tooltip should be shown (default: false). (optional)
	ShowInfo bool `json:"showInfo,omitempty"`

	// Whether the rulers should be shown (default: false). (optional)
	ShowRulers bool `json:"showRulers,omitempty"`

	// Whether the extension lines from node to the rulers should be shown (default: false). (optional)
	ShowExtensionLines bool `json:"showExtensionLines,omitempty"`

	// (optional)
	DisplayAsMaterial bool `json:"displayAsMaterial,omitempty"`

	// The content box highlight fill color (default: transparent). (optional)
	ContentColor interface{} `json:"contentColor,omitempty"`

	// The padding highlight fill color (default: transparent). (optional)
	PaddingColor interface{} `json:"paddingColor,omitempty"`

	// The border highlight fill color (default: transparent). (optional)
	BorderColor interface{} `json:"borderColor,omitempty"`

	// The margin highlight fill color (default: transparent). (optional)
	MarginColor interface{} `json:"marginColor,omitempty"`

	// The event target element highlight fill color (default: transparent). (optional)
	EventTargetColor interface{} `json:"eventTargetColor,omitempty"`

	// The shape outside fill color (default: transparent). (optional)
	ShapeColor interface{} `json:"shapeColor,omitempty"`

	// The shape margin fill color (default: transparent). (optional)
	ShapeMarginColor interface{} `json:"shapeMarginColor,omitempty"`

	// Selectors to highlight relevant nodes. (optional)
	SelectorList string `json:"selectorList,omitempty"`
}

type InspectMode string

// Enables domain notifications.
func (d *Domain) Enable() error {
	return d.Client.Call("Overlay.enable", nil, nil)
}

// Disables domain notifications.
func (d *Domain) Disable() error {
	return d.Client.Call("Overlay.disable", nil, nil)
}

type SetShowPaintRectsOpts struct {
	// True for showing paint rectangles
	Result bool `json:"result"`
}

// Requests that backend shows paint rectangles
func (d *Domain) SetShowPaintRects(opts *SetShowPaintRectsOpts) error {
	return d.Client.Call("Overlay.setShowPaintRects", opts, nil)
}

type SetShowDebugBordersOpts struct {
	// True for showing debug borders
	Show bool `json:"show"`
}

// Requests that backend shows debug borders on layers
func (d *Domain) SetShowDebugBorders(opts *SetShowDebugBordersOpts) error {
	return d.Client.Call("Overlay.setShowDebugBorders", opts, nil)
}

type SetShowFPSCounterOpts struct {
	// True for showing the FPS counter
	Show bool `json:"show"`
}

// Requests that backend shows the FPS counter
func (d *Domain) SetShowFPSCounter(opts *SetShowFPSCounterOpts) error {
	return d.Client.Call("Overlay.setShowFPSCounter", opts, nil)
}

type SetShowScrollBottleneckRectsOpts struct {
	// True for showing scroll bottleneck rects
	Show bool `json:"show"`
}

// Requests that backend shows scroll bottleneck rects
func (d *Domain) SetShowScrollBottleneckRects(opts *SetShowScrollBottleneckRectsOpts) error {
	return d.Client.Call("Overlay.setShowScrollBottleneckRects", opts, nil)
}

type SetShowViewportSizeOnResizeOpts struct {
	// Whether to paint size or not.
	Show bool `json:"show"`
}

// Paints viewport size upon main frame resize.
func (d *Domain) SetShowViewportSizeOnResize(opts *SetShowViewportSizeOnResizeOpts) error {
	return d.Client.Call("Overlay.setShowViewportSizeOnResize", opts, nil)
}

type SetPausedInDebuggerMessageOpts struct {
	// The message to display, also triggers resume and step over controls. (optional)
	Message string `json:"message,omitempty"`
}

func (d *Domain) SetPausedInDebuggerMessage(opts *SetPausedInDebuggerMessageOpts) error {
	return d.Client.Call("Overlay.setPausedInDebuggerMessage", opts, nil)
}

type SetSuspendedOpts struct {
	// Whether overlay should be suspended and not consume any resources until resumed.
	Suspended bool `json:"suspended"`
}

func (d *Domain) SetSuspended(opts *SetSuspendedOpts) error {
	return d.Client.Call("Overlay.setSuspended", opts, nil)
}

type SetInspectModeOpts struct {
	// Set an inspection mode.
	Mode InspectMode `json:"mode"`

	// A descriptor for the highlight appearance of hovered-over nodes. May be omitted if <code>enabled == false</code>. (optional)
	HighlightConfig *HighlightConfig `json:"highlightConfig,omitempty"`
}

// Enters the 'inspect' mode. In this mode, elements that user is hovering over are highlighted. Backend then generates 'inspectNodeRequested' event upon element selection.
func (d *Domain) SetInspectMode(opts *SetInspectModeOpts) error {
	return d.Client.Call("Overlay.setInspectMode", opts, nil)
}

type HighlightRectOpts struct {
	// X coordinate
	X int `json:"x"`

	// Y coordinate
	Y int `json:"y"`

	// Rectangle width
	Width int `json:"width"`

	// Rectangle height
	Height int `json:"height"`

	// The highlight fill color (default: transparent). (optional)
	Color interface{} `json:"color,omitempty"`

	// The highlight outline color (default: transparent). (optional)
	OutlineColor interface{} `json:"outlineColor,omitempty"`
}

// Highlights given rectangle. Coordinates are absolute with respect to the main frame viewport.
func (d *Domain) HighlightRect(opts *HighlightRectOpts) error {
	return d.Client.Call("Overlay.highlightRect", opts, nil)
}

type HighlightQuadOpts struct {
	// Quad to highlight
	Quad interface{} `json:"quad"`

	// The highlight fill color (default: transparent). (optional)
	Color interface{} `json:"color,omitempty"`

	// The highlight outline color (default: transparent). (optional)
	OutlineColor interface{} `json:"outlineColor,omitempty"`
}

// Highlights given quad. Coordinates are absolute with respect to the main frame viewport.
func (d *Domain) HighlightQuad(opts *HighlightQuadOpts) error {
	return d.Client.Call("Overlay.highlightQuad", opts, nil)
}

type HighlightNodeOpts struct {
	// A descriptor for the highlight appearance.
	HighlightConfig *HighlightConfig `json:"highlightConfig"`

	// Identifier of the node to highlight. (optional)
	NodeId interface{} `json:"nodeId,omitempty"`

	// Identifier of the backend node to highlight. (optional)
	BackendNodeId interface{} `json:"backendNodeId,omitempty"`

	// JavaScript object id of the node to be highlighted. (optional)
	ObjectId interface{} `json:"objectId,omitempty"`
}

// Highlights DOM node with given id or with the given JavaScript object wrapper. Either nodeId or objectId must be specified.
func (d *Domain) HighlightNode(opts *HighlightNodeOpts) error {
	return d.Client.Call("Overlay.highlightNode", opts, nil)
}

type HighlightFrameOpts struct {
	// Identifier of the frame to highlight.
	FrameId interface{} `json:"frameId"`

	// The content box highlight fill color (default: transparent). (optional)
	ContentColor interface{} `json:"contentColor,omitempty"`

	// The content box highlight outline color (default: transparent). (optional)
	ContentOutlineColor interface{} `json:"contentOutlineColor,omitempty"`
}

// Highlights owner element of the frame with given id.
func (d *Domain) HighlightFrame(opts *HighlightFrameOpts) error {
	return d.Client.Call("Overlay.highlightFrame", opts, nil)
}

// Hides any highlight.
func (d *Domain) HideHighlight() error {
	return d.Client.Call("Overlay.hideHighlight", nil, nil)
}

type GetHighlightObjectForTestOpts struct {
	// Id of the node to get highlight object for.
	NodeId interface{} `json:"nodeId"`
}

type GetHighlightObjectForTestResult struct {
	// Highlight data for the node.
	Highlight interface{} `json:"highlight"`
}

// For testing.
func (d *Domain) GetHighlightObjectForTest(opts *GetHighlightObjectForTestOpts) (*GetHighlightObjectForTestResult, error) {
	var result GetHighlightObjectForTestResult
	err := d.Client.Call("Overlay.getHighlightObjectForTest", opts, &result)
	return &result, err
}

type NodeHighlightRequestedEvent struct {
	NodeId interface{} `json:"nodeId"`
}

// Fired when the node should be highlighted. This happens after call to <code>setInspectMode</code>.
func (d *Domain) OnNodeHighlightRequested(listener func(*NodeHighlightRequestedEvent)) {
	d.Client.AddListener("Overlay.nodeHighlightRequested", func(params json.RawMessage) {
		var event NodeHighlightRequestedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type InspectNodeRequestedEvent struct {
	// Id of the node to inspect.
	BackendNodeId interface{} `json:"backendNodeId"`
}

// Fired when the node should be inspected. This happens after call to <code>setInspectMode</code> or when user manually inspects an element.
func (d *Domain) OnInspectNodeRequested(listener func(*InspectNodeRequestedEvent)) {
	d.Client.AddListener("Overlay.inspectNodeRequested", func(params json.RawMessage) {
		var event InspectNodeRequestedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}
