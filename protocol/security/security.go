// Security (experimental)
package security

import (
	"github.com/neelance/cdp-go/rpc"
)

// Security (experimental)
type Domain struct {
	Client *rpc.Client
}

// An internal certificate ID value.

type CertificateId int

// The security level of a page or resource.

type SecurityState string

// An explanation of an factor contributing to the security state.

type SecurityStateExplanation struct {
	// Security state representing the severity of the factor being explained.
	SecurityState SecurityState `json:"securityState"`

	// Short phrase describing the type of factor.
	Summary string `json:"summary"`

	// Full text explanation of the factor.
	Description string `json:"description"`

	// True if the page has a certificate.
	HasCertificate bool `json:"hasCertificate"`
}

// Information about insecure content on the page.

type InsecureContentStatus struct {
	// True if the page was loaded over HTTPS and ran mixed (HTTP) content such as scripts.
	RanMixedContent bool `json:"ranMixedContent"`

	// True if the page was loaded over HTTPS and displayed mixed (HTTP) content such as images.
	DisplayedMixedContent bool `json:"displayedMixedContent"`

	// True if the page was loaded over HTTPS and contained a form targeting an insecure url.
	ContainedMixedForm bool `json:"containedMixedForm"`

	// True if the page was loaded over HTTPS without certificate errors, and ran content such as scripts that were loaded with certificate errors.
	RanContentWithCertErrors bool `json:"ranContentWithCertErrors"`

	// True if the page was loaded over HTTPS without certificate errors, and displayed content such as images that were loaded with certificate errors.
	DisplayedContentWithCertErrors bool `json:"displayedContentWithCertErrors"`

	// Security state representing a page that ran insecure content.
	RanInsecureContentStyle SecurityState `json:"ranInsecureContentStyle"`

	// Security state representing a page that displayed insecure content.
	DisplayedInsecureContentStyle SecurityState `json:"displayedInsecureContentStyle"`
}

// The action to take when a certificate error occurs. continue will continue processing the request and cancel will cancel the request.

type CertificateErrorAction string

// Enables tracking security state changes.
type EnableRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) Enable() *EnableRequest {
	return &EnableRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Enables tracking security state changes.
func (r *EnableRequest) Do() error {
	return r.client.Call("Security.enable", r.opts, nil)
}

// Disables tracking security state changes.
type DisableRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) Disable() *DisableRequest {
	return &DisableRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Disables tracking security state changes.
func (r *DisableRequest) Do() error {
	return r.client.Call("Security.disable", r.opts, nil)
}

// Displays native dialog with the certificate details.
type ShowCertificateViewerRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) ShowCertificateViewer() *ShowCertificateViewerRequest {
	return &ShowCertificateViewerRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Displays native dialog with the certificate details.
func (r *ShowCertificateViewerRequest) Do() error {
	return r.client.Call("Security.showCertificateViewer", r.opts, nil)
}

// Handles a certificate error that fired a certificateError event.
type HandleCertificateErrorRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) HandleCertificateError() *HandleCertificateErrorRequest {
	return &HandleCertificateErrorRequest{opts: make(map[string]interface{}), client: d.Client}
}

// The ID of the event.
func (r *HandleCertificateErrorRequest) EventId(v int) *HandleCertificateErrorRequest {
	r.opts["eventId"] = v
	return r
}

// The action to take on the certificate error.
func (r *HandleCertificateErrorRequest) Action(v CertificateErrorAction) *HandleCertificateErrorRequest {
	r.opts["action"] = v
	return r
}

// Handles a certificate error that fired a certificateError event.
func (r *HandleCertificateErrorRequest) Do() error {
	return r.client.Call("Security.handleCertificateError", r.opts, nil)
}

// Enable/disable overriding certificate errors. If enabled, all certificate error events need to be handled by the DevTools client and should be answered with handleCertificateError commands.
type SetOverrideCertificateErrorsRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Domain) SetOverrideCertificateErrors() *SetOverrideCertificateErrorsRequest {
	return &SetOverrideCertificateErrorsRequest{opts: make(map[string]interface{}), client: d.Client}
}

// If true, certificate errors will be overridden.
func (r *SetOverrideCertificateErrorsRequest) Override(v bool) *SetOverrideCertificateErrorsRequest {
	r.opts["override"] = v
	return r
}

// Enable/disable overriding certificate errors. If enabled, all certificate error events need to be handled by the DevTools client and should be answered with handleCertificateError commands.
func (r *SetOverrideCertificateErrorsRequest) Do() error {
	return r.client.Call("Security.setOverrideCertificateErrors", r.opts, nil)
}

func init() {
	rpc.EventTypes["Security.securityStateChanged"] = func() interface{} { return new(SecurityStateChangedEvent) }
	rpc.EventTypes["Security.certificateError"] = func() interface{} { return new(CertificateErrorEvent) }
}

// The security state of the page changed.
type SecurityStateChangedEvent struct {
	// Security state.
	SecurityState SecurityState `json:"securityState"`

	// True if the page was loaded over cryptographic transport such as HTTPS.
	SchemeIsCryptographic bool `json:"schemeIsCryptographic"`

	// List of explanations for the security state. If the overall security state is `insecure` or `warning`, at least one corresponding explanation should be included.
	Explanations []*SecurityStateExplanation `json:"explanations"`

	// Information about insecure content on the page.
	InsecureContentStatus *InsecureContentStatus `json:"insecureContentStatus"`

	// Overrides user-visible description of the state. (optional)
	Summary string `json:"summary"`
}

// There is a certificate error. If overriding certificate errors is enabled, then it should be handled with the handleCertificateError command. Note: this event does not fire if the certificate error has been allowed internally.
type CertificateErrorEvent struct {
	// The ID of the event.
	EventId int `json:"eventId"`

	// The type of the error.
	ErrorType string `json:"errorType"`

	// The url that was requested.
	RequestURL string `json:"requestURL"`
}
