// This domain exposes DOM read/write operations. Each DOM Node is represented with its mirror object that has an <code>id</code>. This <code>id</code> can be used to get additional information on the Node, resolve it into the JavaScript object wrapper, etc. It is important that client receives DOM events only for the nodes that are known to the client. Backend keeps track of the nodes that were sent to the client and never sends the same node twice. It is client's responsibility to collect information about the nodes that were sent to the client.<p>Note that <code>iframe</code> owner elements will return corresponding document elements as their child nodes.</p>
package dom

import (
	"github.com/neelance/cdp-go/rpc"
)

// This domain exposes DOM read/write operations. Each DOM Node is represented with its mirror object that has an <code>id</code>. This <code>id</code> can be used to get additional information on the Node, resolve it into the JavaScript object wrapper, etc. It is important that client receives DOM events only for the nodes that are known to the client. Backend keeps track of the nodes that were sent to the client and never sends the same node twice. It is client's responsibility to collect information about the nodes that were sent to the client.<p>Note that <code>iframe</code> owner elements will return corresponding document elements as their child nodes.</p>
type Domain struct {
	Client *rpc.Client
}

// Unique DOM node identifier.

type NodeId int

// Unique DOM node identifier used to reference a node that may not have been pushed to the front-end. (experimental)

type BackendNodeId int

// Backend node with a friendly name. (experimental)

type BackendNode struct {
	// <code>Node</code>'s nodeType.
	NodeType int `json:"nodeType"`

	// <code>Node</code>'s nodeName.
	NodeName string `json:"nodeName"`

	BackendNodeId BackendNodeId `json:"backendNodeId"`
}

// Pseudo element type.

type PseudoType string

// Shadow root type.

type ShadowRootType string

// DOM interaction is implemented in terms of mirror objects that represent the actual DOM nodes. DOMNode is a base node mirror type.

type Node struct {
	// Node identifier that is passed into the rest of the DOM messages as the <code>nodeId</code>. Backend will only push node with given <code>id</code> once. It is aware of all requested nodes and will only fire DOM events for nodes known to the client.
	NodeId NodeId `json:"nodeId"`

	// The id of the parent node if any. (optional, experimental)
	ParentId NodeId `json:"parentId,omitempty"`

	// The BackendNodeId for this node.
	BackendNodeId BackendNodeId `json:"backendNodeId"`

	// <code>Node</code>'s nodeType.
	NodeType int `json:"nodeType"`

	// <code>Node</code>'s nodeName.
	NodeName string `json:"nodeName"`

	// <code>Node</code>'s localName.
	LocalName string `json:"localName"`

	// <code>Node</code>'s nodeValue.
	NodeValue string `json:"nodeValue"`

	// Child count for <code>Container</code> nodes. (optional)
	ChildNodeCount int `json:"childNodeCount,omitempty"`

	// Child nodes of this node when requested with children. (optional)
	Children []*Node `json:"children,omitempty"`

	// Attributes of the <code>Element</code> node in the form of flat array <code>[name1, value1, name2, value2]</code>. (optional)
	Attributes []string `json:"attributes,omitempty"`

	// Document URL that <code>Document</code> or <code>FrameOwner</code> node points to. (optional)
	DocumentURL string `json:"documentURL,omitempty"`

	// Base URL that <code>Document</code> or <code>FrameOwner</code> node uses for URL completion. (optional, experimental)
	BaseURL string `json:"baseURL,omitempty"`

	// <code>DocumentType</code>'s publicId. (optional)
	PublicId string `json:"publicId,omitempty"`

	// <code>DocumentType</code>'s systemId. (optional)
	SystemId string `json:"systemId,omitempty"`

	// <code>DocumentType</code>'s internalSubset. (optional)
	InternalSubset string `json:"internalSubset,omitempty"`

	// <code>Document</code>'s XML version in case of XML documents. (optional)
	XmlVersion string `json:"xmlVersion,omitempty"`

	// <code>Attr</code>'s name. (optional)
	Name string `json:"name,omitempty"`

	// <code>Attr</code>'s value. (optional)
	Value string `json:"value,omitempty"`

	// Pseudo element type for this node. (optional)
	PseudoType PseudoType `json:"pseudoType,omitempty"`

	// Shadow root type. (optional)
	ShadowRootType ShadowRootType `json:"shadowRootType,omitempty"`

	// Frame ID for frame owner elements. (optional, experimental)
	FrameId interface{} `json:"frameId,omitempty"`

	// Content document for frame owner elements. (optional)
	ContentDocument *Node `json:"contentDocument,omitempty"`

	// Shadow root list for given element host. (optional, experimental)
	ShadowRoots []*Node `json:"shadowRoots,omitempty"`

	// Content document fragment for template elements. (optional, experimental)
	TemplateContent *Node `json:"templateContent,omitempty"`

	// Pseudo elements associated with this node. (optional, experimental)
	PseudoElements []*Node `json:"pseudoElements,omitempty"`

	// Import document for the HTMLImport links. (optional)
	ImportedDocument *Node `json:"importedDocument,omitempty"`

	// Distributed nodes for given insertion point. (optional, experimental)
	DistributedNodes []*BackendNode `json:"distributedNodes,omitempty"`

	// Whether the node is SVG. (optional, experimental)
	IsSVG bool `json:"isSVG,omitempty"`
}

