// Actions and events related to the inspected page belong to the page domain.
package page

import (
	"encoding/json"
	"log"

	"github.com/neelance/cdp-go/rpc"
)

// Actions and events related to the inspected page belong to the page domain.
type Domain struct {
	Client *rpc.Client
}

// Resource type as it was perceived by the rendering engine.

type ResourceType string

// Unique frame identifier.

type FrameId string

// Information about the Frame on the page.

type Frame struct {
	// Frame unique identifier.
	Id string `json:"id"`

	// Parent frame identifier. (optional)
	ParentId string `json:"parentId,omitempty"`

	// Identifier of the loader associated with this frame.
	LoaderId interface{} `json:"loaderId"`

	// Frame's name as specified in the tag. (optional)
	Name string `json:"name,omitempty"`

	// Frame document's URL.
	URL string `json:"url"`

	// Frame document's security origin.
	SecurityOrigin string `json:"securityOrigin"`

	// Frame document's mimeType as determined by the browser.
	MimeType string `json:"mimeType"`
}

// Information about the Resource on the page. (experimental)

type FrameResource struct {
	// Resource URL.
	URL string `json:"url"`

	// Type of this resource.
	Type *ResourceType `json:"type"`

	// Resource mimeType as determined by the browser.
	MimeType string `json:"mimeType"`

	// last-modified timestamp as reported by server. (optional)
	LastModified interface{} `json:"lastModified,omitempty"`

	// Resource content size. (optional)
	ContentSize float64 `json:"contentSize,omitempty"`

	// True if the resource failed to load. (optional)
	Failed bool `json:"failed,omitempty"`

	// True if the resource was canceled during loading. (optional)
	Canceled bool `json:"canceled,omitempty"`
}

// Information about the Frame hierarchy along with their cached resources. (experimental)

type FrameResourceTree struct {
	// Frame information for this tree item.
	Frame *Frame `json:"frame"`

	// Child frames. (optional)
	ChildFrames []*FrameResourceTree `json:"childFrames,omitempty"`

	// Information about frame resources.
	Resources []*FrameResource `json:"resources"`
}

// Unique script identifier. (experimental)

type ScriptIdentifier string

// Navigation history entry. (experimental)

type NavigationEntry struct {
	// Unique id of the navigation history entry.
	Id int `json:"id"`

	// URL of the navigation history entry.
	URL string `json:"url"`

	// Title of the navigation history entry.
	Title string `json:"title"`
}

// Screencast frame metadata. (experimental)

type ScreencastFrameMetadata struct {
	// Top offset in DIP.
	OffsetTop float64 `json:"offsetTop"`

	// Page scale factor.
	PageScaleFactor float64 `json:"pageScaleFactor"`

	// Device screen width in DIP.
	DeviceWidth float64 `json:"deviceWidth"`

	// Device screen height in DIP.
	DeviceHeight float64 `json:"deviceHeight"`

	// Position of horizontal scroll in CSS pixels.
	ScrollOffsetX float64 `json:"scrollOffsetX"`

	// Position of vertical scroll in CSS pixels.
	ScrollOffsetY float64 `json:"scrollOffsetY"`

	// Frame swap timestamp. (optional, experimental)
	Timestamp float64 `json:"timestamp,omitempty"`
}

// Javascript dialog type. (experimental)

type DialogType string

// Error while paring app manifest. (experimental)

type AppManifestError struct {
	// Error message.
	Message string `json:"message"`

	// If criticial, this is a non-recoverable parse error.
	Critical int `json:"critical"`

	// Error line.
	Line int `json:"line"`

	// Error column.
	Column int `json:"column"`
}

// Proceed: allow the navigation; Cancel: cancel the navigation; CancelAndIgnore: cancels the navigation and makes the requester of the navigation acts like the request was never made. (experimental)

type NavigationResponse string

// Layout viewport position and dimensions. (experimental)

type LayoutViewport struct {
	// Horizontal offset relative to the document (CSS pixels).
	PageX int `json:"pageX"`

	// Vertical offset relative to the document (CSS pixels).
	PageY int `json:"pageY"`

	// Width (CSS pixels), excludes scrollbar if present.
	ClientWidth int `json:"clientWidth"`

	// Height (CSS pixels), excludes scrollbar if present.
	ClientHeight int `json:"clientHeight"`
}

// Visual viewport position, dimensions, and scale. (experimental)

