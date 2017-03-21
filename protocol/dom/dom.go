// This domain exposes DOM read/write operations. Each DOM Node is represented with its mirror object that has an <code>id</code>. This <code>id</code> can be used to get additional information on the Node, resolve it into the JavaScript object wrapper, etc. It is important that client receives DOM events only for the nodes that are known to the client. Backend keeps track of the nodes that were sent to the client and never sends the same node twice. It is client's responsibility to collect information about the nodes that were sent to the client.<p>Note that <code>iframe</code> owner elements will return corresponding document elements as their child nodes.</p>
package dom

import (
	"encoding/json"
	"log"

	"github.com/neelance/cdp-go/rpc"
)

// This domain exposes DOM read/write operations. Each DOM Node is represented with its mirror object that has an <code>id</code>. This <code>id</code> can be used to get additional information on the Node, resolve it into the JavaScript object wrapper, etc. It is important that client receives DOM events only for the nodes that are known to the client. Backend keeps track of the nodes that were sent to the client and never sends the same node twice. It is client's responsibility to collect information about the nodes that were sent to the client.<p>Note that <code>iframe</code> owner elements will return corresponding document elements as their child nodes.</p>
type Domain struct {
	Client *rpc.Client
}

// Unique DOM node identifier.
type NodeId interface{}

// Unique DOM node identifier used to reference a node that may not have been pushed to the front-end. (experimental)
type BackendNodeId interface{}

// Backend node with a friendly name. (experimental)
type BackendNode interface{}

// Pseudo element type.
type PseudoType interface{}

// Shadow root type.
type ShadowRootType interface{}

// DOM interaction is implemented in terms of mirror objects that represent the actual DOM nodes. DOMNode is a base node mirror type.
type Node interface{}

// A structure holding an RGBA color.
type RGBA interface{}

// An array of quad vertices, x immediately followed by y for each point, points clock-wise. (experimental)
type Quad interface{}

// Box model. (experimental)
type BoxModel interface{}

// CSS Shape Outside details. (experimental)
type ShapeOutsideInfo interface{}

// Rectangle. (experimental)
type Rect interface{}

// Configuration data for the highlighting of page elements.
type HighlightConfig interface{}

// (experimental)
type InspectMode interface{}

// Enables DOM agent for the given page.
func (d *Domain) Enable() error {
	return d.Client.Call("DOM.enable", nil, nil)
}

// Disables DOM agent for the given page.
func (d *Domain) Disable() error {
	return d.Client.Call("DOM.disable", nil, nil)
}

type GetDocumentOpts struct {
	// The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0. (optional, experimental)
	Depth int `json:"depth,omitempty"`

	// Whether or not iframes and shadow roots should be traversed when returning the subtree (default is false). (optional, experimental)
	Pierce bool `json:"pierce,omitempty"`
}

type GetDocumentResult struct {
	// Resulting node.
	Root Node `json:"root"`
}

// Returns the root DOM node (and optionally the subtree) to the caller.
func (d *Domain) GetDocument(opts *GetDocumentOpts) (*GetDocumentResult, error) {
	var result GetDocumentResult
	err := d.Client.Call("DOM.getDocument", opts, &result)
	return &result, err
}

type GetFlattenedDocumentOpts struct {
	// The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0. (optional, experimental)
	Depth int `json:"depth,omitempty"`

	// Whether or not iframes and shadow roots should be traversed when returning the subtree (default is false). (optional, experimental)
	Pierce bool `json:"pierce,omitempty"`
}

type GetFlattenedDocumentResult struct {
	// Resulting node.
	Nodes []Node `json:"nodes"`
}

// Returns the root DOM node (and optionally the subtree) to the caller.
func (d *Domain) GetFlattenedDocument(opts *GetFlattenedDocumentOpts) (*GetFlattenedDocumentResult, error) {
	var result GetFlattenedDocumentResult
	err := d.Client.Call("DOM.getFlattenedDocument", opts, &result)
	return &result, err
}

type CollectClassNamesFromSubtreeOpts struct {
	// Id of the node to collect class names.
	NodeId NodeId `json:"nodeId"`
}

