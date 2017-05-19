// This domain exposes CSS read/write operations. All CSS objects (stylesheets, rules, and styles) have an associated <code>id</code> used in subsequent operations on the related object. Each object type has a specific <code>id</code> structure, and those are not interchangeable between objects of different kinds. CSS objects can be loaded using the <code>get*ForNode()</code> calls (which accept a DOM node id). A client can also discover all the existing stylesheets with the <code>getAllStyleSheets()</code> method (or keeping track of the <code>styleSheetAdded</code>/<code>styleSheetRemoved</code> events) and subsequently load the required stylesheet contents using the <code>getStyleSheet[Text]()</code> methods. (experimental)
package css

import (
	"encoding/json"
	"log"

	"github.com/neelance/cdp-go/rpc"
)

// This domain exposes CSS read/write operations. All CSS objects (stylesheets, rules, and styles) have an associated <code>id</code> used in subsequent operations on the related object. Each object type has a specific <code>id</code> structure, and those are not interchangeable between objects of different kinds. CSS objects can be loaded using the <code>get*ForNode()</code> calls (which accept a DOM node id). A client can also discover all the existing stylesheets with the <code>getAllStyleSheets()</code> method (or keeping track of the <code>styleSheetAdded</code>/<code>styleSheetRemoved</code> events) and subsequently load the required stylesheet contents using the <code>getStyleSheet[Text]()</code> methods. (experimental)
type Domain struct {
	Client *rpc.Client
}

type StyleSheetId string

// Stylesheet type: "injected" for stylesheets injected via extension, "user-agent" for user-agent stylesheets, "inspector" for stylesheets created by the inspector (i.e. those holding the "via inspector" rules), "regular" for regular stylesheets.

type StyleSheetOrigin string

// CSS rule collection for a single pseudo style.

type PseudoElementMatches struct {
	// Pseudo element type.
	PseudoType interface{} `json:"pseudoType"`

	// Matches of CSS rules applicable to the pseudo style.
	Matches []*RuleMatch `json:"matches"`
}

// Inherited CSS rule collection from ancestor node.

type InheritedStyleEntry struct {
	// The ancestor node's inline style, if any, in the style inheritance chain. (optional)
	InlineStyle *CSSStyle `json:"inlineStyle,omitempty"`

	// Matches of CSS rules matching the ancestor node in the style inheritance chain.
	MatchedCSSRules []*RuleMatch `json:"matchedCSSRules"`
}

// Match data for a CSS rule.

type RuleMatch struct {
	// CSS rule in the match.
	Rule *CSSRule `json:"rule"`

	// Matching selector indices in the rule's selectorList selectors (0-based).
	MatchingSelectors []int `json:"matchingSelectors"`
}

// Data for a simple selector (these are delimited by commas in a selector list).

type Value struct {
	// Value text.
	Text string `json:"text"`

	// Value range in the underlying resource (if available). (optional)
	Range *SourceRange `json:"range,omitempty"`
}

// Selector list data.

type SelectorList struct {
	// Selectors in the list.
	Selectors []*Value `json:"selectors"`

	// Rule selector text.
	Text string `json:"text"`
}

// CSS stylesheet metainformation.

type CSSStyleSheetHeader struct {
	// The stylesheet identifier.
	StyleSheetId *StyleSheetId `json:"styleSheetId"`

	// Owner frame identifier.
	FrameId interface{} `json:"frameId"`

	// Stylesheet resource URL.
	SourceURL string `json:"sourceURL"`

	// URL of source map associated with the stylesheet (if any). (optional)
	SourceMapURL string `json:"sourceMapURL,omitempty"`

	// Stylesheet origin.
	Origin *StyleSheetOrigin `json:"origin"`

	// Stylesheet title.
	Title string `json:"title"`

	// The backend id for the owner node of the stylesheet. (optional)
	OwnerNode interface{} `json:"ownerNode,omitempty"`

	// Denotes whether the stylesheet is disabled.
	Disabled bool `json:"disabled"`

	// Whether the sourceURL field value comes from the sourceURL comment. (optional)
	HasSourceURL bool `json:"hasSourceURL,omitempty"`

	// Whether this stylesheet is created for STYLE tag by parser. This flag is not set for document.written STYLE tags.
	IsInline bool `json:"isInline"`

	// Line offset of the stylesheet within the resource (zero based).
	StartLine float64 `json:"startLine"`

	// Column offset of the stylesheet within the resource (zero based).
	StartColumn float64 `json:"startColumn"`

	// Size of the content (in characters).
	Length float64 `json:"length"`
}