type VisualViewport struct {
	// Horizontal offset relative to the layout viewport (CSS pixels).
	OffsetX float64 `json:"offsetX"`

	// Vertical offset relative to the layout viewport (CSS pixels).
	OffsetY float64 `json:"offsetY"`

	// Horizontal offset relative to the document (CSS pixels).
	PageX float64 `json:"pageX"`

	// Vertical offset relative to the document (CSS pixels).
	PageY float64 `json:"pageY"`

	// Width (CSS pixels), excludes scrollbar if present.
	ClientWidth float64 `json:"clientWidth"`

	// Height (CSS pixels), excludes scrollbar if present.
	ClientHeight float64 `json:"clientHeight"`

	// Scale relative to the ideal viewport (size at width=device-width).
	Scale float64 `json:"scale"`
}

// Enables page domain notifications.
func (d *Domain) Enable() error {
	return d.Client.Call("Page.enable", nil, nil)
}

// Disables page domain notifications.
func (d *Domain) Disable() error {
	return d.Client.Call("Page.disable", nil, nil)
}

type AddScriptToEvaluateOnLoadOpts struct {
	ScriptSource string `json:"scriptSource"`
}

type AddScriptToEvaluateOnLoadResult struct {
	// Identifier of the added script.
	Identifier *ScriptIdentifier `json:"identifier"`
}

// (experimental)
func (d *Domain) AddScriptToEvaluateOnLoad(opts *AddScriptToEvaluateOnLoadOpts) (*AddScriptToEvaluateOnLoadResult, error) {
	var result AddScriptToEvaluateOnLoadResult
	err := d.Client.Call("Page.addScriptToEvaluateOnLoad", opts, &result)
	return &result, err
}

type RemoveScriptToEvaluateOnLoadOpts struct {
	Identifier *ScriptIdentifier `json:"identifier"`
}

// (experimental)
func (d *Domain) RemoveScriptToEvaluateOnLoad(opts *RemoveScriptToEvaluateOnLoadOpts) error {
	return d.Client.Call("Page.removeScriptToEvaluateOnLoad", opts, nil)
}

type SetAutoAttachToCreatedPagesOpts struct {
	// If true, browser will open a new inspector window for every page created from this one.
	AutoAttach bool `json:"autoAttach"`
}

// Controls whether browser will open a new inspector window for connected pages. (experimental)
func (d *Domain) SetAutoAttachToCreatedPages(opts *SetAutoAttachToCreatedPagesOpts) error {
	return d.Client.Call("Page.setAutoAttachToCreatedPages", opts, nil)
}

type ReloadOpts struct {
	// If true, browser cache is ignored (as if the user pressed Shift+refresh). (optional)
	IgnoreCache bool `json:"ignoreCache,omitempty"`

	// If set, the script will be injected into all frames of the inspected page after reload. (optional)
	ScriptToEvaluateOnLoad string `json:"scriptToEvaluateOnLoad,omitempty"`
}

// Reloads given page optionally ignoring the cache.
func (d *Domain) Reload(opts *ReloadOpts) error {
	return d.Client.Call("Page.reload", opts, nil)
}

type NavigateOpts struct {
	// URL to navigate the page to.
	URL string `json:"url"`

	// Referrer URL. (optional, experimental)
	Referrer string `json:"referrer,omitempty"`
}

type NavigateResult struct {
	// Frame id that will be navigated.
	FrameId *FrameId `json:"frameId"`
}

// Navigates current page to the given URL.
func (d *Domain) Navigate(opts *NavigateOpts) (*NavigateResult, error) {
	var result NavigateResult
	err := d.Client.Call("Page.navigate", opts, &result)
	return &result, err
}

// Force the page stop all navigations and pending resource fetches. (experimental)
func (d *Domain) StopLoading() error {
	return d.Client.Call("Page.stopLoading", nil, nil)
}

type GetNavigationHistoryResult struct {
	// Index of the current navigation history entry.
	CurrentIndex int `json:"currentIndex"`

	// Array of navigation history entries.
	Entries []*NavigationEntry `json:"entries"`
}

// Returns navigation history for the current page. (experimental)
func (d *Domain) GetNavigationHistory() (*GetNavigationHistoryResult, error) {
	var result GetNavigationHistoryResult
	err := d.Client.Call("Page.getNavigationHistory", nil, &result)
	return &result, err
}

