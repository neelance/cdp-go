// Security (experimental)
package security

import (
	"encoding/json"
	"log"

	"github.com/neelance/cdp-go/rpc"
)

// Security (experimental)
type Domain struct {
	Client *rpc.Client
}

// An internal certificate ID value.
type CertificateId interface{}

// The security level of a page or resource.
type SecurityState interface{}

// An explanation of an factor contributing to the security state.
type SecurityStateExplanation interface{}

// Information about insecure content on the page.
type InsecureContentStatus interface{}

// The action to take when a certificate error occurs. continue will continue processing the request and cancel will cancel the request.
type CertificateErrorAction interface{}

// Enables tracking security state changes.
func (d *Domain) Enable() error {
	return d.Client.Call("Security.enable", nil, nil)
}

// Disables tracking security state changes.
func (d *Domain) Disable() error {
	return d.Client.Call("Security.disable", nil, nil)
}

// Displays native dialog with the certificate details.
func (d *Domain) ShowCertificateViewer() error {
	return d.Client.Call("Security.showCertificateViewer", nil, nil)
}

type HandleCertificateErrorOpts struct {
	// The ID of the event.
	EventId int `json:"eventId"`

	// The action to take on the certificate error.
	Action CertificateErrorAction `json:"action"`
}

// Handles a certificate error that fired a certificateError event.
func (d *Domain) HandleCertificateError(opts *HandleCertificateErrorOpts) error {
	return d.Client.Call("Security.handleCertificateError", opts, nil)
}

type SetOverrideCertificateErrorsOpts struct {
	// If true, certificate errors will be overridden.
	Override bool `json:"override"`
}

// Enable/disable overriding certificate errors. If enabled, all certificate error events need to be handled by the DevTools client and should be answered with handleCertificateError commands.
func (d *Domain) SetOverrideCertificateErrors(opts *SetOverrideCertificateErrorsOpts) error {
	return d.Client.Call("Security.setOverrideCertificateErrors", opts, nil)
}

type SecurityStateChangedEvent struct {
	// Security state.
	SecurityState SecurityState `json:"securityState"`

	// True if the page was loaded over cryptographic transport such as HTTPS.
	SchemeIsCryptographic bool `json:"schemeIsCryptographic"`

	// List of explanations for the security state. If the overall security state is `insecure` or `warning`, at least one corresponding explanation should be included.
	Explanations []SecurityStateExplanation `json:"explanations"`

	// Information about insecure content on the page.
	InsecureContentStatus InsecureContentStatus `json:"insecureContentStatus"`

	// Overrides user-visible description of the state. (optional)
	Summary string `json:"summary"`
}

// The security state of the page changed.
func (d *Domain) OnSecurityStateChanged(listener func(*SecurityStateChangedEvent)) {
	d.Client.AddListener("Security.securityStateChanged", func(params json.RawMessage) {
		var event SecurityStateChangedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type CertificateErrorEvent struct {
	// The ID of the event.
	EventId int `json:"eventId"`

	// The type of the error.
	ErrorType string `json:"errorType"`

	// The url that was requested.
	RequestURL string `json:"requestURL"`
}

// There is a certificate error. If overriding certificate errors is enabled, then it should be handled with the handleCertificateError command. Note: this event does not fire if the certificate error has been allowed internally.
func (d *Domain) OnCertificateError(listener func(*CertificateErrorEvent)) {
	d.Client.AddListener("Security.certificateError", func(params json.RawMessage) {
		var event CertificateErrorEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}
