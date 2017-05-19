// Network domain allows tracking network activities of the page. It exposes information about http, file, data and other requests and responses, their headers, bodies, timing, etc.
package network

import (
	"encoding/json"
	"log"

	"github.com/neelance/cdp-go/rpc"
)

// Network domain allows tracking network activities of the page. It exposes information about http, file, data and other requests and responses, their headers, bodies, timing, etc.
type Domain struct {
	Client *rpc.Client
}

// Unique loader identifier.

type LoaderId string

// Unique request identifier.

type RequestId string

// Number of seconds since epoch.

type Timestamp float64

// Request / response headers as keys / values of JSON object.

type Headers struct {
}

// Loading priority of a resource request.

type ConnectionType string

// Represents the cookie's 'SameSite' status: https://tools.ietf.org/html/draft-west-first-party-cookies

type CookieSameSite string

// Timing information for the request.

type ResourceTiming struct {
	// Timing's requestTime is a baseline in seconds, while the other numbers are ticks in milliseconds relatively to this requestTime.
	RequestTime float64 `json:"requestTime"`

	// Started resolving proxy.
	ProxyStart float64 `json:"proxyStart"`

	// Finished resolving proxy.
	ProxyEnd float64 `json:"proxyEnd"`

	// Started DNS address resolve.
	DnsStart float64 `json:"dnsStart"`

	// Finished DNS address resolve.
	DnsEnd float64 `json:"dnsEnd"`

	// Started connecting to the remote host.
	ConnectStart float64 `json:"connectStart"`

	// Connected to the remote host.
	ConnectEnd float64 `json:"connectEnd"`

	// Started SSL handshake.
	SslStart float64 `json:"sslStart"`

	// Finished SSL handshake.
	SslEnd float64 `json:"sslEnd"`

	// Started running ServiceWorker.
	WorkerStart float64 `json:"workerStart"`

	// Finished Starting ServiceWorker.
	WorkerReady float64 `json:"workerReady"`

	// Started sending request.
	SendStart float64 `json:"sendStart"`

	// Finished sending request.
	SendEnd float64 `json:"sendEnd"`

	// Time the server started pushing request.
	PushStart float64 `json:"pushStart"`

	// Time the server finished pushing request.
	PushEnd float64 `json:"pushEnd"`

	// Finished receiving response headers.
	ReceiveHeadersEnd float64 `json:"receiveHeadersEnd"`
}

// Loading priority of a resource request.

type ResourcePriority string

// HTTP request data.

type Request struct {
	// Request URL.
	URL string `json:"url"`

	// HTTP request method.
	Method string `json:"method"`

	// HTTP request headers.
	Headers *Headers `json:"headers"`

	// HTTP POST request data. (optional)
	PostData string `json:"postData,omitempty"`

	// The mixed content status of the request, as defined in http://www.w3.org/TR/mixed-content/ (optional)
	MixedContentType string `json:"mixedContentType,omitempty"`

	// Priority of the resource request at the time request is sent.
	InitialPriority *ResourcePriority `json:"initialPriority"`

	// The referrer policy of the request, as defined in https://www.w3.org/TR/referrer-policy/
	ReferrerPolicy string `json:"referrerPolicy"`

	// Whether is loaded via link preload. (optional)
	IsLinkPreload bool `json:"isLinkPreload,omitempty"`
}

// Details of a signed certificate timestamp (SCT).

type SignedCertificateTimestamp struct {
	// Validation status.
	Status string `json:"status"`

	// Origin.
	Origin string `json:"origin"`

	// Log name / description.
	LogDescription string `json:"logDescription"`

	// Log ID.
	LogId string `json:"logId"`

	// Issuance date.
	Timestamp *Timestamp `json:"timestamp"`

	// Hash algorithm.
	HashAlgorithm string `json:"hashAlgorithm"`

	// Signature algorithm.
	SignatureAlgorithm string `json:"signatureAlgorithm"`

	// Signature data.
	SignatureData string `json:"signatureData"`
}

// Security details about a request.