type NavigateToHistoryEntryOpts struct {
	// Unique id of the entry to navigate to.
	EntryId int `json:"entryId"`
}

// Navigates current page to the given history entry. (experimental)
func (d *Domain) NavigateToHistoryEntry(opts *NavigateToHistoryEntryOpts) error {
	return d.Client.Call("Page.navigateToHistoryEntry", opts, nil)
}

type GetCookiesResult struct {
	// Array of cookie objects.
	Cookies []interface{} `json:"cookies"`
}

// Returns all browser cookies. Depending on the backend support, will return detailed cookie information in the <code>cookies</code> field. (experimental)
func (d *Domain) GetCookies() (*GetCookiesResult, error) {
	var result GetCookiesResult
	err := d.Client.Call("Page.getCookies", nil, &result)
	return &result, err
}

type DeleteCookieOpts struct {
	// Name of the cookie to remove.
	CookieName string `json:"cookieName"`

	// URL to match cooke domain and path.
	URL string `json:"url"`
}

// Deletes browser cookie with given name, domain and path. (experimental)
func (d *Domain) DeleteCookie(opts *DeleteCookieOpts) error {
	return d.Client.Call("Page.deleteCookie", opts, nil)
}

type GetResourceTreeResult struct {
	// Present frame / resource tree structure.
	FrameTree *FrameResourceTree `json:"frameTree"`
}

// Returns present frame / resource tree structure. (experimental)
func (d *Domain) GetResourceTree() (*GetResourceTreeResult, error) {
	var result GetResourceTreeResult
	err := d.Client.Call("Page.getResourceTree", nil, &result)
	return &result, err
}

type GetResourceContentOpts struct {
	// Frame id to get resource for.
	FrameId *FrameId `json:"frameId"`

	// URL of the resource to get content for.
	URL string `json:"url"`
}

type GetResourceContentResult struct {
	// Resource content.
	Content string `json:"content"`

	// True, if content was served as base64.
	Base64Encoded bool `json:"base64Encoded"`
}

// Returns content of the given resource. (experimental)
func (d *Domain) GetResourceContent(opts *GetResourceContentOpts) (*GetResourceContentResult, error) {
	var result GetResourceContentResult
	err := d.Client.Call("Page.getResourceContent", opts, &result)
	return &result, err
}

type SearchInResourceOpts struct {
	// Frame id for resource to search in.
	FrameId *FrameId `json:"frameId"`

	// URL of the resource to search in.
	URL string `json:"url"`

	// String to search for.
	Query string `json:"query"`

	// If true, search is case sensitive. (optional)
	CaseSensitive bool `json:"caseSensitive,omitempty"`

	// If true, treats string parameter as regex. (optional)
	IsRegex bool `json:"isRegex,omitempty"`
}

type SearchInResourceResult struct {
	// List of search matches.
	Result []interface{} `json:"result"`
}

// Searches for given string in resource content. (experimental)
func (d *Domain) SearchInResource(opts *SearchInResourceOpts) (*SearchInResourceResult, error) {
	var result SearchInResourceResult
	err := d.Client.Call("Page.searchInResource", opts, &result)
	return &result, err
}

type SetDocumentContentOpts struct {
	// Frame id to set HTML for.
	FrameId *FrameId `json:"frameId"`

	// HTML content to set.
	Html string `json:"html"`
}

// Sets given markup as the document's HTML. (experimental)
func (d *Domain) SetDocumentContent(opts *SetDocumentContentOpts) error {
	return d.Client.Call("Page.setDocumentContent", opts, nil)
}

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

	// Scale to apply to resulting view image. Ignored in |fitWindow| mode. (optional)
	Scale float64 `json:"scale,omitempty"`

	// X offset to shift resulting view image by. Ignored in |fitWindow| mode. (optional)
	OffsetX float64 `json:"offsetX,omitempty"`

	// Y offset to shift resulting view image by. Ignored in |fitWindow| mode. (optional)
	OffsetY float64 `json:"offsetY,omitempty"`

	// Overriding screen width value in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|. (optional)
	ScreenWidth int `json:"screenWidth,omitempty"`

	// Overriding screen height value in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|. (optional)
	ScreenHeight int `json:"screenHeight,omitempty"`

	// Overriding view X position on screen in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|. (optional)
	PositionX int `json:"positionX,omitempty"`

	// Overriding view Y position on screen in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|. (optional)
	PositionY int `json:"positionY,omitempty"`

	// Screen orientation override. (optional)
	ScreenOrientation interface{} `json:"screenOrientation,omitempty"`
}