type CollectClassNamesFromSubtreeResult struct {
	// Class name list.
	ClassNames []string `json:"classNames"`
}

// Collects class names for the node with given id and all of it's child nodes. (experimental)
func (d *Domain) CollectClassNamesFromSubtree(opts *CollectClassNamesFromSubtreeOpts) (*CollectClassNamesFromSubtreeResult, error) {
	var result CollectClassNamesFromSubtreeResult
	err := d.Client.Call("DOM.collectClassNamesFromSubtree", opts, &result)
	return &result, err
}

type RequestChildNodesOpts struct {
	// Id of the node to get children for.
	NodeId NodeId `json:"nodeId"`

	// The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0. (optional, experimental)
	Depth int `json:"depth,omitempty"`

	// Whether or not iframes and shadow roots should be traversed when returning the sub-tree (default is false). (optional, experimental)
	Pierce bool `json:"pierce,omitempty"`
}

// Requests that children of the node with given id are returned to the caller in form of <code>setChildNodes</code> events where not only immediate children are retrieved, but all children down to the specified depth.
func (d *Domain) RequestChildNodes(opts *RequestChildNodesOpts) error {
	return d.Client.Call("DOM.requestChildNodes", opts, nil)
}

type QuerySelectorOpts struct {
	// Id of the node to query upon.
	NodeId NodeId `json:"nodeId"`

	// Selector string.
	Selector string `json:"selector"`
}

type QuerySelectorResult struct {
	// Query selector result.
	NodeId NodeId `json:"nodeId"`
}

// Executes <code>querySelector</code> on a given node.
func (d *Domain) QuerySelector(opts *QuerySelectorOpts) (*QuerySelectorResult, error) {
	var result QuerySelectorResult
	err := d.Client.Call("DOM.querySelector", opts, &result)
	return &result, err
}

type QuerySelectorAllOpts struct {
	// Id of the node to query upon.
	NodeId NodeId `json:"nodeId"`

	// Selector string.
	Selector string `json:"selector"`
}

type QuerySelectorAllResult struct {
	// Query selector result.
	NodeIds []NodeId `json:"nodeIds"`
}

// Executes <code>querySelectorAll</code> on a given node.
func (d *Domain) QuerySelectorAll(opts *QuerySelectorAllOpts) (*QuerySelectorAllResult, error) {
	var result QuerySelectorAllResult
	err := d.Client.Call("DOM.querySelectorAll", opts, &result)
	return &result, err
}

type SetNodeNameOpts struct {
	// Id of the node to set name for.
	NodeId NodeId `json:"nodeId"`

	// New node's name.
	Name string `json:"name"`
}

type SetNodeNameResult struct {
	// New node's id.
	NodeId NodeId `json:"nodeId"`
}

// Sets node name for a node with given id.
func (d *Domain) SetNodeName(opts *SetNodeNameOpts) (*SetNodeNameResult, error) {
	var result SetNodeNameResult
	err := d.Client.Call("DOM.setNodeName", opts, &result)
	return &result, err
}

type SetNodeValueOpts struct {
	// Id of the node to set value for.
	NodeId NodeId `json:"nodeId"`

	// New node's value.
	Value string `json:"value"`
}

// Sets node value for a node with given id.
func (d *Domain) SetNodeValue(opts *SetNodeValueOpts) error {
	return d.Client.Call("DOM.setNodeValue", opts, nil)
}

type RemoveNodeOpts struct {
	// Id of the node to remove.
	NodeId NodeId `json:"nodeId"`
}

// Removes node with given id.
func (d *Domain) RemoveNode(opts *RemoveNodeOpts) error {
	return d.Client.Call("DOM.removeNode", opts, nil)
}

type SetAttributeValueOpts struct {
	// Id of the element to set attribute for.
	NodeId NodeId `json:"nodeId"`

	// Attribute name.
	Name string `json:"name"`

	// Attribute value.
	Value string `json:"value"`
}