type SecurityDetails struct {
	// Protocol name (e.g. "TLS 1.2" or "QUIC").
	Protocol string `json:"protocol"`

	// Key Exchange used by the connection, or the empty string if not applicable.
	KeyExchange string `json:"keyExchange"`

	// (EC)DH group used by the connection, if applicable. (optional)
	KeyExchangeGroup string `json:"keyExchangeGroup,omitempty"`

	// Cipher name.
	Cipher string `json:"cipher"`

	// TLS MAC. Note that AEAD ciphers do not have separate MACs. (optional)
	Mac string `json:"mac,omitempty"`

	// Certificate ID value.
	CertificateId interface{} `json:"certificateId"`

	// Certificate subject name.
	SubjectName string `json:"subjectName"`

	// Subject Alternative Name (SAN) DNS names and IP addresses.
	SanList []string `json:"sanList"`

	// Name of the issuing CA.
	Issuer string `json:"issuer"`

	// Certificate valid from date.
	ValidFrom *Timestamp `json:"validFrom"`

	// Certificate valid to (expiration) date
	ValidTo *Timestamp `json:"validTo"`

	// List of signed certificate timestamps (SCTs).
	SignedCertificateTimestampList []*SignedCertificateTimestamp `json:"signedCertificateTimestampList"`
}

// The reason why request was blocked. (experimental)

type BlockedReason string

// HTTP response data.

type Response struct {
	// Response URL. This URL can be different from CachedResource.url in case of redirect.
	URL string `json:"url"`

	// HTTP response status code.
	Status float64 `json:"status"`

	// HTTP response status text.
	StatusText string `json:"statusText"`

	// HTTP response headers.
	Headers *Headers `json:"headers"`

	// HTTP response headers text. (optional)
	HeadersText string `json:"headersText,omitempty"`

	// Resource mimeType as determined by the browser.
	MimeType string `json:"mimeType"`

	// Refined HTTP request headers that were actually transmitted over the network. (optional)
	RequestHeaders *Headers `json:"requestHeaders,omitempty"`

	// HTTP request headers text. (optional)
	RequestHeadersText string `json:"requestHeadersText,omitempty"`

	// Specifies whether physical connection was actually reused for this request.
	ConnectionReused bool `json:"connectionReused"`

	// Physical connection id that was actually used for this request.
	ConnectionId float64 `json:"connectionId"`

	// Remote IP address. (optional, experimental)
	RemoteIPAddress string `json:"remoteIPAddress,omitempty"`

	// Remote port. (optional, experimental)
	RemotePort int `json:"remotePort,omitempty"`

	// Specifies that the request was served from the disk cache. (optional)
	FromDiskCache bool `json:"fromDiskCache,omitempty"`

	// Specifies that the request was served from the ServiceWorker. (optional)
	FromServiceWorker bool `json:"fromServiceWorker,omitempty"`

	// Total number of bytes received for this request so far.
	EncodedDataLength float64 `json:"encodedDataLength"`

	// Timing information for the given request. (optional)
	Timing *ResourceTiming `json:"timing,omitempty"`

	// Protocol used to fetch this request. (optional)
	Protocol string `json:"protocol,omitempty"`

	// Security state of the request resource.
	SecurityState interface{} `json:"securityState"`

	// Security details for the request. (optional)
	SecurityDetails *SecurityDetails `json:"securityDetails,omitempty"`
}

// WebSocket request data. (experimental)

type WebSocketRequest struct {
	// HTTP request headers.
	Headers *Headers `json:"headers"`
}

// WebSocket response data. (experimental)

type WebSocketResponse struct {
	// HTTP response status code.
	Status float64 `json:"status"`

	// HTTP response status text.
	StatusText string `json:"statusText"`

	// HTTP response headers.
	Headers *Headers `json:"headers"`

	// HTTP response headers text. (optional)
	HeadersText string `json:"headersText,omitempty"`

	// HTTP request headers. (optional)
	RequestHeaders *Headers `json:"requestHeaders,omitempty"`

	// HTTP request headers text. (optional)
	RequestHeadersText string `json:"requestHeadersText,omitempty"`
}

// WebSocket frame data. (experimental)

type WebSocketFrame struct {
	// WebSocket frame opcode.
	Opcode float64 `json:"opcode"`

	// WebSocke frame mask.
	Mask bool `json:"mask"`

	// WebSocke frame payload data.
	PayloadData string `json:"payloadData"`
}