// Overrides the values of device screen dimensions (window.screen.width, window.screen.height, window.innerWidth, window.innerHeight, and "device-width"/"device-height"-related CSS media query results). (experimental)
func (d *Domain) SetDeviceMetricsOverride(opts *SetDeviceMetricsOverrideOpts) error {
	return d.Client.Call("Page.setDeviceMetricsOverride", opts, nil)
}

// Clears the overriden device metrics. (experimental)
func (d *Domain) ClearDeviceMetricsOverride() error {
	return d.Client.Call("Page.clearDeviceMetricsOverride", nil, nil)
}

type SetGeolocationOverrideOpts struct {
	// Mock latitude (optional)
	Latitude float64 `json:"latitude,omitempty"`

	// Mock longitude (optional)
	Longitude float64 `json:"longitude,omitempty"`

	// Mock accuracy (optional)
	Accuracy float64 `json:"accuracy,omitempty"`
}

// Overrides the Geolocation Position or Error. Omitting any of the parameters emulates position unavailable.
func (d *Domain) SetGeolocationOverride(opts *SetGeolocationOverrideOpts) error {
	return d.Client.Call("Page.setGeolocationOverride", opts, nil)
}

// Clears the overriden Geolocation Position and Error.
func (d *Domain) ClearGeolocationOverride() error {
	return d.Client.Call("Page.clearGeolocationOverride", nil, nil)
}

type SetDeviceOrientationOverrideOpts struct {
	// Mock alpha
	Alpha float64 `json:"alpha"`

	// Mock beta
	Beta float64 `json:"beta"`

	// Mock gamma
	Gamma float64 `json:"gamma"`
}

// Overrides the Device Orientation. (experimental)
func (d *Domain) SetDeviceOrientationOverride(opts *SetDeviceOrientationOverrideOpts) error {
	return d.Client.Call("Page.setDeviceOrientationOverride", opts, nil)
}

// Clears the overridden Device Orientation. (experimental)
func (d *Domain) ClearDeviceOrientationOverride() error {
	return d.Client.Call("Page.clearDeviceOrientationOverride", nil, nil)
}

type SetTouchEmulationEnabledOpts struct {
	// Whether the touch event emulation should be enabled.
	Enabled bool `json:"enabled"`

	// Touch/gesture events configuration. Default: current platform. (optional)
	Configuration string `json:"configuration,omitempty"`
}

// Toggles mouse event-based touch event emulation. (experimental)
func (d *Domain) SetTouchEmulationEnabled(opts *SetTouchEmulationEnabledOpts) error {
	return d.Client.Call("Page.setTouchEmulationEnabled", opts, nil)
}

type CaptureScreenshotOpts struct {
	// Image compression format (defaults to png). (optional)
	Format string `json:"format,omitempty"`

	// Compression quality from range [0..100] (jpeg only). (optional)
	Quality int `json:"quality,omitempty"`

	// Capture the screenshot from the surface, rather than the view. Defaults to false. (optional, experimental)
	FromSurface bool `json:"fromSurface,omitempty"`
}

type CaptureScreenshotResult struct {
	// Base64-encoded image data.
	Data string `json:"data"`
}

// Capture page screenshot. (experimental)
func (d *Domain) CaptureScreenshot(opts *CaptureScreenshotOpts) (*CaptureScreenshotResult, error) {
	var result CaptureScreenshotResult
	err := d.Client.Call("Page.captureScreenshot", opts, &result)
	return &result, err
}

type PrintToPDFOpts struct {
	// Paper orientation. Defaults to false. (optional)
	Landscape bool `json:"landscape,omitempty"`

	// Display header and footer. Defaults to false. (optional)
	DisplayHeaderFooter bool `json:"displayHeaderFooter,omitempty"`

	// Print background graphics. Defaults to false. (optional)
	PrintBackground bool `json:"printBackground,omitempty"`

	// Scale of the webpage rendering. Defaults to 1. (optional)
	Scale float64 `json:"scale,omitempty"`

	// Paper width in inches. Defaults to 8.5 inches. (optional)
	PaperWidth float64 `json:"paperWidth,omitempty"`

	// Paper height in inches. Defaults to 11 inches. (optional)
	PaperHeight float64 `json:"paperHeight,omitempty"`

	// Top margin in inches. Defaults to 1cm (~0.4 inches). (optional)
	MarginTop float64 `json:"marginTop,omitempty"`

	// Bottom margin in inches. Defaults to 1cm (~0.4 inches). (optional)
	MarginBottom float64 `json:"marginBottom,omitempty"`

	// Left margin in inches. Defaults to 1cm (~0.4 inches). (optional)
	MarginLeft float64 `json:"marginLeft,omitempty"`

	// Right margin in inches. Defaults to 1cm (~0.4 inches). (optional)
	MarginRight float64 `json:"marginRight,omitempty"`

	// Paper ranges to print, e.g., '1-5, 8, 11-13'. Defaults to the empty string, which means print all pages. (optional)
	PageRanges string `json:"pageRanges,omitempty"`
}