// Sets attribute for an element with given id.
func (d *Domain) SetAttributeValue(opts *SetAttributeValueOpts) error {
	return d.Client.Call("DOM.setAttributeValue", opts, nil)
}

type SetAttributesAsTextOpts struct {
	// Id of the element to set attributes for.
	NodeId NodeId `json:"nodeId"`

	// Text with a number of attributes. Will parse this text using HTML parser.
	Text string `json:"text"`

	// Attribute name to replace with new attributes derived from text in case text parsed successfully. (optional)
	Name string `json:"name,omitempty"`
}

// Sets attributes on element with given id. This method is useful when user edits some existing attribute value and types in several attribute name/value pairs.
func (d *Domain) SetAttributesAsText(opts *SetAttributesAsTextOpts) error {
	return d.Client.Call("DOM.setAttributesAsText", opts, nil)
}

type RemoveAttributeOpts struct {
	// Id of the element to remove attribute from.
	NodeId NodeId `json:"nodeId"`

	// Name of the attribute to remove.
	Name string `json:"name"`
}

// Removes attribute with given name from an element with given id.
func (d *Domain) RemoveAttribute(opts *RemoveAttributeOpts) error {
	return d.Client.Call("DOM.removeAttribute", opts, nil)
}

type GetOuterHTMLOpts struct {
	// Id of the node to get markup for.
	NodeId NodeId `json:"nodeId"`
}

type GetOuterHTMLResult struct {
	// Outer HTML markup.
	OuterHTML string `json:"outerHTML"`
}

// Returns node's HTML markup.
func (d *Domain) GetOuterHTML(opts *GetOuterHTMLOpts) (*GetOuterHTMLResult, error) {
	var result GetOuterHTMLResult
	err := d.Client.Call("DOM.getOuterHTML", opts, &result)
	return &result, err
}

type SetOuterHTMLOpts struct {
	// Id of the node to set markup for.
	NodeId NodeId `json:"nodeId"`

	// Outer HTML markup to set.
	OuterHTML string `json:"outerHTML"`
}

// Sets node HTML markup, returns new node id.
func (d *Domain) SetOuterHTML(opts *SetOuterHTMLOpts) error {
	return d.Client.Call("DOM.setOuterHTML", opts, nil)
}

type PerformSearchOpts struct {
	// Plain text or query selector or XPath search query.
	Query string `json:"query"`

	// True to search in user agent shadow DOM. (optional, experimental)
	IncludeUserAgentShadowDOM bool `json:"includeUserAgentShadowDOM,omitempty"`
}

type PerformSearchResult struct {
	// Unique search session identifier.
	SearchId string `json:"searchId"`

	// Number of search results.
	ResultCount int `json:"resultCount"`
}

// Searches for a given string in the DOM tree. Use <code>getSearchResults</code> to access search results or <code>cancelSearch</code> to end this search session. (experimental)
func (d *Domain) PerformSearch(opts *PerformSearchOpts) (*PerformSearchResult, error) {
	var result PerformSearchResult
	err := d.Client.Call("DOM.performSearch", opts, &result)
	return &result, err
}

type GetSearchResultsOpts struct {
	// Unique search session identifier.
	SearchId string `json:"searchId"`

	// Start index of the search result to be returned.
	FromIndex int `json:"fromIndex"`

	// End index of the search result to be returned.
	ToIndex int `json:"toIndex"`
}

type GetSearchResultsResult struct {
	// Ids of the search result nodes.
	NodeIds []NodeId `json:"nodeIds"`
}

// Returns search results from given <code>fromIndex</code> to given <code>toIndex</code> from the sarch with the given identifier. (experimental)
func (d *Domain) GetSearchResults(opts *GetSearchResultsOpts) (*GetSearchResultsResult, error) {
	var result GetSearchResultsResult
	err := d.Client.Call("DOM.getSearchResults", opts, &result)
	return &result, err
}

type DiscardSearchResultsOpts struct {
	// Unique search session identifier.
	SearchId string `json:"searchId"`
}