// A structure holding an RGBA color.

type RGBA struct {
	// The red component, in the [0-255] range.
	R int `json:"r"`

	// The green component, in the [0-255] range.
	G int `json:"g"`

	// The blue component, in the [0-255] range.
	B int `json:"b"`

	// The alpha component, in the [0-1] range (default: 1). (optional)
	A float64 `json:"a,omitempty"`
}

// An array of quad vertices, x immediately followed by y for each point, points clock-wise. (experimental)

type Quad []float64

// Box model. (experimental)

type BoxModel struct {
	// Content box
	Content Quad `json:"content"`

	// Padding box
	Padding Quad `json:"padding"`

	// Border box
	Border Quad `json:"border"`

	// Margin box
	Margin Quad `json:"margin"`

	// Node width
	Width int `json:"width"`

	// Node height
	Height int `json:"height"`

	// Shape outside coordinates (optional)
	ShapeOutside *ShapeOutsideInfo `json:"shapeOutside,omitempty"`
}

// CSS Shape Outside details. (experimental)

type ShapeOutsideInfo struct {
	// Shape bounds
	Bounds Quad `json:"bounds"`

	// Shape coordinate details
	Shape []interface{} `json:"shape"`

	// Margin shape bounds
	MarginShape []interface{} `json:"marginShape"`
}

// Rectangle. (experimental)

type Rect struct {
	// X coordinate
	X float64 `json:"x"`

	// Y coordinate
	Y float64 `json:"y"`

	// Rectangle width
	Width float64 `json:"width"`

	// Rectangle height
	Height float64 `json:"height"`
}

// Enables DOM agent for the given page.
type EnableRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) Enable() *EnableRequest {
	return &EnableRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Enables DOM agent for the given page.
func (r *EnableRequest) Do() error {
	return r.client.Call("DOM.enable", r.opts, nil)
}

// Disables DOM agent for the given page.
type DisableRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) Disable() *DisableRequest {
	return &DisableRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Disables DOM agent for the given page.
func (r *DisableRequest) Do() error {
	return r.client.Call("DOM.disable", r.opts, nil)
}

// Returns the root DOM node (and optionally the subtree) to the caller.
type GetDocumentRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) GetDocument() *GetDocumentRequest {
	return &GetDocumentRequest{opts: make(map[string]interface{}), client: d.Client}
}

// The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0. (optional, experimental)
func (r *GetDocumentRequest) Depth(v int) *GetDocumentRequest {
	r.opts["depth"] = v
	return r
}

// Whether or not iframes and shadow roots should be traversed when returning the subtree (default is false). (optional, experimental)
func (r *GetDocumentRequest) Pierce(v bool) *GetDocumentRequest {
	r.opts["pierce"] = v
	return r
}

type GetDocumentResult struct {
	// Resulting node.
	Root *Node `json:"root"`
}

func (r *GetDocumentRequest) Do() (*GetDocumentResult, error) {
	var result GetDocumentResult
	err := r.client.Call("DOM.getDocument", r.opts, &result)
	return &result, err
}

// Returns the root DOM node (and optionally the subtree) to the caller.
type GetFlattenedDocumentRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) GetFlattenedDocument() *GetFlattenedDocumentRequest {
	return &GetFlattenedDocumentRequest{opts: make(map[string]interface{}), client: d.Client}
}

// The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0. (optional, experimental)
func (r *GetFlattenedDocumentRequest) Depth(v int) *GetFlattenedDocumentRequest {
	r.opts["depth"] = v
	return r
}

// Whether or not iframes and shadow roots should be traversed when returning the subtree (default is false). (optional, experimental)
func (r *GetFlattenedDocumentRequest) Pierce(v bool) *GetFlattenedDocumentRequest {
	r.opts["pierce"] = v
	return r
}

type GetFlattenedDocumentResult struct {
	// Resulting node.
	Nodes []*Node `json:"nodes"`
}

func (r *GetFlattenedDocumentRequest) Do() (*GetFlattenedDocumentResult, error) {
	var result GetFlattenedDocumentResult
	err := r.client.Call("DOM.getFlattenedDocument", r.opts, &result)
	return &result, err
}

