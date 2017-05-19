// (experimental)
package serviceworker

import (
	"encoding/json"
	"log"

	"github.com/neelance/cdp-go/rpc"
)

// (experimental)
type Domain struct {
	Client *rpc.Client
}

// ServiceWorker registration.

type ServiceWorkerRegistration struct {
	RegistrationId string `json:"registrationId"`

	ScopeURL string `json:"scopeURL"`

	IsDeleted bool `json:"isDeleted"`
}

type ServiceWorkerVersionRunningStatus string

type ServiceWorkerVersionStatus string

// ServiceWorker version.

type ServiceWorkerVersion struct {
	VersionId string `json:"versionId"`

	RegistrationId string `json:"registrationId"`

	ScriptURL string `json:"scriptURL"`

	RunningStatus ServiceWorkerVersionRunningStatus `json:"runningStatus"`

	Status ServiceWorkerVersionStatus `json:"status"`

	// The Last-Modified header value of the main script. (optional)
	ScriptLastModified float64 `json:"scriptLastModified,omitempty"`

	// The time at which the response headers of the main script were received from the server.  For cached script it is the last time the cache entry was validated. (optional)
	ScriptResponseTime float64 `json:"scriptResponseTime,omitempty"`

	// (optional)
	ControlledClients []interface{} `json:"controlledClients,omitempty"`

	// (optional)
	TargetId interface{} `json:"targetId,omitempty"`
}

// ServiceWorker error message.

type ServiceWorkerErrorMessage struct {
	ErrorMessage string `json:"errorMessage"`

	RegistrationId string `json:"registrationId"`

	VersionId string `json:"versionId"`

	SourceURL string `json:"sourceURL"`

	LineNumber int `json:"lineNumber"`

	ColumnNumber int `json:"columnNumber"`
}

func (d *Domain) Enable() error {
	return d.Client.Call("ServiceWorker.enable", nil, nil)
}

func (d *Domain) Disable() error {
	return d.Client.Call("ServiceWorker.disable", nil, nil)
}

type UnregisterOpts struct {
	ScopeURL string `json:"scopeURL"`
}

func (d *Domain) Unregister(opts *UnregisterOpts) error {
	return d.Client.Call("ServiceWorker.unregister", opts, nil)
}

type UpdateRegistrationOpts struct {
	ScopeURL string `json:"scopeURL"`
}

func (d *Domain) UpdateRegistration(opts *UpdateRegistrationOpts) error {
	return d.Client.Call("ServiceWorker.updateRegistration", opts, nil)
}

type StartWorkerOpts struct {
	ScopeURL string `json:"scopeURL"`
}

func (d *Domain) StartWorker(opts *StartWorkerOpts) error {
	return d.Client.Call("ServiceWorker.startWorker", opts, nil)
}

type SkipWaitingOpts struct {
	ScopeURL string `json:"scopeURL"`
}

func (d *Domain) SkipWaiting(opts *SkipWaitingOpts) error {
	return d.Client.Call("ServiceWorker.skipWaiting", opts, nil)
}

type StopWorkerOpts struct {
	VersionId string `json:"versionId"`
}

func (d *Domain) StopWorker(opts *StopWorkerOpts) error {
	return d.Client.Call("ServiceWorker.stopWorker", opts, nil)
}

type InspectWorkerOpts struct {
	VersionId string `json:"versionId"`
}

func (d *Domain) InspectWorker(opts *InspectWorkerOpts) error {
	return d.Client.Call("ServiceWorker.inspectWorker", opts, nil)
}

type SetForceUpdateOnPageLoadOpts struct {
	ForceUpdateOnPageLoad bool `json:"forceUpdateOnPageLoad"`
}

func (d *Domain) SetForceUpdateOnPageLoad(opts *SetForceUpdateOnPageLoadOpts) error {
	return d.Client.Call("ServiceWorker.setForceUpdateOnPageLoad", opts, nil)
}

type DeliverPushMessageOpts struct {
	Origin string `json:"origin"`

	RegistrationId string `json:"registrationId"`

	Data string `json:"data"`
}

func (d *Domain) DeliverPushMessage(opts *DeliverPushMessageOpts) error {
	return d.Client.Call("ServiceWorker.deliverPushMessage", opts, nil)
}

type DispatchSyncEventOpts struct {
	Origin string `json:"origin"`

	RegistrationId string `json:"registrationId"`

	Tag string `json:"tag"`

	LastChance bool `json:"lastChance"`
}

func (d *Domain) DispatchSyncEvent(opts *DispatchSyncEventOpts) error {
	return d.Client.Call("ServiceWorker.dispatchSyncEvent", opts, nil)
}

type WorkerRegistrationUpdatedEvent struct {
	Registrations []*ServiceWorkerRegistration `json:"registrations"`
}

func (d *Domain) OnWorkerRegistrationUpdated(listener func(*WorkerRegistrationUpdatedEvent)) {
	d.Client.AddListener("ServiceWorker.workerRegistrationUpdated", func(params json.RawMessage) {
		var event WorkerRegistrationUpdatedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type WorkerVersionUpdatedEvent struct {
	Versions []*ServiceWorkerVersion `json:"versions"`
}

func (d *Domain) OnWorkerVersionUpdated(listener func(*WorkerVersionUpdatedEvent)) {
	d.Client.AddListener("ServiceWorker.workerVersionUpdated", func(params json.RawMessage) {
		var event WorkerVersionUpdatedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type WorkerErrorReportedEvent struct {
	ErrorMessage *ServiceWorkerErrorMessage `json:"errorMessage"`
}

func (d *Domain) OnWorkerErrorReported(listener func(*WorkerErrorReportedEvent)) {
	d.Client.AddListener("ServiceWorker.workerErrorReported", func(params json.RawMessage) {
		var event WorkerErrorReportedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}