// Discards search results from the session with the given id. <code>getSearchResults</code> should no longer be called for that search. (experimental)
func (d *Domain) DiscardSearchResults(opts *DiscardSearchResultsOpts) error {
	return d.Client.Call("DOM.discardSearchResults", opts, nil)
}

type RequestNodeOpts struct {
	// JavaScript object id to convert into node.
	ObjectId interface{} `json:"objectId"`
}

type RequestNodeResult struct {
	// Node id for given object.
	NodeId NodeId `json:"nodeId"`
}

// Requests that the node is sent to the caller given the JavaScript node object reference. All nodes that form the path from the node to the root are also sent to the client as a series of <code>setChildNodes</code> notifications.
func (d *Domain) RequestNode(opts *RequestNodeOpts) (*RequestNodeResult, error) {
	var result RequestNodeResult
	err := d.Client.Call("DOM.requestNode", opts, &result)
	return &result, err
}

type SetInspectModeOpts struct {
	// Set an inspection mode.
	Mode InspectMode `json:"mode"`

	// A descriptor for the highlight appearance of hovered-over nodes. May be omitted if <code>enabled == false</code>. (optional)
	HighlightConfig HighlightConfig `json:"highlightConfig,omitempty"`
}

// Enters the 'inspect' mode. In this mode, elements that user is hovering over are highlighted. Backend then generates 'inspectNodeRequested' event upon element selection. (experimental)
func (d *Domain) SetInspectMode(opts *SetInspectModeOpts) error {
	return d.Client.Call("DOM.setInspectMode", opts, nil)
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
	Color RGBA `json:"color,omitempty"`

	// The highlight outline color (default: transparent). (optional)
	OutlineColor RGBA `json:"outlineColor,omitempty"`
}

// Highlights given rectangle. Coordinates are absolute with respect to the main frame viewport.
func (d *Domain) HighlightRect(opts *HighlightRectOpts) error {
	return d.Client.Call("DOM.highlightRect", opts, nil)
}

type HighlightQuadOpts struct {
	// Quad to highlight
	Quad Quad `json:"quad"`

	// The highlight fill color (default: transparent). (optional)
	Color RGBA `json:"color,omitempty"`

	// The highlight outline color (default: transparent). (optional)
	OutlineColor RGBA `json:"outlineColor,omitempty"`
}

// Highlights given quad. Coordinates are absolute with respect to the main frame viewport. (experimental)
func (d *Domain) HighlightQuad(opts *HighlightQuadOpts) error {
	return d.Client.Call("DOM.highlightQuad", opts, nil)
}

type HighlightNodeOpts struct {
	// A descriptor for the highlight appearance.
	HighlightConfig HighlightConfig `json:"highlightConfig"`

	// Identifier of the node to highlight. (optional)
	NodeId NodeId `json:"nodeId,omitempty"`

	// Identifier of the backend node to highlight. (optional)
	BackendNodeId BackendNodeId `json:"backendNodeId,omitempty"`

	// JavaScript object id of the node to be highlighted. (optional, experimental)
	ObjectId interface{} `json:"objectId,omitempty"`
}

// Highlights DOM node with given id or with the given JavaScript object wrapper. Either nodeId or objectId must be specified.
func (d *Domain) HighlightNode(opts *HighlightNodeOpts) error {
	return d.Client.Call("DOM.highlightNode", opts, nil)
}

// Hides DOM node highlight.
func (d *Domain) HideHighlight() error {
	return d.Client.Call("DOM.hideHighlight", nil, nil)
}

type HighlightFrameOpts struct {
	// Identifier of the frame to highlight.
	FrameId interface{} `json:"frameId"`

	// The content box highlight fill color (default: transparent). (optional)
	ContentColor RGBA `json:"contentColor,omitempty"`

	// The content box highlight outline color (default: transparent). (optional)
	ContentOutlineColor RGBA `json:"contentOutlineColor,omitempty"`
}

// Highlights owner element of the frame with given id. (experimental)
func (d *Domain) HighlightFrame(opts *HighlightFrameOpts) error {
	return d.Client.Call("DOM.highlightFrame", opts, nil)
}