// Collects class names for the node with given id and all of it's child nodes. (experimental)
type CollectClassNamesFromSubtreeRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) CollectClassNamesFromSubtree() *CollectClassNamesFromSubtreeRequest {
	return &CollectClassNamesFromSubtreeRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Id of the node to collect class names.
func (r *CollectClassNamesFromSubtreeRequest) NodeId(v NodeId) *CollectClassNamesFromSubtreeRequest {
	r.opts["nodeId"] = v
	return r
}

type CollectClassNamesFromSubtreeResult struct {
	// Class name list.
	ClassNames []string `json:"classNames"`
}

func (r *CollectClassNamesFromSubtreeRequest) Do() (*CollectClassNamesFromSubtreeResult, error) {
	var result CollectClassNamesFromSubtreeResult
	err := r.client.Call("DOM.collectClassNamesFromSubtree", r.opts, &result)
	return &result, err
}

// Requests that children of the node with given id are returned to the caller in form of <code>setChildNodes</code> events where not only immediate children are retrieved, but all children down to the specified depth.
type RequestChildNodesRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) RequestChildNodes() *RequestChildNodesRequest {
	return &RequestChildNodesRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Id of the node to get children for.
func (r *RequestChildNodesRequest) NodeId(v NodeId) *RequestChildNodesRequest {
	r.opts["nodeId"] = v
	return r
}

// The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0. (optional, experimental)
func (r *RequestChildNodesRequest) Depth(v int) *RequestChildNodesRequest {
	r.opts["depth"] = v
	return r
}

// Whether or not iframes and shadow roots should be traversed when returning the sub-tree (default is false). (optional, experimental)
func (r *RequestChildNodesRequest) Pierce(v bool) *RequestChildNodesRequest {
	r.opts["pierce"] = v
	return r
}

// Requests that children of the node with given id are returned to the caller in form of <code>setChildNodes</code> events where not only immediate children are retrieved, but all children down to the specified depth.
func (r *RequestChildNodesRequest) Do() error {
	return r.client.Call("DOM.requestChildNodes", r.opts, nil)
}

// Executes <code>querySelector</code> on a given node.
type QuerySelectorRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) QuerySelector() *QuerySelectorRequest {
	return &QuerySelectorRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Id of the node to query upon.
func (r *QuerySelectorRequest) NodeId(v NodeId) *QuerySelectorRequest {
	r.opts["nodeId"] = v
	return r
}

// Selector string.
func (r *QuerySelectorRequest) Selector(v string) *QuerySelectorRequest {
	r.opts["selector"] = v
	return r
}

type QuerySelectorResult struct {
	// Query selector result.
	NodeId NodeId `json:"nodeId"`
}

func (r *QuerySelectorRequest) Do() (*QuerySelectorResult, error) {
	var result QuerySelectorResult
	err := r.client.Call("DOM.querySelector", r.opts, &result)
	return &result, err
}

// Executes <code>querySelectorAll</code> on a given node.
type QuerySelectorAllRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) QuerySelectorAll() *QuerySelectorAllRequest {
	return &QuerySelectorAllRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Id of the node to query upon.
func (r *QuerySelectorAllRequest) NodeId(v NodeId) *QuerySelectorAllRequest {
	r.opts["nodeId"] = v
	return r
}

// Selector string.
func (r *QuerySelectorAllRequest) Selector(v string) *QuerySelectorAllRequest {
	r.opts["selector"] = v
	return r
}

type QuerySelectorAllResult struct {
	// Query selector result.
	NodeIds []NodeId `json:"nodeIds"`
}

func (r *QuerySelectorAllRequest) Do() (*QuerySelectorAllResult, error) {
	var result QuerySelectorAllResult
	err := r.client.Call("DOM.querySelectorAll", r.opts, &result)
	return &result, err
}

// Sets node name for a node with given id.
type SetNodeNameRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) SetNodeName() *SetNodeNameRequest {
	return &SetNodeNameRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Id of the node to set name for.
func (r *SetNodeNameRequest) NodeId(v NodeId) *SetNodeNameRequest {
	r.opts["nodeId"] = v
	return r
}

// New node's name.
func (r *SetNodeNameRequest) Name(v string) *SetNodeNameRequest {
	r.opts["name"] = v
	return r
}

type SetNodeNameResult struct {
	// New node's id.
	NodeId NodeId `json:"nodeId"`
}

func (r *SetNodeNameRequest) Do() (*SetNodeNameResult, error) {
	var result SetNodeNameResult
	err := r.client.Call("DOM.setNodeName", r.opts, &result)
	return &result, err
}

// Sets node value for a node with given id.
type SetNodeValueRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) SetNodeValue() *SetNodeValueRequest {
	return &SetNodeValueRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Id of the node to set value for.
func (r *SetNodeValueRequest) NodeId(v NodeId) *SetNodeValueRequest {
	r.opts["nodeId"] = v
	return r
}

// New node's value.
func (r *SetNodeValueRequest) Value(v string) *SetNodeValueRequest {
	r.opts["value"] = v
	return r
}

// Sets node value for a node with given id.
func (r *SetNodeValueRequest) Do() error {
	return r.client.Call("DOM.setNodeValue", r.opts, nil)
}

// Removes node with given id.
type RemoveNodeRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) RemoveNode() *RemoveNodeRequest {
	return &RemoveNodeRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Id of the node to remove.
func (r *RemoveNodeRequest) NodeId(v NodeId) *RemoveNodeRequest {
	r.opts["nodeId"] = v
	return r
}

// Removes node with given id.
func (r *RemoveNodeRequest) Do() error {
	return r.client.Call("DOM.removeNode", r.opts, nil)
}

// Sets attribute for an element with given id.
type SetAttributeValueRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) SetAttributeValue() *SetAttributeValueRequest {
	return &SetAttributeValueRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Id of the element to set attribute for.
func (r *SetAttributeValueRequest) NodeId(v NodeId) *SetAttributeValueRequest {
	r.opts["nodeId"] = v
	return r
}

// Attribute name.
func (r *SetAttributeValueRequest) Name(v string) *SetAttributeValueRequest {
	r.opts["name"] = v
	return r
}

// Attribute value.
func (r *SetAttributeValueRequest) Value(v string) *SetAttributeValueRequest {
	r.opts["value"] = v
	return r
}

// Sets attribute for an element with given id.
func (r *SetAttributeValueRequest) Do() error {
	return r.client.Call("DOM.setAttributeValue", r.opts, nil)
}

// Sets attributes on element with given id. This method is useful when user edits some existing attribute value and types in several attribute name/value pairs.
type SetAttributesAsTextRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) SetAttributesAsText() *SetAttributesAsTextRequest {
	return &SetAttributesAsTextRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Id of the element to set attributes for.
func (r *SetAttributesAsTextRequest) NodeId(v NodeId) *SetAttributesAsTextRequest {
	r.opts["nodeId"] = v
	return r
}

// Text with a number of attributes. Will parse this text using HTML parser.
func (r *SetAttributesAsTextRequest) Text(v string) *SetAttributesAsTextRequest {
	r.opts["text"] = v
	return r
}

// Attribute name to replace with new attributes derived from text in case text parsed successfully. (optional)
func (r *SetAttributesAsTextRequest) Name(v string) *SetAttributesAsTextRequest {
	r.opts["name"] = v
	return r
}

// Sets attributes on element with given id. This method is useful when user edits some existing attribute value and types in several attribute name/value pairs.
func (r *SetAttributesAsTextRequest) Do() error {
	return r.client.Call("DOM.setAttributesAsText", r.opts, nil)
}

// Removes attribute with given name from an element with given id.
type RemoveAttributeRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) RemoveAttribute() *RemoveAttributeRequest {
	return &RemoveAttributeRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Id of the element to remove attribute from.
func (r *RemoveAttributeRequest) NodeId(v NodeId) *RemoveAttributeRequest {
	r.opts["nodeId"] = v
	return r
}

// Name of the attribute to remove.
func (r *RemoveAttributeRequest) Name(v string) *RemoveAttributeRequest {
	r.opts["name"] = v
	return r
}

// Removes attribute with given name from an element with given id.
func (r *RemoveAttributeRequest) Do() error {
	return r.client.Call("DOM.removeAttribute", r.opts, nil)
}

// Returns node's HTML markup.
type GetOuterHTMLRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) GetOuterHTML() *GetOuterHTMLRequest {
	return &GetOuterHTMLRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Id of the node to get markup for.
func (r *GetOuterHTMLRequest) NodeId(v NodeId) *GetOuterHTMLRequest {
	r.opts["nodeId"] = v
	return r
}

type GetOuterHTMLResult struct {
	// Outer HTML markup.
	OuterHTML string `json:"outerHTML"`
}

func (r *GetOuterHTMLRequest) Do() (*GetOuterHTMLResult, error) {
	var result GetOuterHTMLResult
	err := r.client.Call("DOM.getOuterHTML", r.opts, &result)
	return &result, err
}

// Sets node HTML markup, returns new node id.
type SetOuterHTMLRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) SetOuterHTML() *SetOuterHTMLRequest {
	return &SetOuterHTMLRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Id of the node to set markup for.
func (r *SetOuterHTMLRequest) NodeId(v NodeId) *SetOuterHTMLRequest {
	r.opts["nodeId"] = v
	return r
}

// Outer HTML markup to set.
func (r *SetOuterHTMLRequest) OuterHTML(v string) *SetOuterHTMLRequest {
	r.opts["outerHTML"] = v
	return r
}

// Sets node HTML markup, returns new node id.
func (r *SetOuterHTMLRequest) Do() error {
	return r.client.Call("DOM.setOuterHTML", r.opts, nil)
}

// Searches for a given string in the DOM tree. Use <code>getSearchResults</code> to access search results or <code>cancelSearch</code> to end this search session. (experimental)
type PerformSearchRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) PerformSearch() *PerformSearchRequest {
	return &PerformSearchRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Plain text or query selector or XPath search query.
func (r *PerformSearchRequest) Query(v string) *PerformSearchRequest {
	r.opts["query"] = v
	return r
}

// True to search in user agent shadow DOM. (optional, experimental)
func (r *PerformSearchRequest) IncludeUserAgentShadowDOM(v bool) *PerformSearchRequest {
	r.opts["includeUserAgentShadowDOM"] = v
	return r
}

type PerformSearchResult struct {
	// Unique search session identifier.
	SearchId string `json:"searchId"`

	// Number of search results.
	ResultCount int `json:"resultCount"`
}

func (r *PerformSearchRequest) Do() (*PerformSearchResult, error) {
	var result PerformSearchResult
	err := r.client.Call("DOM.performSearch", r.opts, &result)
	return &result, err
}

// Returns search results from given <code>fromIndex</code> to given <code>toIndex</code> from the sarch with the given identifier. (experimental)
type GetSearchResultsRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) GetSearchResults() *GetSearchResultsRequest {
	return &GetSearchResultsRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Unique search session identifier.
func (r *GetSearchResultsRequest) SearchId(v string) *GetSearchResultsRequest {
	r.opts["searchId"] = v
	return r
}

// Start index of the search result to be returned.
func (r *GetSearchResultsRequest) FromIndex(v int) *GetSearchResultsRequest {
	r.opts["fromIndex"] = v
	return r
}

// End index of the search result to be returned.
func (r *GetSearchResultsRequest) ToIndex(v int) *GetSearchResultsRequest {
	r.opts["toIndex"] = v
	return r
}

type GetSearchResultsResult struct {
	// Ids of the search result nodes.
	NodeIds []NodeId `json:"nodeIds"`
}

func (r *GetSearchResultsRequest) Do() (*GetSearchResultsResult, error) {
	var result GetSearchResultsResult
	err := r.client.Call("DOM.getSearchResults", r.opts, &result)
	return &result, err
}

// Discards search results from the session with the given id. <code>getSearchResults</code> should no longer be called for that search. (experimental)
type DiscardSearchResultsRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) DiscardSearchResults() *DiscardSearchResultsRequest {
	return &DiscardSearchResultsRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Unique search session identifier.
func (r *DiscardSearchResultsRequest) SearchId(v string) *DiscardSearchResultsRequest {
	r.opts["searchId"] = v
	return r
}

// Discards search results from the session with the given id. <code>getSearchResults</code> should no longer be called for that search. (experimental)
func (r *DiscardSearchResultsRequest) Do() error {
	return r.client.Call("DOM.discardSearchResults", r.opts, nil)
}

// Requests that the node is sent to the caller given the JavaScript node object reference. All nodes that form the path from the node to the root are also sent to the client as a series of <code>setChildNodes</code> notifications.
type RequestNodeRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) RequestNode() *RequestNodeRequest {
	return &RequestNodeRequest{opts: make(map[string]interface{}), client: d.Client}
}