// CSS rule representation.

type CSSRule struct {
	// The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from. (optional)
	StyleSheetId *StyleSheetId `json:"styleSheetId,omitempty"`

	// Rule selector data.
	SelectorList *SelectorList `json:"selectorList"`

	// Parent stylesheet's origin.
	Origin *StyleSheetOrigin `json:"origin"`

	// Associated style declaration.
	Style *CSSStyle `json:"style"`

	// Media list array (for rules involving media queries). The array enumerates media queries starting with the innermost one, going outwards. (optional)
	Media []*CSSMedia `json:"media,omitempty"`
}

// CSS coverage information. (experimental)

type RuleUsage struct {
	// The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
	StyleSheetId *StyleSheetId `json:"styleSheetId"`

	// Offset of the start of the rule (including selector) from the beginning of the stylesheet.
	StartOffset float64 `json:"startOffset"`

	// Offset of the end of the rule body from the beginning of the stylesheet.
	EndOffset float64 `json:"endOffset"`

	// Indicates whether the rule was actually used by some element in the page.
	Used bool `json:"used"`
}

// Text range within a resource. All numbers are zero-based.

type SourceRange struct {
	// Start line of range.
	StartLine int `json:"startLine"`

	// Start column of range (inclusive).
	StartColumn int `json:"startColumn"`

	// End line of range
	EndLine int `json:"endLine"`

	// End column of range (exclusive).
	EndColumn int `json:"endColumn"`
}

type ShorthandEntry struct {
	// Shorthand name.
	Name string `json:"name"`

	// Shorthand value.
	Value string `json:"value"`

	// Whether the property has "!important" annotation (implies <code>false</code> if absent). (optional)
	Important bool `json:"important,omitempty"`
}

type CSSComputedStyleProperty struct {
	// Computed style property name.
	Name string `json:"name"`

	// Computed style property value.
	Value string `json:"value"`
}

// CSS style representation.

type CSSStyle struct {
	// The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from. (optional)
	StyleSheetId *StyleSheetId `json:"styleSheetId,omitempty"`

	// CSS properties in the style.
	CssProperties []*CSSProperty `json:"cssProperties"`

	// Computed values for all shorthands found in the style.
	ShorthandEntries []*ShorthandEntry `json:"shorthandEntries"`

	// Style declaration text (if available). (optional)
	CssText string `json:"cssText,omitempty"`

	// Style declaration range in the enclosing stylesheet (if available). (optional)
	Range *SourceRange `json:"range,omitempty"`
}

// CSS property declaration data.

type CSSProperty struct {
	// The property name.
	Name string `json:"name"`

	// The property value.
	Value string `json:"value"`

	// Whether the property has "!important" annotation (implies <code>false</code> if absent). (optional)
	Important bool `json:"important,omitempty"`

	// Whether the property is implicit (implies <code>false</code> if absent). (optional)
	Implicit bool `json:"implicit,omitempty"`

	// The full property text as specified in the style. (optional)
	Text string `json:"text,omitempty"`

	// Whether the property is understood by the browser (implies <code>true</code> if absent). (optional)
	ParsedOk bool `json:"parsedOk,omitempty"`

	// Whether the property is disabled by the user (present for source-based properties only). (optional)
	Disabled bool `json:"disabled,omitempty"`

	// The entire property range in the enclosing style declaration (if available). (optional)
	Range *SourceRange `json:"range,omitempty"`
}

// CSS media rule descriptor.

