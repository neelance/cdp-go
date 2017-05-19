// (experimental)
package layertree

import (
	"encoding/json"
	"log"

	"github.com/neelance/cdp-go/rpc"
)

// (experimental)
type Domain struct {
	Client *rpc.Client
}

// Unique Layer identifier.

type LayerId string

// Unique snapshot identifier.

type SnapshotId string

// Rectangle where scrolling happens on the main thread.

type ScrollRect struct {
	// Rectangle itself.
	Rect interface{} `json:"rect"`

	// Reason for rectangle to force scrolling on the main thread
	Type string `json:"type"`
}

// Serialized fragment of layer picture along with its offset within the layer.

type PictureTile struct {
	// Offset from owning layer left boundary
	X float64 `json:"x"`

	// Offset from owning layer top boundary
	Y float64 `json:"y"`

	// Base64-encoded snapshot data.
	Picture string `json:"picture"`
}

// Information about a compositing layer.

type Layer struct {
	// The unique id for this layer.
	LayerId LayerId `json:"layerId"`

	// The id of parent (not present for root). (optional)
	ParentLayerId LayerId `json:"parentLayerId,omitempty"`

	// The backend id for the node associated with this layer. (optional)
	BackendNodeId interface{} `json:"backendNodeId,omitempty"`

	// Offset from parent layer, X coordinate.
	OffsetX float64 `json:"offsetX"`

	// Offset from parent layer, Y coordinate.
	OffsetY float64 `json:"offsetY"`

	// Layer width.
	Width float64 `json:"width"`

	// Layer height.
	Height float64 `json:"height"`

	// Transformation matrix for layer, default is identity matrix (optional)
	Transform []float64 `json:"transform,omitempty"`

	// Transform anchor point X, absent if no transform specified (optional)
	AnchorX float64 `json:"anchorX,omitempty"`

	// Transform anchor point Y, absent if no transform specified (optional)
	AnchorY float64 `json:"anchorY,omitempty"`

	// Transform anchor point Z, absent if no transform specified (optional)
	AnchorZ float64 `json:"anchorZ,omitempty"`

	// Indicates how many time this layer has painted.
	PaintCount int `json:"paintCount"`

	// Indicates whether this layer hosts any content, rather than being used for transform/scrolling purposes only.
	DrawsContent bool `json:"drawsContent"`

	// Set if layer is not visible. (optional)
	Invisible bool `json:"invisible,omitempty"`

	// Rectangles scrolling on main thread only. (optional)
	ScrollRects []*ScrollRect `json:"scrollRects,omitempty"`
}

// Array of timings, one per paint step.

type PaintProfile []float64

// Enables compositing tree inspection.
func (d *Domain) Enable() error {
	return d.Client.Call("LayerTree.enable", nil, nil)
}

// Disables compositing tree inspection.
func (d *Domain) Disable() error {
	return d.Client.Call("LayerTree.disable", nil, nil)
}

type CompositingReasonsOpts struct {
	// The id of the layer for which we want to get the reasons it was composited.
	LayerId LayerId `json:"layerId"`
}

type CompositingReasonsResult struct {
	// A list of strings specifying reasons for the given layer to become composited.
	CompositingReasons []string `json:"compositingReasons"`
}

// Provides the reasons why the given layer was composited.
func (d *Domain) CompositingReasons(opts *CompositingReasonsOpts) (*CompositingReasonsResult, error) {
	var result CompositingReasonsResult
	err := d.Client.Call("LayerTree.compositingReasons", opts, &result)
	return &result, err
}

type MakeSnapshotOpts struct {
	// The id of the layer.
	LayerId LayerId `json:"layerId"`
}

type MakeSnapshotResult struct {
	// The id of the layer snapshot.
	SnapshotId SnapshotId `json:"snapshotId"`
}

// Returns the layer snapshot identifier.
func (d *Domain) MakeSnapshot(opts *MakeSnapshotOpts) (*MakeSnapshotResult, error) {
	var result MakeSnapshotResult
	err := d.Client.Call("LayerTree.makeSnapshot", opts, &result)
	return &result, err
}

type LoadSnapshotOpts struct {
	// An array of tiles composing the snapshot.
	Tiles []*PictureTile `json:"tiles"`
}