// JavaScript object id to convert into node.
func (r *RequestNodeRequest) ObjectId(v interface{}) *RequestNodeRequest {
	r.opts["objectId"] = v
	return r
}

type RequestNodeResult struct {
	// Node id for given object.
	NodeId NodeId `json:"nodeId"`
}

func (r *RequestNodeRequest) Do() (*RequestNodeResult, error) {
	var result RequestNodeResult
	err := r.client.Call("DOM.requestNode", r.opts, &result)
	return &result, err
}

// Highlights given rectangle.
type HighlightRectRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) HighlightRect() *HighlightRectRequest {
	return &HighlightRectRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Highlights given rectangle.
func (r *HighlightRectRequest) Do() error {
	return r.client.Call("DOM.highlightRect", r.opts, nil)
}

// Highlights DOM node.
type HighlightNodeRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) HighlightNode() *HighlightNodeRequest {
	return &HighlightNodeRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Highlights DOM node.
func (r *HighlightNodeRequest) Do() error {
	return r.client.Call("DOM.highlightNode", r.opts, nil)
}

// Hides any highlight.
type HideHighlightRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) HideHighlight() *HideHighlightRequest {
	return &HideHighlightRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Hides any highlight.
func (r *HideHighlightRequest) Do() error {
	return r.client.Call("DOM.hideHighlight", r.opts, nil)
}

// Requests that the node is sent to the caller given its path. // FIXME, use XPath (experimental)
type PushNodeByPathToFrontendRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) PushNodeByPathToFrontend() *PushNodeByPathToFrontendRequest {
	return &PushNodeByPathToFrontendRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Path to node in the proprietary format.
func (r *PushNodeByPathToFrontendRequest) Path(v string) *PushNodeByPathToFrontendRequest {
	r.opts["path"] = v
	return r
}

type PushNodeByPathToFrontendResult struct {
	// Id of the node for given path.
	NodeId NodeId `json:"nodeId"`
}

func (r *PushNodeByPathToFrontendRequest) Do() (*PushNodeByPathToFrontendResult, error) {
	var result PushNodeByPathToFrontendResult
	err := r.client.Call("DOM.pushNodeByPathToFrontend", r.opts, &result)
	return &result, err
}

// Requests that a batch of nodes is sent to the caller given their backend node ids. (experimental)
type PushNodesByBackendIdsToFrontendRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) PushNodesByBackendIdsToFrontend() *PushNodesByBackendIdsToFrontendRequest {
	return &PushNodesByBackendIdsToFrontendRequest{opts: make(map[string]interface{}), client: d.Client}
}