// Information about the cached resource.

type CachedResource struct {
	// Resource URL. This is the url of the original network request.
	URL string `json:"url"`

	// Type of this resource.
	Type interface{} `json:"type"`

	// Cached response data. (optional)
	Response *Response `json:"response,omitempty"`

	// Cached response body size.
	BodySize float64 `json:"bodySize"`
}

// Information about the request initiator.

type Initiator struct {
	// Type of this initiator.
	Type string `json:"type"`

	// Initiator JavaScript stack trace, set for Script only. (optional)
	Stack interface{} `json:"stack,omitempty"`

	// Initiator URL, set for Parser type only. (optional)
	URL string `json:"url,omitempty"`

	// Initiator line number, set for Parser type only (0-based). (optional)
	LineNumber float64 `json:"lineNumber,omitempty"`
}

// Cookie object (experimental)

type Cookie struct {
	// Cookie name.
	Name string `json:"name"`

	// Cookie value.
	Value string `json:"value"`

	// Cookie domain.
	Domain string `json:"domain"`

	// Cookie path.
	Path string `json:"path"`

	// Cookie expiration date as the number of seconds since the UNIX epoch.
	Expires float64 `json:"expires"`

	// Cookie size.
	Size int `json:"size"`

	// True if cookie is http-only.
	HttpOnly bool `json:"httpOnly"`

	// True if cookie is secure.
	Secure bool `json:"secure"`

	// True in case of session cookie.
	Session bool `json:"session"`

	// Cookie SameSite type. (optional)
	SameSite *CookieSameSite `json:"sameSite,omitempty"`
}

type EnableOpts struct {
	// Buffer size in bytes to use when preserving network payloads (XHRs, etc). (optional, experimental)
	MaxTotalBufferSize int `json:"maxTotalBufferSize,omitempty"`

	// Per-resource buffer size in bytes to use when preserving network payloads (XHRs, etc). (optional, experimental)
	MaxResourceBufferSize int `json:"maxResourceBufferSize,omitempty"`
}

// Enables network tracking, network events will now be delivered to the client.
func (d *Domain) Enable(opts *EnableOpts) error {
	return d.Client.Call("Network.enable", opts, nil)
}

// Disables network tracking, prevents network events from being sent to the client.
func (d *Domain) Disable() error {
	return d.Client.Call("Network.disable", nil, nil)
}

type SetUserAgentOverrideOpts struct {
	// User agent to use.
	UserAgent string `json:"userAgent"`
}

// Allows overriding user agent with the given string.
func (d *Domain) SetUserAgentOverride(opts *SetUserAgentOverrideOpts) error {
	return d.Client.Call("Network.setUserAgentOverride", opts, nil)
}

type SetExtraHTTPHeadersOpts struct {
	// Map with extra HTTP headers.
	Headers *Headers `json:"headers"`
}

// Specifies whether to always send extra HTTP headers with the requests from this page.
func (d *Domain) SetExtraHTTPHeaders(opts *SetExtraHTTPHeadersOpts) error {
	return d.Client.Call("Network.setExtraHTTPHeaders", opts, nil)
}

type GetResponseBodyOpts struct {
	// Identifier of the network request to get content for.
	RequestId *RequestId `json:"requestId"`
}

type GetResponseBodyResult struct {
	// Response body.
	Body string `json:"body"`

	// True, if content was sent as base64.
	Base64Encoded bool `json:"base64Encoded"`
}

// Returns content served for the given request.
func (d *Domain) GetResponseBody(opts *GetResponseBodyOpts) (*GetResponseBodyResult, error) {
	var result GetResponseBodyResult
	err := d.Client.Call("Network.getResponseBody", opts, &result)
	return &result, err
}

type SetBlockedURLsOpts struct {
	// URL patterns to block. Wildcards ('*') are allowed.
	Urls []string `json:"urls"`
}

// Blocks URLs from loading. (experimental)
func (d *Domain) SetBlockedURLs(opts *SetBlockedURLsOpts) error {
	return d.Client.Call("Network.setBlockedURLs", opts, nil)
}