type CSSMedia struct {
	// Media query text.
	Text string `json:"text"`

	// Source of the media query: "mediaRule" if specified by a @media rule, "importRule" if specified by an @import rule, "linkedSheet" if specified by a "media" attribute in a linked stylesheet's LINK tag, "inlineSheet" if specified by a "media" attribute in an inline stylesheet's STYLE tag.
	Source string `json:"source"`

	// URL of the document containing the media query description. (optional)
	SourceURL string `json:"sourceURL,omitempty"`

	// The associated rule (@media or @import) header range in the enclosing stylesheet (if available). (optional)
	Range *SourceRange `json:"range,omitempty"`

	// Identifier of the stylesheet containing this object (if exists). (optional)
	StyleSheetId *StyleSheetId `json:"styleSheetId,omitempty"`

	// Array of media queries. (optional, experimental)
	MediaList []*MediaQuery `json:"mediaList,omitempty"`
}

// Media query descriptor. (experimental)

type MediaQuery struct {
	// Array of media query expressions.
	Expressions []*MediaQueryExpression `json:"expressions"`

	// Whether the media query condition is satisfied.
	Active bool `json:"active"`
}

// Media query expression descriptor. (experimental)

type MediaQueryExpression struct {
	// Media query expression value.
	Value float64 `json:"value"`

	// Media query expression units.
	Unit string `json:"unit"`

	// Media query expression feature.
	Feature string `json:"feature"`

	// The associated range of the value text in the enclosing stylesheet (if available). (optional)
	ValueRange *SourceRange `json:"valueRange,omitempty"`

	// Computed length of media query expression (if applicable). (optional)
	ComputedLength float64 `json:"computedLength,omitempty"`
}

// Information about amount of glyphs that were rendered with given font. (experimental)

type PlatformFontUsage struct {
	// Font's family name reported by platform.
	FamilyName string `json:"familyName"`

	// Indicates if the font was downloaded or resolved locally.
	IsCustomFont bool `json:"isCustomFont"`

	// Amount of glyphs that were rendered with this font.
	GlyphCount float64 `json:"glyphCount"`
}

// CSS keyframes rule representation.

type CSSKeyframesRule struct {
	// Animation name.
	AnimationName *Value `json:"animationName"`

	// List of keyframes.
	Keyframes []*CSSKeyframeRule `json:"keyframes"`
}

// CSS keyframe rule representation.

type CSSKeyframeRule struct {
	// The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from. (optional)
	StyleSheetId *StyleSheetId `json:"styleSheetId,omitempty"`

	// Parent stylesheet's origin.
	Origin *StyleSheetOrigin `json:"origin"`

	// Associated key text.
	KeyText *Value `json:"keyText"`

	// Associated style declaration.
	Style *CSSStyle `json:"style"`
}

// A descriptor of operation to mutate style declaration text.

type StyleDeclarationEdit struct {
	// The css style sheet identifier.
	StyleSheetId *StyleSheetId `json:"styleSheetId"`

	// The range of the style text in the enclosing stylesheet.
	Range *SourceRange `json:"range"`

	// New style text.
	Text string `json:"text"`
}

// Details of post layout rendered text positions. The exact layout should not be regarded as stable and may change between versions. (experimental)

type InlineTextBox struct {
	// The absolute position bounding box.
	BoundingBox interface{} `json:"boundingBox"`

	// The starting index in characters, for this post layout textbox substring.
	StartCharacterIndex int `json:"startCharacterIndex"`

	// The number of characters in this post layout textbox substring.
	NumCharacters int `json:"numCharacters"`
}

// Details of an element in the DOM tree with a LayoutObject. (experimental)

type LayoutTreeNode struct {
	// The id of the related DOM node matching one from DOM.GetDocument.
	NodeId interface{} `json:"nodeId"`

	// The absolute position bounding box.
	BoundingBox interface{} `json:"boundingBox"`

	// Contents of the LayoutText if any (optional)
	LayoutText string `json:"layoutText,omitempty"`

	// The post layout inline text nodes, if any. (optional)
	InlineTextNodes []*InlineTextBox `json:"inlineTextNodes,omitempty"`

	// Index into the computedStyles array returned by getLayoutTreeAndStyles. (optional)
	StyleIndex int `json:"styleIndex,omitempty"`
}