// The array of backend node ids.
func (r *PushNodesByBackendIdsToFrontendRequest) BackendNodeIds(v []BackendNodeId) *PushNodesByBackendIdsToFrontendRequest {
	r.opts["backendNodeIds"] = v
	return r
}

type PushNodesByBackendIdsToFrontendResult struct {
	// The array of ids of pushed nodes that correspond to the backend ids specified in backendNodeIds.
	NodeIds []NodeId `json:"nodeIds"`
}

func (r *PushNodesByBackendIdsToFrontendRequest) Do() (*PushNodesByBackendIdsToFrontendResult, error) {
	var result PushNodesByBackendIdsToFrontendResult
	err := r.client.Call("DOM.pushNodesByBackendIdsToFrontend", r.opts, &result)
	return &result, err
}

// Enables console to refer to the node with given id via $x (see Command Line API for more details $x functions). (experimental)
type SetInspectedNodeRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) SetInspectedNode() *SetInspectedNodeRequest {
	return &SetInspectedNodeRequest{opts: make(map[string]interface{}), client: d.Client}
}

// DOM node id to be accessible by means of $x command line API.
func (r *SetInspectedNodeRequest) NodeId(v NodeId) *SetInspectedNodeRequest {
	r.opts["nodeId"] = v
	return r
}