type ReplayXHROpts struct {
	// Identifier of XHR to replay.
	RequestId *RequestId `json:"requestId"`
}

// This method sends a new XMLHttpRequest which is identical to the original one. The following parameters should be identical: method, url, async, request body, extra headers, withCredentials attribute, user, password. (experimental)
func (d *Domain) ReplayXHR(opts *ReplayXHROpts) error {
	return d.Client.Call("Network.replayXHR", opts, nil)
}

type CanClearBrowserCacheResult struct {
	// True if browser cache can be cleared.
	Result bool `json:"result"`
}

// Tells whether clearing browser cache is supported.
func (d *Domain) CanClearBrowserCache() (*CanClearBrowserCacheResult, error) {
	var result CanClearBrowserCacheResult
	err := d.Client.Call("Network.canClearBrowserCache", nil, &result)
	return &result, err
}

// Clears browser cache.
func (d *Domain) ClearBrowserCache() error {
	return d.Client.Call("Network.clearBrowserCache", nil, nil)
}

type CanClearBrowserCookiesResult struct {
	// True if browser cookies can be cleared.
	Result bool `json:"result"`
}

// Tells whether clearing browser cookies is supported.
func (d *Domain) CanClearBrowserCookies() (*CanClearBrowserCookiesResult, error) {
	var result CanClearBrowserCookiesResult
	err := d.Client.Call("Network.canClearBrowserCookies", nil, &result)
	return &result, err
}

// Clears browser cookies.
func (d *Domain) ClearBrowserCookies() error {
	return d.Client.Call("Network.clearBrowserCookies", nil, nil)
}

type GetCookiesOpts struct {
	// The list of URLs for which applicable cookies will be fetched (optional)
	Urls []string `json:"urls,omitempty"`
}

type GetCookiesResult struct {
	// Array of cookie objects.
	Cookies []*Cookie `json:"cookies"`
}

// Returns all browser cookies for the current URL. Depending on the backend support, will return detailed cookie information in the <code>cookies</code> field. (experimental)
func (d *Domain) GetCookies(opts *GetCookiesOpts) (*GetCookiesResult, error) {
	var result GetCookiesResult
	err := d.Client.Call("Network.getCookies", opts, &result)
	return &result, err
}

type GetAllCookiesResult struct {
	// Array of cookie objects.
	Cookies []*Cookie `json:"cookies"`
}

// Returns all browser cookies. Depending on the backend support, will return detailed cookie information in the <code>cookies</code> field. (experimental)
func (d *Domain) GetAllCookies() (*GetAllCookiesResult, error) {
	var result GetAllCookiesResult
	err := d.Client.Call("Network.getAllCookies", nil, &result)
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
	return d.Client.Call("Network.deleteCookie", opts, nil)
}

type SetCookieOpts struct {
	// The request-URI to associate with the setting of the cookie. This value can affect the default domain and path values of the created cookie.
	URL string `json:"url"`

	// The name of the cookie.
	Name string `json:"name"`

	// The value of the cookie.
	Value string `json:"value"`

	// If omitted, the cookie becomes a host-only cookie. (optional)
	Domain string `json:"domain,omitempty"`

	// Defaults to the path portion of the url parameter. (optional)
	Path string `json:"path,omitempty"`

	// Defaults ot false. (optional)
	Secure bool `json:"secure,omitempty"`

	// Defaults to false. (optional)
	HttpOnly bool `json:"httpOnly,omitempty"`

	// Defaults to browser default behavior. (optional)
	SameSite *CookieSameSite `json:"sameSite,omitempty"`

	// If omitted, the cookie becomes a session cookie. (optional)
	ExpirationDate *Timestamp `json:"expirationDate,omitempty"`
}

type SetCookieResult struct {
	// True if successfully set cookie.
	Success bool `json:"success"`
}

// Sets a cookie with the given cookie data; may overwrite equivalent cookies if they exist. (experimental)
func (d *Domain) SetCookie(opts *SetCookieOpts) (*SetCookieResult, error) {
	var result SetCookieResult
	err := d.Client.Call("Network.setCookie", opts, &result)
	return &result, err
}