type LoadSnapshotResult struct {
	// The id of the snapshot.
	SnapshotId SnapshotId `json:"snapshotId"`
}

// Returns the snapshot identifier.
func (d *Domain) LoadSnapshot(opts *LoadSnapshotOpts) (*LoadSnapshotResult, error) {
	var result LoadSnapshotResult
	err := d.Client.Call("LayerTree.loadSnapshot", opts, &result)
	return &result, err
}

type ReleaseSnapshotOpts struct {
	// The id of the layer snapshot.
	SnapshotId SnapshotId `json:"snapshotId"`
}

// Releases layer snapshot captured by the back-end.
func (d *Domain) ReleaseSnapshot(opts *ReleaseSnapshotOpts) error {
	return d.Client.Call("LayerTree.releaseSnapshot", opts, nil)
}

type ProfileSnapshotOpts struct {
	// The id of the layer snapshot.
	SnapshotId SnapshotId `json:"snapshotId"`

	// The maximum number of times to replay the snapshot (1, if not specified). (optional)
	MinRepeatCount int `json:"minRepeatCount,omitempty"`

	// The minimum duration (in seconds) to replay the snapshot. (optional)
	MinDuration float64 `json:"minDuration,omitempty"`

	// The clip rectangle to apply when replaying the snapshot. (optional)
	ClipRect interface{} `json:"clipRect,omitempty"`
}

type ProfileSnapshotResult struct {
	// The array of paint profiles, one per run.
	Timings []PaintProfile `json:"timings"`
}

func (d *Domain) ProfileSnapshot(opts *ProfileSnapshotOpts) (*ProfileSnapshotResult, error) {
	var result ProfileSnapshotResult
	err := d.Client.Call("LayerTree.profileSnapshot", opts, &result)
	return &result, err
}

type ReplaySnapshotOpts struct {
	// The id of the layer snapshot.
	SnapshotId SnapshotId `json:"snapshotId"`

	// The first step to replay from (replay from the very start if not specified). (optional)
	FromStep int `json:"fromStep,omitempty"`

	// The last step to replay to (replay till the end if not specified). (optional)
	ToStep int `json:"toStep,omitempty"`

	// The scale to apply while replaying (defaults to 1). (optional)
	Scale float64 `json:"scale,omitempty"`
}

type ReplaySnapshotResult struct {
	// A data: URL for resulting image.
	DataURL string `json:"dataURL"`
}

// Replays the layer snapshot and returns the resulting bitmap.
func (d *Domain) ReplaySnapshot(opts *ReplaySnapshotOpts) (*ReplaySnapshotResult, error) {
	var result ReplaySnapshotResult
	err := d.Client.Call("LayerTree.replaySnapshot", opts, &result)
	return &result, err
}

type SnapshotCommandLogOpts struct {
	// The id of the layer snapshot.
	SnapshotId SnapshotId `json:"snapshotId"`
}

type SnapshotCommandLogResult struct {
	// The array of canvas function calls.
	CommandLog []interface{} `json:"commandLog"`
}

// Replays the layer snapshot and returns canvas log.
func (d *Domain) SnapshotCommandLog(opts *SnapshotCommandLogOpts) (*SnapshotCommandLogResult, error) {
	var result SnapshotCommandLogResult
	err := d.Client.Call("LayerTree.snapshotCommandLog", opts, &result)
	return &result, err
}

type LayerTreeDidChangeEvent struct {
	// Layer tree, absent if not in the comspositing mode. (optional)
	Layers []*Layer `json:"layers"`
}

func (d *Domain) OnLayerTreeDidChange(listener func(*LayerTreeDidChangeEvent)) {
	d.Client.AddListener("LayerTree.layerTreeDidChange", func(params json.RawMessage) {
		var event LayerTreeDidChangeEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type LayerPaintedEvent struct {
	// The id of the painted layer.
	LayerId LayerId `json:"layerId"`

	// Clip rectangle.
	Clip interface{} `json:"clip"`
}

func (d *Domain) OnLayerPainted(listener func(*LayerPaintedEvent)) {
	d.Client.AddListener("LayerTree.layerPainted", func(params json.RawMessage) {
		var event LayerPaintedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}