type PushNodeByPathToFrontendOpts struct {
	// Path to node in the proprietary format.
	Path string `json:"path"`
}

type PushNodeByPathToFrontendResult struct {
	// Id of the node for given path.
	NodeId NodeId `json:"nodeId"`
}

// Requests that the node is sent to the caller given its path. // FIXME, use XPath (experimental)
func (d *Domain) PushNodeByPathToFrontend(opts *PushNodeByPathToFrontendOpts) (*PushNodeByPathToFrontendResult, error) {
	var result PushNodeByPathToFrontendResult
	err := d.Client.Call("DOM.pushNodeByPathToFrontend", opts, &result)
	return &result, err
}

type PushNodesByBackendIdsToFrontendOpts struct {
	// The array of backend node ids.
	BackendNodeIds []BackendNodeId `json:"backendNodeIds"`
}

type PushNodesByBackendIdsToFrontendResult struct {
	// The array of ids of pushed nodes that correspond to the backend ids specified in backendNodeIds.
	NodeIds []NodeId `json:"nodeIds"`
}

// Requests that a batch of nodes is sent to the caller given their backend node ids. (experimental)
func (d *Domain) PushNodesByBackendIdsToFrontend(opts *PushNodesByBackendIdsToFrontendOpts) (*PushNodesByBackendIdsToFrontendResult, error) {
	var result PushNodesByBackendIdsToFrontendResult
	err := d.Client.Call("DOM.pushNodesByBackendIdsToFrontend", opts, &result)
	return &result, err
}

type SetInspectedNodeOpts struct {
	// DOM node id to be accessible by means of $x command line API.
	NodeId NodeId `json:"nodeId"`
}

// Enables console to refer to the node with given id via $x (see Command Line API for more details $x functions). (experimental)
func (d *Domain) SetInspectedNode(opts *SetInspectedNodeOpts) error {
	return d.Client.Call("DOM.setInspectedNode", opts, nil)
}

type ResolveNodeOpts struct {
	// Id of the node to resolve.
	NodeId NodeId `json:"nodeId"`

	// Symbolic group name that can be used to release multiple objects. (optional)
	ObjectGroup string `json:"objectGroup,omitempty"`
}

type ResolveNodeResult struct {
	// JavaScript object wrapper for given node.
	Object interface{} `json:"object"`
}

// Resolves JavaScript node object for given node id.
func (d *Domain) ResolveNode(opts *ResolveNodeOpts) (*ResolveNodeResult, error) {
	var result ResolveNodeResult
	err := d.Client.Call("DOM.resolveNode", opts, &result)
	return &result, err
}

type GetAttributesOpts struct {
	// Id of the node to retrieve attibutes for.
	NodeId NodeId `json:"nodeId"`
}

type GetAttributesResult struct {
	// An interleaved array of node attribute names and values.
	Attributes []string `json:"attributes"`
}

// Returns attributes for the specified node.
func (d *Domain) GetAttributes(opts *GetAttributesOpts) (*GetAttributesResult, error) {
	var result GetAttributesResult
	err := d.Client.Call("DOM.getAttributes", opts, &result)
	return &result, err
}

type CopyToOpts struct {
	// Id of the node to copy.
	NodeId NodeId `json:"nodeId"`

	// Id of the element to drop the copy into.
	TargetNodeId NodeId `json:"targetNodeId"`

	// Drop the copy before this node (if absent, the copy becomes the last child of <code>targetNodeId</code>). (optional)
	InsertBeforeNodeId NodeId `json:"insertBeforeNodeId,omitempty"`
}

type CopyToResult struct {
	// Id of the node clone.
	NodeId NodeId `json:"nodeId"`
}

// Creates a deep copy of the specified node and places it into the target container before the given anchor. (experimental)
func (d *Domain) CopyTo(opts *CopyToOpts) (*CopyToResult, error) {
	var result CopyToResult
	err := d.Client.Call("DOM.copyTo", opts, &result)
	return &result, err
}