// Enables console to refer to the node with given id via $x (see Command Line API for more details $x functions). (experimental)
func (r *SetInspectedNodeRequest) Do() error {
	return r.client.Call("DOM.setInspectedNode", r.opts, nil)
}

// Resolves JavaScript node object for given node id.
type ResolveNodeRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) ResolveNode() *ResolveNodeRequest {
	return &ResolveNodeRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Id of the node to resolve.
func (r *ResolveNodeRequest) NodeId(v NodeId) *ResolveNodeRequest {
	r.opts["nodeId"] = v
	return r
}

// Symbolic group name that can be used to release multiple objects. (optional)
func (r *ResolveNodeRequest) ObjectGroup(v string) *ResolveNodeRequest {
	r.opts["objectGroup"] = v
	return r
}

type ResolveNodeResult struct {
	// JavaScript object wrapper for given node.
	Object interface{} `json:"object"`
}

func (r *ResolveNodeRequest) Do() (*ResolveNodeResult, error) {
	var result ResolveNodeResult
	err := r.client.Call("DOM.resolveNode", r.opts, &result)
	return &result, err
}

// Returns attributes for the specified node.
type GetAttributesRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) GetAttributes() *GetAttributesRequest {
	return &GetAttributesRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Id of the node to retrieve attibutes for.
func (r *GetAttributesRequest) NodeId(v NodeId) *GetAttributesRequest {
	r.opts["nodeId"] = v
	return r
}

type GetAttributesResult struct {
	// An interleaved array of node attribute names and values.
	Attributes []string `json:"attributes"`
}

func (r *GetAttributesRequest) Do() (*GetAttributesResult, error) {
	var result GetAttributesResult
	err := r.client.Call("DOM.getAttributes", r.opts, &result)
	return &result, err
}

// Creates a deep copy of the specified node and places it into the target container before the given anchor. (experimental)
type CopyToRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) CopyTo() *CopyToRequest {
	return &CopyToRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Id of the node to copy.
func (r *CopyToRequest) NodeId(v NodeId) *CopyToRequest {
	r.opts["nodeId"] = v
	return r
}

// Id of the element to drop the copy into.
func (r *CopyToRequest) TargetNodeId(v NodeId) *CopyToRequest {
	r.opts["targetNodeId"] = v
	return r
}

// Drop the copy before this node (if absent, the copy becomes the last child of <code>targetNodeId</code>). (optional)
func (r *CopyToRequest) InsertBeforeNodeId(v NodeId) *CopyToRequest {
	r.opts["insertBeforeNodeId"] = v
	return r
}

type CopyToResult struct {
	// Id of the node clone.
	NodeId NodeId `json:"nodeId"`
}

func (r *CopyToRequest) Do() (*CopyToResult, error) {
	var result CopyToResult
	err := r.client.Call("DOM.copyTo", r.opts, &result)
	return &result, err
}

// Moves node into the new container, places it before the given anchor.
type MoveToRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) MoveTo() *MoveToRequest {
	return &MoveToRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Id of the node to move.
func (r *MoveToRequest) NodeId(v NodeId) *MoveToRequest {
	r.opts["nodeId"] = v
	return r
}

// Id of the element to drop the moved node into.
func (r *MoveToRequest) TargetNodeId(v NodeId) *MoveToRequest {
	r.opts["targetNodeId"] = v
	return r
}

// Drop node before this one (if absent, the moved node becomes the last child of <code>targetNodeId</code>). (optional)
func (r *MoveToRequest) InsertBeforeNodeId(v NodeId) *MoveToRequest {
	r.opts["insertBeforeNodeId"] = v
	return r
}

type MoveToResult struct {
	// New id of the moved node.
	NodeId NodeId `json:"nodeId"`
}

func (r *MoveToRequest) Do() (*MoveToResult, error) {
	var result MoveToResult
	err := r.client.Call("DOM.moveTo", r.opts, &result)
	return &result, err
}

// Undoes the last performed action. (experimental)
type UndoRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) Undo() *UndoRequest {
	return &UndoRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Undoes the last performed action. (experimental)
func (r *UndoRequest) Do() error {
	return r.client.Call("DOM.undo", r.opts, nil)
}

// Re-does the last undone action. (experimental)
type RedoRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) Redo() *RedoRequest {
	return &RedoRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Re-does the last undone action. (experimental)
func (r *RedoRequest) Do() error {
	return r.client.Call("DOM.redo", r.opts, nil)
}

// Marks last undoable state. (experimental)
type MarkUndoableStateRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) MarkUndoableState() *MarkUndoableStateRequest {
	return &MarkUndoableStateRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Marks last undoable state. (experimental)
func (r *MarkUndoableStateRequest) Do() error {
	return r.client.Call("DOM.markUndoableState", r.opts, nil)
}

// Focuses the given element. (experimental)
type FocusRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) Focus() *FocusRequest {
	return &FocusRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Id of the node to focus.
func (r *FocusRequest) NodeId(v NodeId) *FocusRequest {
	r.opts["nodeId"] = v
	return r
}

// Focuses the given element. (experimental)
func (r *FocusRequest) Do() error {
	return r.client.Call("DOM.focus", r.opts, nil)
}

// Sets files for the given file input element. (experimental)
type SetFileInputFilesRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) SetFileInputFiles() *SetFileInputFilesRequest {
	return &SetFileInputFilesRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Id of the file input node to set files for.
func (r *SetFileInputFilesRequest) NodeId(v NodeId) *SetFileInputFilesRequest {
	r.opts["nodeId"] = v
	return r
}

// Array of file paths to set.
func (r *SetFileInputFilesRequest) Files(v []string) *SetFileInputFilesRequest {
	r.opts["files"] = v
	return r
}

// Sets files for the given file input element. (experimental)
func (r *SetFileInputFilesRequest) Do() error {
	return r.client.Call("DOM.setFileInputFiles", r.opts, nil)
}

// Returns boxes for the currently selected nodes. (experimental)
type GetBoxModelRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) GetBoxModel() *GetBoxModelRequest {
	return &GetBoxModelRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Id of the node to get box model for.
func (r *GetBoxModelRequest) NodeId(v NodeId) *GetBoxModelRequest {
	r.opts["nodeId"] = v
	return r
}

type GetBoxModelResult struct {
	// Box model for the node.
	Model *BoxModel `json:"model"`
}

func (r *GetBoxModelRequest) Do() (*GetBoxModelResult, error) {
	var result GetBoxModelResult
	err := r.client.Call("DOM.getBoxModel", r.opts, &result)
	return &result, err
}

// Returns node id at given location. (experimental)
type GetNodeForLocationRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) GetNodeForLocation() *GetNodeForLocationRequest {
	return &GetNodeForLocationRequest{opts: make(map[string]interface{}), client: d.Client}
}