type PrintToPDFResult struct {
	// Base64-encoded pdf data.
	Data string `json:"data"`
}

// Print page as PDF. (experimental)
func (d *Domain) PrintToPDF(opts *PrintToPDFOpts) (*PrintToPDFResult, error) {
	var result PrintToPDFResult
	err := d.Client.Call("Page.printToPDF", opts, &result)
	return &result, err
}

type StartScreencastOpts struct {
	// Image compression format. (optional)
	Format string `json:"format,omitempty"`

	// Compression quality from range [0..100]. (optional)
	Quality int `json:"quality,omitempty"`

	// Maximum screenshot width. (optional)
	MaxWidth int `json:"maxWidth,omitempty"`

	// Maximum screenshot height. (optional)
	MaxHeight int `json:"maxHeight,omitempty"`

	// Send every n-th frame. (optional)
	EveryNthFrame int `json:"everyNthFrame,omitempty"`
}

// Starts sending each frame using the <code>screencastFrame</code> event. (experimental)
func (d *Domain) StartScreencast(opts *StartScreencastOpts) error {
	return d.Client.Call("Page.startScreencast", opts, nil)
}

// Stops sending each frame in the <code>screencastFrame</code>. (experimental)
func (d *Domain) StopScreencast() error {
	return d.Client.Call("Page.stopScreencast", nil, nil)
}

type ScreencastFrameAckOpts struct {
	// Frame number.
	SessionId int `json:"sessionId"`
}

// Acknowledges that a screencast frame has been received by the frontend. (experimental)
func (d *Domain) ScreencastFrameAck(opts *ScreencastFrameAckOpts) error {
	return d.Client.Call("Page.screencastFrameAck", opts, nil)
}

type HandleJavaScriptDialogOpts struct {
	// Whether to accept or dismiss the dialog.
	Accept bool `json:"accept"`

	// The text to enter into the dialog prompt before accepting. Used only if this is a prompt dialog. (optional)
	PromptText string `json:"promptText,omitempty"`
}

// Accepts or dismisses a JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload).
func (d *Domain) HandleJavaScriptDialog(opts *HandleJavaScriptDialogOpts) error {
	return d.Client.Call("Page.handleJavaScriptDialog", opts, nil)
}

type GetAppManifestResult struct {
	// Manifest location.
	URL string `json:"url"`

	Errors []*AppManifestError `json:"errors"`

	// Manifest content. (optional)
	Data string `json:"data"`
}

// (experimental)
func (d *Domain) GetAppManifest() (*GetAppManifestResult, error) {
	var result GetAppManifestResult
	err := d.Client.Call("Page.getAppManifest", nil, &result)
	return &result, err
}

// (experimental)
func (d *Domain) RequestAppBanner() error {
	return d.Client.Call("Page.requestAppBanner", nil, nil)
}

type SetControlNavigationsOpts struct {
	Enabled bool `json:"enabled"`
}

// Toggles navigation throttling which allows programatic control over navigation and redirect response. (experimental)
func (d *Domain) SetControlNavigations(opts *SetControlNavigationsOpts) error {
	return d.Client.Call("Page.setControlNavigations", opts, nil)
}

type ProcessNavigationOpts struct {
	Response *NavigationResponse `json:"response"`

	NavigationId int `json:"navigationId"`
}

// Should be sent in response to a navigationRequested or a redirectRequested event, telling the browser how to handle the navigation. (experimental)
func (d *Domain) ProcessNavigation(opts *ProcessNavigationOpts) error {
	return d.Client.Call("Page.processNavigation", opts, nil)
}

