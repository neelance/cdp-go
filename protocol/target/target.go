// Supports additional targets discovery and allows to attach to them. (experimental)
package target

import (
	"encoding/json"
	"log"

	"github.com/neelance/cdp-go/rpc"
)

// Supports additional targets discovery and allows to attach to them. (experimental)
type Domain struct {
	Client *rpc.Client
}

type TargetID string

type BrowserContextID string

type TargetInfo struct {
	TargetId TargetID `json:"targetId"`

	Type string `json:"type"`

	Title string `json:"title"`

	URL string `json:"url"`
}

type RemoteLocation struct {
	Host string `json:"host"`

	Port int `json:"port"`
}

type SetDiscoverTargetsOpts struct {
	// Whether to discover available targets.
	Discover bool `json:"discover"`
}

// Controls whether to discover available targets and notify via <code>targetCreated/targetDestroyed</code> events.
func (d *Domain) SetDiscoverTargets(opts *SetDiscoverTargetsOpts) error {
	return d.Client.Call("Target.setDiscoverTargets", opts, nil)
}

type SetAutoAttachOpts struct {
	// Whether to auto-attach to related targets.
	AutoAttach bool `json:"autoAttach"`

	// Whether to pause new targets when attaching to them. Use <code>Runtime.runIfWaitingForDebugger</code> to run paused targets.
	WaitForDebuggerOnStart bool `json:"waitForDebuggerOnStart"`
}

// Controls whether to automatically attach to new targets which are considered to be related to this one. When turned on, attaches to all existing related targets as well. When turned off, automatically detaches from all currently attached targets.
func (d *Domain) SetAutoAttach(opts *SetAutoAttachOpts) error {
	return d.Client.Call("Target.setAutoAttach", opts, nil)
}

type SetAttachToFramesOpts struct {
	// Whether to attach to frames.
	Value bool `json:"value"`
}

func (d *Domain) SetAttachToFrames(opts *SetAttachToFramesOpts) error {
	return d.Client.Call("Target.setAttachToFrames", opts, nil)
}

type SetRemoteLocationsOpts struct {
	// List of remote locations.
	Locations []*RemoteLocation `json:"locations"`
}

// Enables target discovery for the specified locations, when <code>setDiscoverTargets</code> was set to <code>true</code>.
func (d *Domain) SetRemoteLocations(opts *SetRemoteLocationsOpts) error {
	return d.Client.Call("Target.setRemoteLocations", opts, nil)
}

type SendMessageToTargetOpts struct {
	TargetId TargetID `json:"targetId"`

	Message string `json:"message"`
}

// Sends protocol message to the target with given id.
func (d *Domain) SendMessageToTarget(opts *SendMessageToTargetOpts) error {
	return d.Client.Call("Target.sendMessageToTarget", opts, nil)
}

type GetTargetInfoOpts struct {
	TargetId TargetID `json:"targetId"`
}

type GetTargetInfoResult struct {
	TargetInfo *TargetInfo `json:"targetInfo"`
}

// Returns information about a target.
func (d *Domain) GetTargetInfo(opts *GetTargetInfoOpts) (*GetTargetInfoResult, error) {
	var result GetTargetInfoResult
	err := d.Client.Call("Target.getTargetInfo", opts, &result)
	return &result, err
}

type ActivateTargetOpts struct {
	TargetId TargetID `json:"targetId"`
}

// Activates (focuses) the target.
func (d *Domain) ActivateTarget(opts *ActivateTargetOpts) error {
	return d.Client.Call("Target.activateTarget", opts, nil)
}

type CloseTargetOpts struct {
	TargetId TargetID `json:"targetId"`
}

type CloseTargetResult struct {
	Success bool `json:"success"`
}

// Closes the target. If the target is a page that gets closed too.
func (d *Domain) CloseTarget(opts *CloseTargetOpts) (*CloseTargetResult, error) {
	var result CloseTargetResult
	err := d.Client.Call("Target.closeTarget", opts, &result)
	return &result, err
}

type AttachToTargetOpts struct {
	TargetId TargetID `json:"targetId"`
}

type AttachToTargetResult struct {
	// Whether attach succeeded.
	Success bool `json:"success"`
}

// Attaches to the target with given id.
func (d *Domain) AttachToTarget(opts *AttachToTargetOpts) (*AttachToTargetResult, error) {
	var result AttachToTargetResult
	err := d.Client.Call("Target.attachToTarget", opts, &result)
	return &result, err
}