// X coordinate.
func (r *GetNodeForLocationRequest) X(v int) *GetNodeForLocationRequest {
	r.opts["x"] = v
	return r
}

// Y coordinate.
func (r *GetNodeForLocationRequest) Y(v int) *GetNodeForLocationRequest {
	r.opts["y"] = v
	return r
}

// False to skip to the nearest non-UA shadow root ancestor (default: false). (optional)
func (r *GetNodeForLocationRequest) IncludeUserAgentShadowDOM(v bool) *GetNodeForLocationRequest {
	r.opts["includeUserAgentShadowDOM"] = v
	return r
}

type GetNodeForLocationResult struct {
	// Id of the node at given coordinates.
	NodeId NodeId `json:"nodeId"`
}

func (r *GetNodeForLocationRequest) Do() (*GetNodeForLocationResult, error) {
	var result GetNodeForLocationResult
	err := r.client.Call("DOM.getNodeForLocation", r.opts, &result)
	return &result, err
}

// Returns the id of the nearest ancestor that is a relayout boundary. (experimental)
type GetRelayoutBoundaryRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) GetRelayoutBoundary() *GetRelayoutBoundaryRequest {
	return &GetRelayoutBoundaryRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Id of the node.
func (r *GetRelayoutBoundaryRequest) NodeId(v NodeId) *GetRelayoutBoundaryRequest {
	r.opts["nodeId"] = v
	return r
}

type GetRelayoutBoundaryResult struct {
	// Relayout boundary node id for the given node.
	NodeId NodeId `json:"nodeId"`
}

func (r *GetRelayoutBoundaryRequest) Do() (*GetRelayoutBoundaryResult, error) {
	var result GetRelayoutBoundaryResult
	err := r.client.Call("DOM.getRelayoutBoundary", r.opts, &result)
	return &result, err
}

func init() {
	rpc.EventTypes["DOM.documentUpdated"] = func() interface{} { return new(DocumentUpdatedEvent) }
	rpc.EventTypes["DOM.setChildNodes"] = func() interface{} { return new(SetChildNodesEvent) }
	rpc.EventTypes["DOM.attributeModified"] = func() interface{} { return new(AttributeModifiedEvent) }
	rpc.EventTypes["DOM.attributeRemoved"] = func() interface{} { return new(AttributeRemovedEvent) }
	rpc.EventTypes["DOM.inlineStyleInvalidated"] = func() interface{} { return new(InlineStyleInvalidatedEvent) }
	rpc.EventTypes["DOM.characterDataModified"] = func() interface{} { return new(CharacterDataModifiedEvent) }
	rpc.EventTypes["DOM.childNodeCountUpdated"] = func() interface{} { return new(ChildNodeCountUpdatedEvent) }
	rpc.EventTypes["DOM.childNodeInserted"] = func() interface{} { return new(ChildNodeInsertedEvent) }
	rpc.EventTypes["DOM.childNodeRemoved"] = func() interface{} { return new(ChildNodeRemovedEvent) }
	rpc.EventTypes["DOM.shadowRootPushed"] = func() interface{} { return new(ShadowRootPushedEvent) }
	rpc.EventTypes["DOM.shadowRootPopped"] = func() interface{} { return new(ShadowRootPoppedEvent) }
	rpc.EventTypes["DOM.pseudoElementAdded"] = func() interface{} { return new(PseudoElementAddedEvent) }
	rpc.EventTypes["DOM.pseudoElementRemoved"] = func() interface{} { return new(PseudoElementRemovedEvent) }
	rpc.EventTypes["DOM.distributedNodesUpdated"] = func() interface{} { return new(DistributedNodesUpdatedEvent) }
}

// Fired when <code>Document</code> has been totally updated. Node ids are no longer valid.
type DocumentUpdatedEvent struct {
}

// Fired when backend wants to provide client with the missing DOM structure. This happens upon most of the calls requesting node ids.
type SetChildNodesEvent struct {
	// Parent node id to populate with children.
	ParentId NodeId `json:"parentId"`

	// Child nodes array.
	Nodes []*Node `json:"nodes"`
}

// Fired when <code>Element</code>'s attribute is modified.
type AttributeModifiedEvent struct {
	// Id of the node that has changed.
	NodeId NodeId `json:"nodeId"`

	// Attribute name.
	Name string `json:"name"`

	// Attribute value.
	Value string `json:"value"`
}

// Fired when <code>Element</code>'s attribute is removed.
type AttributeRemovedEvent struct {
	// Id of the node that has changed.
	NodeId NodeId `json:"nodeId"`

	// A ttribute name.
	Name string `json:"name"`
}

// Fired when <code>Element</code>'s inline style is modified via a CSS property modification. (experimental)
type InlineStyleInvalidatedEvent struct {
	// Ids of the nodes for which the inline styles have been invalidated.
	NodeIds []NodeId `json:"nodeIds"`
}

// Mirrors <code>DOMCharacterDataModified</code> event.
type CharacterDataModifiedEvent struct {
	// Id of the node that has changed.
	NodeId NodeId `json:"nodeId"`

	// New text value.
	CharacterData string `json:"characterData"`
}

// Fired when <code>Container</code>'s child node count has changed.
type ChildNodeCountUpdatedEvent struct {
	// Id of the node that has changed.
	NodeId NodeId `json:"nodeId"`

	// New node count.
	ChildNodeCount int `json:"childNodeCount"`
}

// Mirrors <code>DOMNodeInserted</code> event.
type ChildNodeInsertedEvent struct {
	// Id of the node that has changed.
	ParentNodeId NodeId `json:"parentNodeId"`

	// If of the previous siblint.
	PreviousNodeId NodeId `json:"previousNodeId"`

	// Inserted node data.
	Node *Node `json:"node"`
}

// Mirrors <code>DOMNodeRemoved</code> event.
type ChildNodeRemovedEvent struct {
	// Parent id.
	ParentNodeId NodeId `json:"parentNodeId"`

	// Id of the node that has been removed.
	NodeId NodeId `json:"nodeId"`
}

// Called when shadow root is pushed into the element. (experimental)
type ShadowRootPushedEvent struct {
	// Host element id.
	HostId NodeId `json:"hostId"`

	// Shadow root.
	Root *Node `json:"root"`
}

// Called when shadow root is popped from the element. (experimental)
type ShadowRootPoppedEvent struct {
	// Host element id.
	HostId NodeId `json:"hostId"`

	// Shadow root id.
	RootId NodeId `json:"rootId"`
}

// Called when a pseudo element is added to an element. (experimental)
type PseudoElementAddedEvent struct {
	// Pseudo element's parent element id.
	ParentId NodeId `json:"parentId"`

	// The added pseudo element.
	PseudoElement *Node `json:"pseudoElement"`
}

// Called when a pseudo element is removed from an element. (experimental)
type PseudoElementRemovedEvent struct {
	// Pseudo element's parent element id.
	ParentId NodeId `json:"parentId"`

	// The removed pseudo element id.
	PseudoElementId NodeId `json:"pseudoElementId"`
}

// Called when distrubution is changed. (experimental)
type DistributedNodesUpdatedEvent struct {
	// Insertion point where distrubuted nodes were updated.
	InsertionPointId NodeId `json:"insertionPointId"`

	// Distributed nodes for given insertion point.
	DistributedNodes []*BackendNode `json:"distributedNodes"`
}