// A subset of the full ComputedStyle as defined by the request whitelist. (experimental)

type ComputedStyle struct {
	Properties []*CSSComputedStyleProperty `json:"properties"`
}

// Enables the CSS agent for the given page. Clients should not assume that the CSS agent has been enabled until the result of this command is received.
func (d *Domain) Enable() error {
	return d.Client.Call("CSS.enable", nil, nil)
}

// Disables the CSS agent for the given page.
func (d *Domain) Disable() error {
	return d.Client.Call("CSS.disable", nil, nil)
}

type GetMatchedStylesForNodeOpts struct {
	NodeId interface{} `json:"nodeId"`
}

type GetMatchedStylesForNodeResult struct {
	// Inline style for the specified DOM node. (optional)
	InlineStyle *CSSStyle `json:"inlineStyle"`

	// Attribute-defined element style (e.g. resulting from "width=20 height=100%"). (optional)
	AttributesStyle *CSSStyle `json:"attributesStyle"`

	// CSS rules matching this node, from all applicable stylesheets. (optional)
	MatchedCSSRules []*RuleMatch `json:"matchedCSSRules"`

	// Pseudo style matches for this node. (optional)
	PseudoElements []*PseudoElementMatches `json:"pseudoElements"`

	// A chain of inherited styles (from the immediate node parent up to the DOM tree root). (optional)
	Inherited []*InheritedStyleEntry `json:"inherited"`

	// A list of CSS keyframed animations matching this node. (optional)
	CssKeyframesRules []*CSSKeyframesRule `json:"cssKeyframesRules"`
}

// Returns requested styles for a DOM node identified by <code>nodeId</code>.
func (d *Domain) GetMatchedStylesForNode(opts *GetMatchedStylesForNodeOpts) (*GetMatchedStylesForNodeResult, error) {
	var result GetMatchedStylesForNodeResult
	err := d.Client.Call("CSS.getMatchedStylesForNode", opts, &result)
	return &result, err
}

type GetInlineStylesForNodeOpts struct {
	NodeId interface{} `json:"nodeId"`
}

type GetInlineStylesForNodeResult struct {
	// Inline style for the specified DOM node. (optional)
	InlineStyle *CSSStyle `json:"inlineStyle"`

	// Attribute-defined element style (e.g. resulting from "width=20 height=100%"). (optional)
	AttributesStyle *CSSStyle `json:"attributesStyle"`
}

// Returns the styles defined inline (explicitly in the "style" attribute and implicitly, using DOM attributes) for a DOM node identified by <code>nodeId</code>.
func (d *Domain) GetInlineStylesForNode(opts *GetInlineStylesForNodeOpts) (*GetInlineStylesForNodeResult, error) {
	var result GetInlineStylesForNodeResult
	err := d.Client.Call("CSS.getInlineStylesForNode", opts, &result)
	return &result, err
}

type GetComputedStyleForNodeOpts struct {
	NodeId interface{} `json:"nodeId"`
}

type GetComputedStyleForNodeResult struct {
	// Computed style for the specified DOM node.
	ComputedStyle []*CSSComputedStyleProperty `json:"computedStyle"`
}

// Returns the computed style for a DOM node identified by <code>nodeId</code>.
func (d *Domain) GetComputedStyleForNode(opts *GetComputedStyleForNodeOpts) (*GetComputedStyleForNodeResult, error) {
	var result GetComputedStyleForNodeResult
	err := d.Client.Call("CSS.getComputedStyleForNode", opts, &result)
	return &result, err
}

type GetPlatformFontsForNodeOpts struct {
	NodeId interface{} `json:"nodeId"`
}

type GetPlatformFontsForNodeResult struct {
	// Usage statistics for every employed platform font.
	Fonts []*PlatformFontUsage `json:"fonts"`
}