type GetLayoutMetricsResult struct {
	// Metrics relating to the layout viewport.
	LayoutViewport *LayoutViewport `json:"layoutViewport"`

	// Metrics relating to the visual viewport.
	VisualViewport *VisualViewport `json:"visualViewport"`

	// Size of scrollable area.
	ContentSize interface{} `json:"contentSize"`
}

// Returns metrics relating to the layouting of the page, such as viewport bounds/scale. (experimental)
func (d *Domain) GetLayoutMetrics() (*GetLayoutMetricsResult, error) {
	var result GetLayoutMetricsResult
	err := d.Client.Call("Page.getLayoutMetrics", nil, &result)
	return &result, err
}

type CreateIsolatedWorldOpts struct {
	// Id of the frame in which the isolated world should be created.
	FrameId *FrameId `json:"frameId"`

	// An optional name which is reported in the Execution Context. (optional)
	WorldName string `json:"worldName,omitempty"`

	// Whether or not universal access should be granted to the isolated world. This is a powerful option, use with caution. (optional)
	GrantUniveralAccess bool `json:"grantUniveralAccess,omitempty"`
}

// Creates an isolated world for the given frame. (experimental)
func (d *Domain) CreateIsolatedWorld(opts *CreateIsolatedWorldOpts) error {
	return d.Client.Call("Page.createIsolatedWorld", opts, nil)
}

type DomContentEventFiredEvent struct {
	Timestamp float64 `json:"timestamp"`
}