type CanEmulateNetworkConditionsResult struct {
	// True if emulation of network conditions is supported.
	Result bool `json:"result"`
}

// Tells whether emulation of network conditions is supported. (experimental)
func (d *Domain) CanEmulateNetworkConditions() (*CanEmulateNetworkConditionsResult, error) {
	var result CanEmulateNetworkConditionsResult
	err := d.Client.Call("Network.canEmulateNetworkConditions", nil, &result)
	return &result, err
}

type EmulateNetworkConditionsOpts struct {
	// True to emulate internet disconnection.
	Offline bool `json:"offline"`

	// Additional latency (ms).
	Latency float64 `json:"latency"`

	// Maximal aggregated download throughput.
	DownloadThroughput float64 `json:"downloadThroughput"`

	// Maximal aggregated upload throughput.
	UploadThroughput float64 `json:"uploadThroughput"`

	// Connection type if known. (optional)
	ConnectionType *ConnectionType `json:"connectionType,omitempty"`
}

// Activates emulation of network conditions.
func (d *Domain) EmulateNetworkConditions(opts *EmulateNetworkConditionsOpts) error {
	return d.Client.Call("Network.emulateNetworkConditions", opts, nil)
}

type SetCacheDisabledOpts struct {
	// Cache disabled state.
	CacheDisabled bool `json:"cacheDisabled"`
}

// Toggles ignoring cache for each request. If <code>true</code>, cache will not be used.
func (d *Domain) SetCacheDisabled(opts *SetCacheDisabledOpts) error {
	return d.Client.Call("Network.setCacheDisabled", opts, nil)
}

type SetBypassServiceWorkerOpts struct {
	// Bypass service worker and load from network.
	Bypass bool `json:"bypass"`
}

// Toggles ignoring of service worker for each request. (experimental)
func (d *Domain) SetBypassServiceWorker(opts *SetBypassServiceWorkerOpts) error {
	return d.Client.Call("Network.setBypassServiceWorker", opts, nil)
}

type SetDataSizeLimitsForTestOpts struct {
	// Maximum total buffer size.
	MaxTotalSize int `json:"maxTotalSize"`

	// Maximum per-resource size.
	MaxResourceSize int `json:"maxResourceSize"`
}

// For testing. (experimental)
func (d *Domain) SetDataSizeLimitsForTest(opts *SetDataSizeLimitsForTestOpts) error {
	return d.Client.Call("Network.setDataSizeLimitsForTest", opts, nil)
}

type GetCertificateOpts struct {
	// Origin to get certificate for.
	Origin string `json:"origin"`
}

type GetCertificateResult struct {
	TableNames []string `json:"tableNames"`
}

// Returns the DER-encoded certificate. (experimental)
func (d *Domain) GetCertificate(opts *GetCertificateOpts) (*GetCertificateResult, error) {
	var result GetCertificateResult
	err := d.Client.Call("Network.getCertificate", opts, &result)
	return &result, err
}

type ResourceChangedPriorityEvent struct {
	// Request identifier.
	RequestId *RequestId `json:"requestId"`

	// New priority
	NewPriority *ResourcePriority `json:"newPriority"`

	// Timestamp.
	Timestamp *Timestamp `json:"timestamp"`
}