// Requests information about platform fonts which we used to render child TextNodes in the given node. (experimental)
func (d *Domain) GetPlatformFontsForNode(opts *GetPlatformFontsForNodeOpts) (*GetPlatformFontsForNodeResult, error) {
	var result GetPlatformFontsForNodeResult
	err := d.Client.Call("CSS.getPlatformFontsForNode", opts, &result)
	return &result, err
}

type GetStyleSheetTextOpts struct {
	StyleSheetId *StyleSheetId `json:"styleSheetId"`
}

type GetStyleSheetTextResult struct {
	// The stylesheet text.
	Text string `json:"text"`
}

// Returns the current textual content and the URL for a stylesheet.
func (d *Domain) GetStyleSheetText(opts *GetStyleSheetTextOpts) (*GetStyleSheetTextResult, error) {
	var result GetStyleSheetTextResult
	err := d.Client.Call("CSS.getStyleSheetText", opts, &result)
	return &result, err
}

type CollectClassNamesOpts struct {
	StyleSheetId *StyleSheetId `json:"styleSheetId"`
}

type CollectClassNamesResult struct {
	// Class name list.
	ClassNames []string `json:"classNames"`
}

// Returns all class names from specified stylesheet. (experimental)
func (d *Domain) CollectClassNames(opts *CollectClassNamesOpts) (*CollectClassNamesResult, error) {
	var result CollectClassNamesResult
	err := d.Client.Call("CSS.collectClassNames", opts, &result)
	return &result, err
}

type SetStyleSheetTextOpts struct {
	StyleSheetId *StyleSheetId `json:"styleSheetId"`

	Text string `json:"text"`
}

type SetStyleSheetTextResult struct {
	// URL of source map associated with script (if any). (optional)
	SourceMapURL string `json:"sourceMapURL"`
}

// Sets the new stylesheet text.
func (d *Domain) SetStyleSheetText(opts *SetStyleSheetTextOpts) (*SetStyleSheetTextResult, error) {
	var result SetStyleSheetTextResult
	err := d.Client.Call("CSS.setStyleSheetText", opts, &result)
	return &result, err
}

type SetRuleSelectorOpts struct {
	StyleSheetId *StyleSheetId `json:"styleSheetId"`

	Range *SourceRange `json:"range"`

	Selector string `json:"selector"`
}

type SetRuleSelectorResult struct {
	// The resulting selector list after modification.
	SelectorList *SelectorList `json:"selectorList"`
}

// Modifies the rule selector.
func (d *Domain) SetRuleSelector(opts *SetRuleSelectorOpts) (*SetRuleSelectorResult, error) {
	var result SetRuleSelectorResult
	err := d.Client.Call("CSS.setRuleSelector", opts, &result)
	return &result, err
}

type SetKeyframeKeyOpts struct {
	StyleSheetId *StyleSheetId `json:"styleSheetId"`

	Range *SourceRange `json:"range"`

	KeyText string `json:"keyText"`
}

type SetKeyframeKeyResult struct {
	// The resulting key text after modification.
	KeyText *Value `json:"keyText"`
}

// Modifies the keyframe rule key text.
func (d *Domain) SetKeyframeKey(opts *SetKeyframeKeyOpts) (*SetKeyframeKeyResult, error) {
	var result SetKeyframeKeyResult
	err := d.Client.Call("CSS.setKeyframeKey", opts, &result)
	return &result, err
}

type SetStyleTextsOpts struct {
	Edits []*StyleDeclarationEdit `json:"edits"`
}

type SetStyleTextsResult struct {
	// The resulting styles after modification.
	Styles []*CSSStyle `json:"styles"`
}

// Applies specified style edits one after another in the given order.
func (d *Domain) SetStyleTexts(opts *SetStyleTextsOpts) (*SetStyleTextsResult, error) {
	var result SetStyleTextsResult
	err := d.Client.Call("CSS.setStyleTexts", opts, &result)
	return &result, err
}

type SetMediaTextOpts struct {
	StyleSheetId *StyleSheetId `json:"styleSheetId"`

	Range *SourceRange `json:"range"`

	Text string `json:"text"`
}

type SetMediaTextResult struct {
	// The resulting CSS media rule after modification.
	Media *CSSMedia `json:"media"`
}

