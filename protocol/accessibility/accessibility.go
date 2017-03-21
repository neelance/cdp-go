// (experimental)
package accessibility

import (
	"github.com/neelance/cdp-go/rpc"
)

// (experimental)
type Domain struct {
	Client *rpc.Client
}

// Unique accessibility node identifier.
type AXNodeId interface{}

// Enum of possible property types.
type AXValueType interface{}

// Enum of possible property sources.
type AXValueSourceType interface{}

// Enum of possible native property sources (as a subtype of a particular AXValueSourceType).
type AXValueNativeSourceType interface{}

// A single source for a computed AX property.
type AXValueSource interface{}

type AXRelatedNode interface{}

type AXProperty interface{}

// A single computed AX property.
type AXValue interface{}

// States which apply to every AX node.
type AXGlobalStates interface{}

// Attributes which apply to nodes in live regions.
type AXLiveRegionAttributes interface{}

// Attributes which apply to widgets.
type AXWidgetAttributes interface{}

// States which apply to widgets.
type AXWidgetStates interface{}

// Relationships between elements other than parent/child/sibling.
type AXRelationshipAttributes interface{}

// A node in the accessibility tree.
type AXNode interface{}

type GetPartialAXTreeOpts struct {
	// ID of node to get the partial accessibility tree for.
	NodeId interface{} `json:"nodeId"`

	// Whether to fetch this nodes ancestors, siblings and children. Defaults to true. (optional)
	FetchRelatives bool `json:"fetchRelatives,omitempty"`
}

type GetPartialAXTreeResult struct {
	// The <code>Accessibility.AXNode</code> for this DOM node, if it exists, plus its ancestors, siblings and children, if requested.
	Nodes []AXNode `json:"nodes"`
}

// Fetches the accessibility node and partial accessibility tree for this DOM node, if it exists. (experimental)
func (d *Domain) GetPartialAXTree(opts *GetPartialAXTreeOpts) (*GetPartialAXTreeResult, error) {
	var result GetPartialAXTreeResult
	err := d.Client.Call("Accessibility.getPartialAXTree", opts, &result)
	return &result, err
}