func (d *Domain) OnDomContentEventFired(listener func(*DomContentEventFiredEvent)) {
	d.Client.AddListener("Page.domContentEventFired", func(params json.RawMessage) {
		var event DomContentEventFiredEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type LoadEventFiredEvent struct {
	Timestamp float64 `json:"timestamp"`
}

func (d *Domain) OnLoadEventFired(listener func(*LoadEventFiredEvent)) {
	d.Client.AddListener("Page.loadEventFired", func(params json.RawMessage) {
		var event LoadEventFiredEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type FrameAttachedEvent struct {
	// Id of the frame that has been attached.
	FrameId *FrameId `json:"frameId"`

	// Parent frame identifier.
	ParentFrameId *FrameId `json:"parentFrameId"`

	// JavaScript stack trace of when frame was attached, only set if frame initiated from script. (optional, experimental)
	Stack interface{} `json:"stack"`
}

// Fired when frame has been attached to its parent.
func (d *Domain) OnFrameAttached(listener func(*FrameAttachedEvent)) {
	d.Client.AddListener("Page.frameAttached", func(params json.RawMessage) {
		var event FrameAttachedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type FrameNavigatedEvent struct {
	// Frame object.
	Frame *Frame `json:"frame"`
}

// Fired once navigation of the frame has completed. Frame is now associated with the new loader.
func (d *Domain) OnFrameNavigated(listener func(*FrameNavigatedEvent)) {
	d.Client.AddListener("Page.frameNavigated", func(params json.RawMessage) {
		var event FrameNavigatedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type FrameDetachedEvent struct {
	// Id of the frame that has been detached.
	FrameId *FrameId `json:"frameId"`
}

// Fired when frame has been detached from its parent.
func (d *Domain) OnFrameDetached(listener func(*FrameDetachedEvent)) {
	d.Client.AddListener("Page.frameDetached", func(params json.RawMessage) {
		var event FrameDetachedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type FrameStartedLoadingEvent struct {
	// Id of the frame that has started loading.
	FrameId *FrameId `json:"frameId"`
}

// Fired when frame has started loading. (experimental)
func (d *Domain) OnFrameStartedLoading(listener func(*FrameStartedLoadingEvent)) {
	d.Client.AddListener("Page.frameStartedLoading", func(params json.RawMessage) {
		var event FrameStartedLoadingEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type FrameStoppedLoadingEvent struct {
	// Id of the frame that has stopped loading.
	FrameId *FrameId `json:"frameId"`
}

// Fired when frame has stopped loading. (experimental)
func (d *Domain) OnFrameStoppedLoading(listener func(*FrameStoppedLoadingEvent)) {
	d.Client.AddListener("Page.frameStoppedLoading", func(params json.RawMessage) {
		var event FrameStoppedLoadingEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type FrameScheduledNavigationEvent struct {
	// Id of the frame that has scheduled a navigation.
	FrameId *FrameId `json:"frameId"`

	// Delay (in seconds) until the navigation is scheduled to begin. The navigation is not guaranteed to start.
	Delay float64 `json:"delay"`
}

// Fired when frame schedules a potential navigation. (experimental)
func (d *Domain) OnFrameScheduledNavigation(listener func(*FrameScheduledNavigationEvent)) {
	d.Client.AddListener("Page.frameScheduledNavigation", func(params json.RawMessage) {
		var event FrameScheduledNavigationEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type FrameClearedScheduledNavigationEvent struct {
	// Id of the frame that has cleared its scheduled navigation.
	FrameId *FrameId `json:"frameId"`
}

// Fired when frame no longer has a scheduled navigation. (experimental)
func (d *Domain) OnFrameClearedScheduledNavigation(listener func(*FrameClearedScheduledNavigationEvent)) {
	d.Client.AddListener("Page.frameClearedScheduledNavigation", func(params json.RawMessage) {
		var event FrameClearedScheduledNavigationEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type FrameResizedEvent struct {
}

// (experimental)
func (d *Domain) OnFrameResized(listener func(*FrameResizedEvent)) {
	d.Client.AddListener("Page.frameResized", func(params json.RawMessage) {
		var event FrameResizedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type JavascriptDialogOpeningEvent struct {
	// Message that will be displayed by the dialog.
	Message string `json:"message"`

	// Dialog type.
	Type *DialogType `json:"type"`
}

// Fired when a JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload) is about to open.
func (d *Domain) OnJavascriptDialogOpening(listener func(*JavascriptDialogOpeningEvent)) {
	d.Client.AddListener("Page.javascriptDialogOpening", func(params json.RawMessage) {
		var event JavascriptDialogOpeningEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type JavascriptDialogClosedEvent struct {
	// Whether dialog was confirmed.
	Result bool `json:"result"`
}

// Fired when a JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload) has been closed.
func (d *Domain) OnJavascriptDialogClosed(listener func(*JavascriptDialogClosedEvent)) {
	d.Client.AddListener("Page.javascriptDialogClosed", func(params json.RawMessage) {
		var event JavascriptDialogClosedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type ScreencastFrameEvent struct {
	// Base64-encoded compressed image.
	Data string `json:"data"`

	// Screencast frame metadata.
	Metadata *ScreencastFrameMetadata `json:"metadata"`

	// Frame number.
	SessionId int `json:"sessionId"`
}

// Compressed image data requested by the <code>startScreencast</code>. (experimental)
func (d *Domain) OnScreencastFrame(listener func(*ScreencastFrameEvent)) {
	d.Client.AddListener("Page.screencastFrame", func(params json.RawMessage) {
		var event ScreencastFrameEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type ScreencastVisibilityChangedEvent struct {
	// True if the page is visible.
	Visible bool `json:"visible"`
}

// Fired when the page with currently enabled screencast was shown or hidden </code>. (experimental)
func (d *Domain) OnScreencastVisibilityChanged(listener func(*ScreencastVisibilityChangedEvent)) {
	d.Client.AddListener("Page.screencastVisibilityChanged", func(params json.RawMessage) {
		var event ScreencastVisibilityChangedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type InterstitialShownEvent struct {
}

// Fired when interstitial page was shown
func (d *Domain) OnInterstitialShown(listener func(*InterstitialShownEvent)) {
	d.Client.AddListener("Page.interstitialShown", func(params json.RawMessage) {
		var event InterstitialShownEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type InterstitialHiddenEvent struct {
}

// Fired when interstitial page was hidden
func (d *Domain) OnInterstitialHidden(listener func(*InterstitialHiddenEvent)) {
	d.Client.AddListener("Page.interstitialHidden", func(params json.RawMessage) {
		var event InterstitialHiddenEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type NavigationRequestedEvent struct {
	// Whether the navigation is taking place in the main frame or in a subframe.
	IsInMainFrame bool `json:"isInMainFrame"`

	// Whether the navigation has encountered a server redirect or not.
	IsRedirect bool `json:"isRedirect"`

	NavigationId int `json:"navigationId"`

	// URL of requested navigation.
	URL string `json:"url"`
}

// Fired when a navigation is started if navigation throttles are enabled.  The navigation will be deferred until processNavigation is called.
func (d *Domain) OnNavigationRequested(listener func(*NavigationRequestedEvent)) {
	d.Client.AddListener("Page.navigationRequested", func(params json.RawMessage) {
		var event NavigationRequestedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}