// Modifies the rule selector.
func (d *Domain) SetMediaText(opts *SetMediaTextOpts) (*SetMediaTextResult, error) {
	var result SetMediaTextResult
	err := d.Client.Call("CSS.setMediaText", opts, &result)
	return &result, err
}

type CreateStyleSheetOpts struct {
	// Identifier of the frame where "via-inspector" stylesheet should be created.
	FrameId interface{} `json:"frameId"`
}

type CreateStyleSheetResult struct {
	// Identifier of the created "via-inspector" stylesheet.
	StyleSheetId *StyleSheetId `json:"styleSheetId"`
}

// Creates a new special "via-inspector" stylesheet in the frame with given <code>frameId</code>.
func (d *Domain) CreateStyleSheet(opts *CreateStyleSheetOpts) (*CreateStyleSheetResult, error) {
	var result CreateStyleSheetResult
	err := d.Client.Call("CSS.createStyleSheet", opts, &result)
	return &result, err
}

type AddRuleOpts struct {
	// The css style sheet identifier where a new rule should be inserted.
	StyleSheetId *StyleSheetId `json:"styleSheetId"`

	// The text of a new rule.
	RuleText string `json:"ruleText"`

	// Text position of a new rule in the target style sheet.
	Location *SourceRange `json:"location"`
}

type AddRuleResult struct {
	// The newly created rule.
	Rule *CSSRule `json:"rule"`
}

// Inserts a new rule with the given <code>ruleText</code> in a stylesheet with given <code>styleSheetId</code>, at the position specified by <code>location</code>.
func (d *Domain) AddRule(opts *AddRuleOpts) (*AddRuleResult, error) {
	var result AddRuleResult
	err := d.Client.Call("CSS.addRule", opts, &result)
	return &result, err
}

type ForcePseudoStateOpts struct {
	// The element id for which to force the pseudo state.
	NodeId interface{} `json:"nodeId"`

	// Element pseudo classes to force when computing the element's style.
	ForcedPseudoClasses []string `json:"forcedPseudoClasses"`
}

// Ensures that the given node will have specified pseudo-classes whenever its style is computed by the browser.
func (d *Domain) ForcePseudoState(opts *ForcePseudoStateOpts) error {
	return d.Client.Call("CSS.forcePseudoState", opts, nil)
}

type GetMediaQueriesResult struct {
	Medias []*CSSMedia `json:"medias"`
}

// Returns all media queries parsed by the rendering engine. (experimental)
func (d *Domain) GetMediaQueries() (*GetMediaQueriesResult, error) {
	var result GetMediaQueriesResult
	err := d.Client.Call("CSS.getMediaQueries", nil, &result)
	return &result, err
}

type SetEffectivePropertyValueForNodeOpts struct {
	// The element id for which to set property.
	NodeId interface{} `json:"nodeId"`

	PropertyName string `json:"propertyName"`

	Value string `json:"value"`
}

// Find a rule with the given active property for the given node and set the new value for this property (experimental)
func (d *Domain) SetEffectivePropertyValueForNode(opts *SetEffectivePropertyValueForNodeOpts) error {
	return d.Client.Call("CSS.setEffectivePropertyValueForNode", opts, nil)
}

type GetBackgroundColorsOpts struct {
	// Id of the node to get background colors for.
	NodeId interface{} `json:"nodeId"`
}

type GetBackgroundColorsResult struct {
	// The range of background colors behind this element, if it contains any visible text. If no visible text is present, this will be undefined. In the case of a flat background color, this will consist of simply that color. In the case of a gradient, this will consist of each of the color stops. For anything more complicated, this will be an empty array. Images will be ignored (as if the image had failed to load). (optional)
	BackgroundColors []string `json:"backgroundColors"`
}

// (experimental)
func (d *Domain) GetBackgroundColors(opts *GetBackgroundColorsOpts) (*GetBackgroundColorsResult, error) {
	var result GetBackgroundColorsResult
	err := d.Client.Call("CSS.getBackgroundColors", opts, &result)
	return &result, err
}