type MoveToOpts struct {
	// Id of the node to move.
	NodeId NodeId `json:"nodeId"`

	// Id of the element to drop the moved node into.
	TargetNodeId NodeId `json:"targetNodeId"`

	// Drop node before this one (if absent, the moved node becomes the last child of <code>targetNodeId</code>). (optional)
	InsertBeforeNodeId NodeId `json:"insertBeforeNodeId,omitempty"`
}

type MoveToResult struct {
	// New id of the moved node.
	NodeId NodeId `json:"nodeId"`
}

// Moves node into the new container, places it before the given anchor.
func (d *Domain) MoveTo(opts *MoveToOpts) (*MoveToResult, error) {
	var result MoveToResult
	err := d.Client.Call("DOM.moveTo", opts, &result)
	return &result, err
}

// Undoes the last performed action. (experimental)
func (d *Domain) Undo() error {
	return d.Client.Call("DOM.undo", nil, nil)
}

// Re-does the last undone action. (experimental)
func (d *Domain) Redo() error {
	return d.Client.Call("DOM.redo", nil, nil)
}

// Marks last undoable state. (experimental)
func (d *Domain) MarkUndoableState() error {
	return d.Client.Call("DOM.markUndoableState", nil, nil)
}

type FocusOpts struct {
	// Id of the node to focus.
	NodeId NodeId `json:"nodeId"`
}

// Focuses the given element. (experimental)
func (d *Domain) Focus(opts *FocusOpts) error {
	return d.Client.Call("DOM.focus", opts, nil)
}

type SetFileInputFilesOpts struct {
	// Id of the file input node to set files for.
	NodeId NodeId `json:"nodeId"`

	// Array of file paths to set.
	Files []string `json:"files"`
}

// Sets files for the given file input element. (experimental)
func (d *Domain) SetFileInputFiles(opts *SetFileInputFilesOpts) error {
	return d.Client.Call("DOM.setFileInputFiles", opts, nil)
}

type GetBoxModelOpts struct {
	// Id of the node to get box model for.
	NodeId NodeId `json:"nodeId"`
}

type GetBoxModelResult struct {
	// Box model for the node.
	Model BoxModel `json:"model"`
}

// Returns boxes for the currently selected nodes. (experimental)
func (d *Domain) GetBoxModel(opts *GetBoxModelOpts) (*GetBoxModelResult, error) {
	var result GetBoxModelResult
	err := d.Client.Call("DOM.getBoxModel", opts, &result)
	return &result, err
}

type GetNodeForLocationOpts struct {
	// X coordinate.
	X int `json:"x"`

	// Y coordinate.
	Y int `json:"y"`
}

type GetNodeForLocationResult struct {
	// Id of the node at given coordinates.
	NodeId NodeId `json:"nodeId"`
}

// Returns node id at given location. (experimental)
func (d *Domain) GetNodeForLocation(opts *GetNodeForLocationOpts) (*GetNodeForLocationResult, error) {
	var result GetNodeForLocationResult
	err := d.Client.Call("DOM.getNodeForLocation", opts, &result)
	return &result, err
}

type GetRelayoutBoundaryOpts struct {
	// Id of the node.
	NodeId NodeId `json:"nodeId"`
}

type GetRelayoutBoundaryResult struct {
	// Relayout boundary node id for the given node.
	NodeId NodeId `json:"nodeId"`
}

// Returns the id of the nearest ancestor that is a relayout boundary. (experimental)
func (d *Domain) GetRelayoutBoundary(opts *GetRelayoutBoundaryOpts) (*GetRelayoutBoundaryResult, error) {
	var result GetRelayoutBoundaryResult
	err := d.Client.Call("DOM.getRelayoutBoundary", opts, &result)
	return &result, err
}

type GetHighlightObjectForTestOpts struct {
	// Id of the node to get highlight object for.
	NodeId NodeId `json:"nodeId"`
}

type GetHighlightObjectForTestResult struct {
	// Highlight data for the node.
	Highlight interface{} `json:"highlight"`
}