// Fired when resource loading priority is changed (experimental)
func (d *Domain) OnResourceChangedPriority(listener func(*ResourceChangedPriorityEvent)) {
	d.Client.AddListener("Network.resourceChangedPriority", func(params json.RawMessage) {
		var event ResourceChangedPriorityEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type RequestWillBeSentEvent struct {
	// Request identifier.
	RequestId *RequestId `json:"requestId"`

	// Frame identifier.
	FrameId interface{} `json:"frameId"`

	// Loader identifier.
	LoaderId *LoaderId `json:"loaderId"`

	// URL of the document this request is loaded for.
	DocumentURL string `json:"documentURL"`

	// Request data.
	Request *Request `json:"request"`

	// Timestamp.
	Timestamp *Timestamp `json:"timestamp"`

	// UTC Timestamp.
	WallTime *Timestamp `json:"wallTime"`

	// Request initiator.
	Initiator *Initiator `json:"initiator"`

	// Redirect response data. (optional)
	RedirectResponse *Response `json:"redirectResponse"`

	// Type of this resource. (optional, experimental)
	Type interface{} `json:"type"`
}

// Fired when page is about to send HTTP request.
func (d *Domain) OnRequestWillBeSent(listener func(*RequestWillBeSentEvent)) {
	d.Client.AddListener("Network.requestWillBeSent", func(params json.RawMessage) {
		var event RequestWillBeSentEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type RequestServedFromCacheEvent struct {
	// Request identifier.
	RequestId *RequestId `json:"requestId"`
}

// Fired if request ended up loading from cache.
func (d *Domain) OnRequestServedFromCache(listener func(*RequestServedFromCacheEvent)) {
	d.Client.AddListener("Network.requestServedFromCache", func(params json.RawMessage) {
		var event RequestServedFromCacheEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type ResponseReceivedEvent struct {
	// Request identifier.
	RequestId *RequestId `json:"requestId"`

	// Frame identifier.
	FrameId interface{} `json:"frameId"`

	// Loader identifier.
	LoaderId *LoaderId `json:"loaderId"`

	// Timestamp.
	Timestamp *Timestamp `json:"timestamp"`

	// Resource type.
	Type interface{} `json:"type"`

	// Response data.
	Response *Response `json:"response"`
}

// Fired when HTTP response is available.
func (d *Domain) OnResponseReceived(listener func(*ResponseReceivedEvent)) {
	d.Client.AddListener("Network.responseReceived", func(params json.RawMessage) {
		var event ResponseReceivedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type DataReceivedEvent struct {
	// Request identifier.
	RequestId *RequestId `json:"requestId"`

	// Timestamp.
	Timestamp *Timestamp `json:"timestamp"`

	// Data chunk length.
	DataLength int `json:"dataLength"`

	// Actual bytes received (might be less than dataLength for compressed encodings).
	EncodedDataLength int `json:"encodedDataLength"`
}

// Fired when data chunk was received over the network.
func (d *Domain) OnDataReceived(listener func(*DataReceivedEvent)) {
	d.Client.AddListener("Network.dataReceived", func(params json.RawMessage) {
		var event DataReceivedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type LoadingFinishedEvent struct {
	// Request identifier.
	RequestId *RequestId `json:"requestId"`

	// Timestamp.
	Timestamp *Timestamp `json:"timestamp"`

	// Total number of bytes received for this request.
	EncodedDataLength float64 `json:"encodedDataLength"`
}

// Fired when HTTP request has finished loading.
func (d *Domain) OnLoadingFinished(listener func(*LoadingFinishedEvent)) {
	d.Client.AddListener("Network.loadingFinished", func(params json.RawMessage) {
		var event LoadingFinishedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type LoadingFailedEvent struct {
	// Request identifier.
	RequestId *RequestId `json:"requestId"`

	// Timestamp.
	Timestamp *Timestamp `json:"timestamp"`

	// Resource type.
	Type interface{} `json:"type"`

	// User friendly error message.
	ErrorText string `json:"errorText"`

	// True if loading was canceled. (optional)
	Canceled bool `json:"canceled"`

	// The reason why loading was blocked, if any. (optional, experimental)
	BlockedReason *BlockedReason `json:"blockedReason"`
}

// Fired when HTTP request has failed to load.
func (d *Domain) OnLoadingFailed(listener func(*LoadingFailedEvent)) {
	d.Client.AddListener("Network.loadingFailed", func(params json.RawMessage) {
		var event LoadingFailedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type WebSocketWillSendHandshakeRequestEvent struct {
	// Request identifier.
	RequestId *RequestId `json:"requestId"`

	// Timestamp.
	Timestamp *Timestamp `json:"timestamp"`

	// UTC Timestamp.
	WallTime *Timestamp `json:"wallTime"`

	// WebSocket request data.
	Request *WebSocketRequest `json:"request"`
}

// Fired when WebSocket is about to initiate handshake. (experimental)
func (d *Domain) OnWebSocketWillSendHandshakeRequest(listener func(*WebSocketWillSendHandshakeRequestEvent)) {
	d.Client.AddListener("Network.webSocketWillSendHandshakeRequest", func(params json.RawMessage) {
		var event WebSocketWillSendHandshakeRequestEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type WebSocketHandshakeResponseReceivedEvent struct {
	// Request identifier.
	RequestId *RequestId `json:"requestId"`

	// Timestamp.
	Timestamp *Timestamp `json:"timestamp"`

	// WebSocket response data.
	Response *WebSocketResponse `json:"response"`
}

// Fired when WebSocket handshake response becomes available. (experimental)
func (d *Domain) OnWebSocketHandshakeResponseReceived(listener func(*WebSocketHandshakeResponseReceivedEvent)) {
	d.Client.AddListener("Network.webSocketHandshakeResponseReceived", func(params json.RawMessage) {
		var event WebSocketHandshakeResponseReceivedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type WebSocketCreatedEvent struct {
	// Request identifier.
	RequestId *RequestId `json:"requestId"`

	// WebSocket request URL.
	URL string `json:"url"`

	// Request initiator. (optional)
	Initiator *Initiator `json:"initiator"`
}

// Fired upon WebSocket creation. (experimental)
func (d *Domain) OnWebSocketCreated(listener func(*WebSocketCreatedEvent)) {
	d.Client.AddListener("Network.webSocketCreated", func(params json.RawMessage) {
		var event WebSocketCreatedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type WebSocketClosedEvent struct {
	// Request identifier.
	RequestId *RequestId `json:"requestId"`

	// Timestamp.
	Timestamp *Timestamp `json:"timestamp"`
}

// Fired when WebSocket is closed. (experimental)
func (d *Domain) OnWebSocketClosed(listener func(*WebSocketClosedEvent)) {
	d.Client.AddListener("Network.webSocketClosed", func(params json.RawMessage) {
		var event WebSocketClosedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type WebSocketFrameReceivedEvent struct {
	// Request identifier.
	RequestId *RequestId `json:"requestId"`

	// Timestamp.
	Timestamp *Timestamp `json:"timestamp"`

	// WebSocket response data.
	Response *WebSocketFrame `json:"response"`
}

// Fired when WebSocket frame is received. (experimental)
func (d *Domain) OnWebSocketFrameReceived(listener func(*WebSocketFrameReceivedEvent)) {
	d.Client.AddListener("Network.webSocketFrameReceived", func(params json.RawMessage) {
		var event WebSocketFrameReceivedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type WebSocketFrameErrorEvent struct {
	// Request identifier.
	RequestId *RequestId `json:"requestId"`

	// Timestamp.
	Timestamp *Timestamp `json:"timestamp"`

	// WebSocket frame error message.
	ErrorMessage string `json:"errorMessage"`
}

// Fired when WebSocket frame error occurs. (experimental)
func (d *Domain) OnWebSocketFrameError(listener func(*WebSocketFrameErrorEvent)) {
	d.Client.AddListener("Network.webSocketFrameError", func(params json.RawMessage) {
		var event WebSocketFrameErrorEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type WebSocketFrameSentEvent struct {
	// Request identifier.
	RequestId *RequestId `json:"requestId"`

	// Timestamp.
	Timestamp *Timestamp `json:"timestamp"`

	// WebSocket response data.
	Response *WebSocketFrame `json:"response"`
}

// Fired when WebSocket frame is sent. (experimental)
func (d *Domain) OnWebSocketFrameSent(listener func(*WebSocketFrameSentEvent)) {
	d.Client.AddListener("Network.webSocketFrameSent", func(params json.RawMessage) {
		var event WebSocketFrameSentEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type EventSourceMessageReceivedEvent struct {
	// Request identifier.
	RequestId *RequestId `json:"requestId"`

	// Timestamp.
	Timestamp *Timestamp `json:"timestamp"`

	// Message type.
	EventName string `json:"eventName"`

	// Message identifier.
	EventId string `json:"eventId"`

	// Message content.
	Data string `json:"data"`
}

// Fired when EventSource message is received. (experimental)
func (d *Domain) OnEventSourceMessageReceived(listener func(*EventSourceMessageReceivedEvent)) {
	d.Client.AddListener("Network.eventSourceMessageReceived", func(params json.RawMessage) {
		var event EventSourceMessageReceivedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}