type GetLayoutTreeAndStylesOpts struct {
	// Whitelist of computed styles to return.
	ComputedStyleWhitelist []string `json:"computedStyleWhitelist"`
}

type GetLayoutTreeAndStylesResult struct {
	LayoutTreeNodes []*LayoutTreeNode `json:"layoutTreeNodes"`

	ComputedStyles []*ComputedStyle `json:"computedStyles"`
}

// For the main document and any content documents, return the LayoutTreeNodes and a whitelisted subset of the computed style. It only returns pushed nodes, on way to pull all nodes is to call DOM.getDocument with a depth of -1. (experimental)
func (d *Domain) GetLayoutTreeAndStyles(opts *GetLayoutTreeAndStylesOpts) (*GetLayoutTreeAndStylesResult, error) {
	var result GetLayoutTreeAndStylesResult
	err := d.Client.Call("CSS.getLayoutTreeAndStyles", opts, &result)
	return &result, err
}

// Enables the selector recording. (experimental)
func (d *Domain) StartRuleUsageTracking() error {
	return d.Client.Call("CSS.startRuleUsageTracking", nil, nil)
}

type TakeCoverageDeltaResult struct {
	Coverage []*RuleUsage `json:"coverage"`
}

// Obtain list of rules that became used since last call to this method (or since start of coverage instrumentation) (experimental)
func (d *Domain) TakeCoverageDelta() (*TakeCoverageDeltaResult, error) {
	var result TakeCoverageDeltaResult
	err := d.Client.Call("CSS.takeCoverageDelta", nil, &result)
	return &result, err
}

type StopRuleUsageTrackingResult struct {
	RuleUsage []*RuleUsage `json:"ruleUsage"`
}

// The list of rules with an indication of whether these were used (experimental)
func (d *Domain) StopRuleUsageTracking() (*StopRuleUsageTrackingResult, error) {
	var result StopRuleUsageTrackingResult
	err := d.Client.Call("CSS.stopRuleUsageTracking", nil, &result)
	return &result, err
}

type MediaQueryResultChangedEvent struct {
}

// Fires whenever a MediaQuery result changes (for example, after a browser window has been resized.) The current implementation considers only viewport-dependent media features.
func (d *Domain) OnMediaQueryResultChanged(listener func(*MediaQueryResultChangedEvent)) {
	d.Client.AddListener("CSS.mediaQueryResultChanged", func(params json.RawMessage) {
		var event MediaQueryResultChangedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type FontsUpdatedEvent struct {
}

// Fires whenever a web font gets loaded.
func (d *Domain) OnFontsUpdated(listener func(*FontsUpdatedEvent)) {
	d.Client.AddListener("CSS.fontsUpdated", func(params json.RawMessage) {
		var event FontsUpdatedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type StyleSheetChangedEvent struct {
	StyleSheetId *StyleSheetId `json:"styleSheetId"`
}

// Fired whenever a stylesheet is changed as a result of the client operation.
func (d *Domain) OnStyleSheetChanged(listener func(*StyleSheetChangedEvent)) {
	d.Client.AddListener("CSS.styleSheetChanged", func(params json.RawMessage) {
		var event StyleSheetChangedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type StyleSheetAddedEvent struct {
	// Added stylesheet metainfo.
	Header *CSSStyleSheetHeader `json:"header"`
}

// Fired whenever an active document stylesheet is added.
func (d *Domain) OnStyleSheetAdded(listener func(*StyleSheetAddedEvent)) {
	d.Client.AddListener("CSS.styleSheetAdded", func(params json.RawMessage) {
		var event StyleSheetAddedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type StyleSheetRemovedEvent struct {
	// Identifier of the removed stylesheet.
	StyleSheetId *StyleSheetId `json:"styleSheetId"`
}

// Fired whenever an active document stylesheet is removed.
func (d *Domain) OnStyleSheetRemoved(listener func(*StyleSheetRemovedEvent)) {
	d.Client.AddListener("CSS.styleSheetRemoved", func(params json.RawMessage) {
		var event StyleSheetRemovedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}