// For testing. (experimental)
func (d *Domain) GetHighlightObjectForTest(opts *GetHighlightObjectForTestOpts) (*GetHighlightObjectForTestResult, error) {
	var result GetHighlightObjectForTestResult
	err := d.Client.Call("DOM.getHighlightObjectForTest", opts, &result)
	return &result, err
}

type DocumentUpdatedEvent struct {
}

// Fired when <code>Document</code> has been totally updated. Node ids are no longer valid.
func (d *Domain) OnDocumentUpdated(listener func(*DocumentUpdatedEvent)) {
	d.Client.AddListener("DOM.documentUpdated", func(params json.RawMessage) {
		var event DocumentUpdatedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type InspectNodeRequestedEvent struct {
	// Id of the node to inspect.
	BackendNodeId BackendNodeId `json:"backendNodeId"`
}

// Fired when the node should be inspected. This happens after call to <code>setInspectMode</code>. (experimental)
func (d *Domain) OnInspectNodeRequested(listener func(*InspectNodeRequestedEvent)) {
	d.Client.AddListener("DOM.inspectNodeRequested", func(params json.RawMessage) {
		var event InspectNodeRequestedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type SetChildNodesEvent struct {
	// Parent node id to populate with children.
	ParentId NodeId `json:"parentId"`

	// Child nodes array.
	Nodes []Node `json:"nodes"`
}

// Fired when backend wants to provide client with the missing DOM structure. This happens upon most of the calls requesting node ids.
func (d *Domain) OnSetChildNodes(listener func(*SetChildNodesEvent)) {
	d.Client.AddListener("DOM.setChildNodes", func(params json.RawMessage) {
		var event SetChildNodesEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type AttributeModifiedEvent struct {
	// Id of the node that has changed.
	NodeId NodeId `json:"nodeId"`

	// Attribute name.
	Name string `json:"name"`

	// Attribute value.
	Value string `json:"value"`
}

// Fired when <code>Element</code>'s attribute is modified.
func (d *Domain) OnAttributeModified(listener func(*AttributeModifiedEvent)) {
	d.Client.AddListener("DOM.attributeModified", func(params json.RawMessage) {
		var event AttributeModifiedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type AttributeRemovedEvent struct {
	// Id of the node that has changed.
	NodeId NodeId `json:"nodeId"`

	// A ttribute name.
	Name string `json:"name"`
}

// Fired when <code>Element</code>'s attribute is removed.
func (d *Domain) OnAttributeRemoved(listener func(*AttributeRemovedEvent)) {
	d.Client.AddListener("DOM.attributeRemoved", func(params json.RawMessage) {
		var event AttributeRemovedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type InlineStyleInvalidatedEvent struct {
	// Ids of the nodes for which the inline styles have been invalidated.
	NodeIds []NodeId `json:"nodeIds"`
}

// Fired when <code>Element</code>'s inline style is modified via a CSS property modification. (experimental)
func (d *Domain) OnInlineStyleInvalidated(listener func(*InlineStyleInvalidatedEvent)) {
	d.Client.AddListener("DOM.inlineStyleInvalidated", func(params json.RawMessage) {
		var event InlineStyleInvalidatedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type CharacterDataModifiedEvent struct {
	// Id of the node that has changed.
	NodeId NodeId `json:"nodeId"`

	// New text value.
	CharacterData string `json:"characterData"`
}

// Mirrors <code>DOMCharacterDataModified</code> event.
func (d *Domain) OnCharacterDataModified(listener func(*CharacterDataModifiedEvent)) {
	d.Client.AddListener("DOM.characterDataModified", func(params json.RawMessage) {
		var event CharacterDataModifiedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type ChildNodeCountUpdatedEvent struct {
	// Id of the node that has changed.
	NodeId NodeId `json:"nodeId"`

	// New node count.
	ChildNodeCount int `json:"childNodeCount"`
}

// Fired when <code>Container</code>'s child node count has changed.
func (d *Domain) OnChildNodeCountUpdated(listener func(*ChildNodeCountUpdatedEvent)) {
	d.Client.AddListener("DOM.childNodeCountUpdated", func(params json.RawMessage) {
		var event ChildNodeCountUpdatedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type ChildNodeInsertedEvent struct {
	// Id of the node that has changed.
	ParentNodeId NodeId `json:"parentNodeId"`

	// If of the previous siblint.
	PreviousNodeId NodeId `json:"previousNodeId"`

	// Inserted node data.
	Node Node `json:"node"`
}

// Mirrors <code>DOMNodeInserted</code> event.
func (d *Domain) OnChildNodeInserted(listener func(*ChildNodeInsertedEvent)) {
	d.Client.AddListener("DOM.childNodeInserted", func(params json.RawMessage) {
		var event ChildNodeInsertedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type ChildNodeRemovedEvent struct {
	// Parent id.
	ParentNodeId NodeId `json:"parentNodeId"`

	// Id of the node that has been removed.
	NodeId NodeId `json:"nodeId"`
}

// Mirrors <code>DOMNodeRemoved</code> event.
func (d *Domain) OnChildNodeRemoved(listener func(*ChildNodeRemovedEvent)) {
	d.Client.AddListener("DOM.childNodeRemoved", func(params json.RawMessage) {
		var event ChildNodeRemovedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type ShadowRootPushedEvent struct {
	// Host element id.
	HostId NodeId `json:"hostId"`

	// Shadow root.
	Root Node `json:"root"`
}

// Called when shadow root is pushed into the element. (experimental)
func (d *Domain) OnShadowRootPushed(listener func(*ShadowRootPushedEvent)) {
	d.Client.AddListener("DOM.shadowRootPushed", func(params json.RawMessage) {
		var event ShadowRootPushedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type ShadowRootPoppedEvent struct {
	// Host element id.
	HostId NodeId `json:"hostId"`

	// Shadow root id.
	RootId NodeId `json:"rootId"`
}

// Called when shadow root is popped from the element. (experimental)
func (d *Domain) OnShadowRootPopped(listener func(*ShadowRootPoppedEvent)) {
	d.Client.AddListener("DOM.shadowRootPopped", func(params json.RawMessage) {
		var event ShadowRootPoppedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type PseudoElementAddedEvent struct {
	// Pseudo element's parent element id.
	ParentId NodeId `json:"parentId"`

	// The added pseudo element.
	PseudoElement Node `json:"pseudoElement"`
}

// Called when a pseudo element is added to an element. (experimental)
func (d *Domain) OnPseudoElementAdded(listener func(*PseudoElementAddedEvent)) {
	d.Client.AddListener("DOM.pseudoElementAdded", func(params json.RawMessage) {
		var event PseudoElementAddedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type PseudoElementRemovedEvent struct {
	// Pseudo element's parent element id.
	ParentId NodeId `json:"parentId"`

	// The removed pseudo element id.
	PseudoElementId NodeId `json:"pseudoElementId"`
}

// Called when a pseudo element is removed from an element. (experimental)
func (d *Domain) OnPseudoElementRemoved(listener func(*PseudoElementRemovedEvent)) {
	d.Client.AddListener("DOM.pseudoElementRemoved", func(params json.RawMessage) {
		var event PseudoElementRemovedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type DistributedNodesUpdatedEvent struct {
	// Insertion point where distrubuted nodes were updated.
	InsertionPointId NodeId `json:"insertionPointId"`

	// Distributed nodes for given insertion point.
	DistributedNodes []BackendNode `json:"distributedNodes"`
}

// Called when distrubution is changed. (experimental)
func (d *Domain) OnDistributedNodesUpdated(listener func(*DistributedNodesUpdatedEvent)) {
	d.Client.AddListener("DOM.distributedNodesUpdated", func(params json.RawMessage) {
		var event DistributedNodesUpdatedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type NodeHighlightRequestedEvent struct {
	NodeId NodeId `json:"nodeId"`
}

// (experimental)
func (d *Domain) OnNodeHighlightRequested(listener func(*NodeHighlightRequestedEvent)) {
	d.Client.AddListener("DOM.nodeHighlightRequested", func(params json.RawMessage) {
		var event NodeHighlightRequestedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}