type DetachFromTargetOpts struct {
	TargetId TargetID `json:"targetId"`
}

// Detaches from the target with given id.
func (d *Domain) DetachFromTarget(opts *DetachFromTargetOpts) error {
	return d.Client.Call("Target.detachFromTarget", opts, nil)
}

type CreateBrowserContextResult struct {
	// The id of the context created.
	BrowserContextId BrowserContextID `json:"browserContextId"`
}

// Creates a new empty BrowserContext. Similar to an incognito profile but you can have more than one.
func (d *Domain) CreateBrowserContext() (*CreateBrowserContextResult, error) {
	var result CreateBrowserContextResult
	err := d.Client.Call("Target.createBrowserContext", nil, &result)
	return &result, err
}

type DisposeBrowserContextOpts struct {
	BrowserContextId BrowserContextID `json:"browserContextId"`
}

type DisposeBrowserContextResult struct {
	Success bool `json:"success"`
}

// Deletes a BrowserContext, will fail of any open page uses it.
func (d *Domain) DisposeBrowserContext(opts *DisposeBrowserContextOpts) (*DisposeBrowserContextResult, error) {
	var result DisposeBrowserContextResult
	err := d.Client.Call("Target.disposeBrowserContext", opts, &result)
	return &result, err
}

type CreateTargetOpts struct {
	// The initial URL the page will be navigated to.
	URL string `json:"url"`

	// Frame width in DIP (headless chrome only). (optional)
	Width int `json:"width,omitempty"`

	// Frame height in DIP (headless chrome only). (optional)
	Height int `json:"height,omitempty"`

	// The browser context to create the page in (headless chrome only). (optional)
	BrowserContextId BrowserContextID `json:"browserContextId,omitempty"`
}

type CreateTargetResult struct {
	// The id of the page opened.
	TargetId TargetID `json:"targetId"`
}

// Creates a new page.
func (d *Domain) CreateTarget(opts *CreateTargetOpts) (*CreateTargetResult, error) {
	var result CreateTargetResult
	err := d.Client.Call("Target.createTarget", opts, &result)
	return &result, err
}

type GetTargetsResult struct {
	// The list of targets.
	TargetInfos []*TargetInfo `json:"targetInfos"`
}

// Retrieves a list of available targets.
func (d *Domain) GetTargets() (*GetTargetsResult, error) {
	var result GetTargetsResult
	err := d.Client.Call("Target.getTargets", nil, &result)
	return &result, err
}

type TargetCreatedEvent struct {
	TargetInfo *TargetInfo `json:"targetInfo"`
}

// Issued when a possible inspection target is created.
func (d *Domain) OnTargetCreated(listener func(*TargetCreatedEvent)) {
	d.Client.AddListener("Target.targetCreated", func(params json.RawMessage) {
		var event TargetCreatedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type TargetDestroyedEvent struct {
	TargetId TargetID `json:"targetId"`
}

// Issued when a target is destroyed.
func (d *Domain) OnTargetDestroyed(listener func(*TargetDestroyedEvent)) {
	d.Client.AddListener("Target.targetDestroyed", func(params json.RawMessage) {
		var event TargetDestroyedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type AttachedToTargetEvent struct {
	TargetInfo *TargetInfo `json:"targetInfo"`

	WaitingForDebugger bool `json:"waitingForDebugger"`
}

// Issued when attached to target because of auto-attach or <code>attachToTarget</code> command.
func (d *Domain) OnAttachedToTarget(listener func(*AttachedToTargetEvent)) {
	d.Client.AddListener("Target.attachedToTarget", func(params json.RawMessage) {
		var event AttachedToTargetEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type DetachedFromTargetEvent struct {
	TargetId TargetID `json:"targetId"`
}

// Issued when detached from target for any reason (including <code>detachFromTarget</code> command).
func (d *Domain) OnDetachedFromTarget(listener func(*DetachedFromTargetEvent)) {
	d.Client.AddListener("Target.detachedFromTarget", func(params json.RawMessage) {
		var event DetachedFromTargetEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type ReceivedMessageFromTargetEvent struct {
	TargetId TargetID `json:"targetId"`

	Message string `json:"message"`
}

// Notifies about new protocol message from attached target.
func (d *Domain) OnReceivedMessageFromTarget(listener func(*ReceivedMessageFromTargetEvent)) {
	d.Client.AddListener("Target.receivedMessageFromTarget", func(params json.RawMessage) {
		var event ReceivedMessageFromTargetEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}
