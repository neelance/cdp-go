// Debugger domain exposes JavaScript debugging capabilities. It allows setting and removing breakpoints, stepping through execution, exploring stack traces, etc.
package debugger

import (
	"github.com/neelance/cdp-go/rpc"
)

// Debugger domain exposes JavaScript debugging capabilities. It allows setting and removing breakpoints, stepping through execution, exploring stack traces, etc.
type Client struct {
	*rpc.Client
}

// Breakpoint identifier.

type BreakpointId string

// Call frame identifier.

type CallFrameId string

// Location in the source code.

type Location struct {
	// Script identifier as reported in the <code>Debugger.scriptParsed</code>.
	ScriptId interface{} `json:"scriptId"`

	// Line number in the script (0-based).
	LineNumber int `json:"lineNumber"`

	// Column number in the script (0-based). (optional)
	ColumnNumber int `json:"columnNumber,omitempty"`
}

// Location in the source code. (experimental)

type ScriptPosition struct {
	LineNumber int `json:"lineNumber"`

	ColumnNumber int `json:"columnNumber"`
}

// JavaScript call frame. Array of call frames form the call stack.

type CallFrame struct {
	// Call frame identifier. This identifier is only valid while the virtual machine is paused.
	CallFrameId CallFrameId `json:"callFrameId"`

	// Name of the JavaScript function called on this call frame.
	FunctionName string `json:"functionName"`

	// Location in the source code. (optional, experimental)
	FunctionLocation *Location `json:"functionLocation,omitempty"`

	// Location in the source code.
	Location *Location `json:"location"`

	// Scope chain for this call frame.
	ScopeChain []*Scope `json:"scopeChain"`

	// <code>this</code> object for this call frame.
	This interface{} `json:"this"`

	// The value being returned, if the function is at return point. (optional)
	ReturnValue interface{} `json:"returnValue,omitempty"`
}

// Scope description.

type Scope struct {
	// Scope type.
	Type string `json:"type"`

	// Object representing the scope. For <code>global</code> and <code>with</code> scopes it represents the actual object; for the rest of the scopes, it is artificial transient object enumerating scope variables as its properties.
	Object interface{} `json:"object"`

	// (optional)
	Name string `json:"name,omitempty"`

	// Location in the source code where scope starts (optional)
	StartLocation *Location `json:"startLocation,omitempty"`

	// Location in the source code where scope ends (optional)
	EndLocation *Location `json:"endLocation,omitempty"`
}

// Search match for resource. (experimental)

type SearchMatch struct {
	// Line number in resource content.
	LineNumber float64 `json:"lineNumber"`

	// Line with match content.
	LineContent string `json:"lineContent"`
}

// (experimental)

type BreakLocation struct {
	// Script identifier as reported in the <code>Debugger.scriptParsed</code>.
	ScriptId interface{} `json:"scriptId"`

	// Line number in the script (0-based).
	LineNumber int `json:"lineNumber"`

	// Column number in the script (0-based). (optional)
	ColumnNumber int `json:"columnNumber,omitempty"`

	// (optional)
	Type string `json:"type,omitempty"`
}

// Enables debugger for the given page. Clients should not assume that the debugging has been enabled until the result for this command is received.
type EnableRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Client) Enable() *EnableRequest {
	return &EnableRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Enables debugger for the given page. Clients should not assume that the debugging has been enabled until the result for this command is received.
func (r *EnableRequest) Do() error {
	return r.client.Call("Debugger.enable", r.opts, nil)
}

// Disables debugger for given page.
type DisableRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Client) Disable() *DisableRequest {
	return &DisableRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Disables debugger for given page.
func (r *DisableRequest) Do() error {
	return r.client.Call("Debugger.disable", r.opts, nil)
}

// Activates / deactivates all breakpoints on the page.
type SetBreakpointsActiveRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Client) SetBreakpointsActive() *SetBreakpointsActiveRequest {
	return &SetBreakpointsActiveRequest{opts: make(map[string]interface{}), client: d.Client}
}

// New value for breakpoints active state.
func (r *SetBreakpointsActiveRequest) Active(v bool) *SetBreakpointsActiveRequest {
	r.opts["active"] = v
	return r
}

// Activates / deactivates all breakpoints on the page.
func (r *SetBreakpointsActiveRequest) Do() error {
	return r.client.Call("Debugger.setBreakpointsActive", r.opts, nil)
}

// Makes page not interrupt on any pauses (breakpoint, exception, dom exception etc).
type SetSkipAllPausesRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Client) SetSkipAllPauses() *SetSkipAllPausesRequest {
	return &SetSkipAllPausesRequest{opts: make(map[string]interface{}), client: d.Client}
}

// New value for skip pauses state.
func (r *SetSkipAllPausesRequest) Skip(v bool) *SetSkipAllPausesRequest {
	r.opts["skip"] = v
	return r
}

// Makes page not interrupt on any pauses (breakpoint, exception, dom exception etc).
func (r *SetSkipAllPausesRequest) Do() error {
	return r.client.Call("Debugger.setSkipAllPauses", r.opts, nil)
}

// Sets JavaScript breakpoint at given location specified either by URL or URL regex. Once this command is issued, all existing parsed scripts will have breakpoints resolved and returned in <code>locations</code> property. Further matching script parsing will result in subsequent <code>breakpointResolved</code> events issued. This logical breakpoint will survive page reloads.
type SetBreakpointByUrlRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Client) SetBreakpointByUrl() *SetBreakpointByUrlRequest {
	return &SetBreakpointByUrlRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Line number to set breakpoint at.
func (r *SetBreakpointByUrlRequest) LineNumber(v int) *SetBreakpointByUrlRequest {
	r.opts["lineNumber"] = v
	return r
}

// URL of the resources to set breakpoint on. (optional)
func (r *SetBreakpointByUrlRequest) URL(v string) *SetBreakpointByUrlRequest {
	r.opts["url"] = v
	return r
}

// Regex pattern for the URLs of the resources to set breakpoints on. Either <code>url</code> or <code>urlRegex</code> must be specified. (optional)
func (r *SetBreakpointByUrlRequest) UrlRegex(v string) *SetBreakpointByUrlRequest {
	r.opts["urlRegex"] = v
	return r
}

// Offset in the line to set breakpoint at. (optional)
func (r *SetBreakpointByUrlRequest) ColumnNumber(v int) *SetBreakpointByUrlRequest {
	r.opts["columnNumber"] = v
	return r
}

// Expression to use as a breakpoint condition. When specified, debugger will only stop on the breakpoint if this expression evaluates to true. (optional)
func (r *SetBreakpointByUrlRequest) Condition(v string) *SetBreakpointByUrlRequest {
	r.opts["condition"] = v
	return r
}

type SetBreakpointByUrlResult struct {
	// Id of the created breakpoint for further reference.
	BreakpointId BreakpointId `json:"breakpointId"`

	// List of the locations this breakpoint resolved into upon addition.
	Locations []*Location `json:"locations"`
}

func (r *SetBreakpointByUrlRequest) Do() (*SetBreakpointByUrlResult, error) {
	var result SetBreakpointByUrlResult
	err := r.client.Call("Debugger.setBreakpointByUrl", r.opts, &result)
	return &result, err
}

// Sets JavaScript breakpoint at a given location.
type SetBreakpointRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Client) SetBreakpoint() *SetBreakpointRequest {
	return &SetBreakpointRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Location to set breakpoint in.
func (r *SetBreakpointRequest) Location(v *Location) *SetBreakpointRequest {
	r.opts["location"] = v
	return r
}

// Expression to use as a breakpoint condition. When specified, debugger will only stop on the breakpoint if this expression evaluates to true. (optional)
func (r *SetBreakpointRequest) Condition(v string) *SetBreakpointRequest {
	r.opts["condition"] = v
	return r
}

type SetBreakpointResult struct {
	// Id of the created breakpoint for further reference.
	BreakpointId BreakpointId `json:"breakpointId"`

	// Location this breakpoint resolved into.
	ActualLocation *Location `json:"actualLocation"`
}

func (r *SetBreakpointRequest) Do() (*SetBreakpointResult, error) {
	var result SetBreakpointResult
	err := r.client.Call("Debugger.setBreakpoint", r.opts, &result)
	return &result, err
}

// Removes JavaScript breakpoint.
type RemoveBreakpointRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Client) RemoveBreakpoint() *RemoveBreakpointRequest {
	return &RemoveBreakpointRequest{opts: make(map[string]interface{}), client: d.Client}
}

func (r *RemoveBreakpointRequest) BreakpointId(v BreakpointId) *RemoveBreakpointRequest {
	r.opts["breakpointId"] = v
	return r
}

// Removes JavaScript breakpoint.
func (r *RemoveBreakpointRequest) Do() error {
	return r.client.Call("Debugger.removeBreakpoint", r.opts, nil)
}

// Returns possible locations for breakpoint. scriptId in start and end range locations should be the same. (experimental)
type GetPossibleBreakpointsRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Client) GetPossibleBreakpoints() *GetPossibleBreakpointsRequest {
	return &GetPossibleBreakpointsRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Start of range to search possible breakpoint locations in.
func (r *GetPossibleBreakpointsRequest) Start(v *Location) *GetPossibleBreakpointsRequest {
	r.opts["start"] = v
	return r
}

// End of range to search possible breakpoint locations in (excluding). When not specified, end of scripts is used as end of range. (optional)
func (r *GetPossibleBreakpointsRequest) End(v *Location) *GetPossibleBreakpointsRequest {
	r.opts["end"] = v
	return r
}

// Only consider locations which are in the same (non-nested) function as start. (optional)
func (r *GetPossibleBreakpointsRequest) RestrictToFunction(v bool) *GetPossibleBreakpointsRequest {
	r.opts["restrictToFunction"] = v
	return r
}

type GetPossibleBreakpointsResult struct {
	// List of the possible breakpoint locations.
	Locations []*BreakLocation `json:"locations"`
}

func (r *GetPossibleBreakpointsRequest) Do() (*GetPossibleBreakpointsResult, error) {
	var result GetPossibleBreakpointsResult
	err := r.client.Call("Debugger.getPossibleBreakpoints", r.opts, &result)
	return &result, err
}

// Continues execution until specific location is reached.
type ContinueToLocationRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Client) ContinueToLocation() *ContinueToLocationRequest {
	return &ContinueToLocationRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Location to continue to.
func (r *ContinueToLocationRequest) Location(v *Location) *ContinueToLocationRequest {
	r.opts["location"] = v
	return r
}

// (optional, experimental)
func (r *ContinueToLocationRequest) TargetCallFrames(v string) *ContinueToLocationRequest {
	r.opts["targetCallFrames"] = v
	return r
}

// Continues execution until specific location is reached.
func (r *ContinueToLocationRequest) Do() error {
	return r.client.Call("Debugger.continueToLocation", r.opts, nil)
}

// Steps over the statement.
type StepOverRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Client) StepOver() *StepOverRequest {
	return &StepOverRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Steps over the statement.
func (r *StepOverRequest) Do() error {
	return r.client.Call("Debugger.stepOver", r.opts, nil)
}

// Steps into the function call.
type StepIntoRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Client) StepInto() *StepIntoRequest {
	return &StepIntoRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Steps into the function call.
func (r *StepIntoRequest) Do() error {
	return r.client.Call("Debugger.stepInto", r.opts, nil)
}

// Steps out of the function call.
type StepOutRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Client) StepOut() *StepOutRequest {
	return &StepOutRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Steps out of the function call.
func (r *StepOutRequest) Do() error {
	return r.client.Call("Debugger.stepOut", r.opts, nil)
}

// Stops on the next JavaScript statement.
type PauseRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Client) Pause() *PauseRequest {
	return &PauseRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Stops on the next JavaScript statement.
func (r *PauseRequest) Do() error {
	return r.client.Call("Debugger.pause", r.opts, nil)
}

// Steps into next scheduled async task if any is scheduled before next pause. Returns success when async task is actually scheduled, returns error if no task were scheduled or another scheduleStepIntoAsync was called. (experimental)
type ScheduleStepIntoAsyncRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Client) ScheduleStepIntoAsync() *ScheduleStepIntoAsyncRequest {
	return &ScheduleStepIntoAsyncRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Steps into next scheduled async task if any is scheduled before next pause. Returns success when async task is actually scheduled, returns error if no task were scheduled or another scheduleStepIntoAsync was called. (experimental)
func (r *ScheduleStepIntoAsyncRequest) Do() error {
	return r.client.Call("Debugger.scheduleStepIntoAsync", r.opts, nil)
}

// Resumes JavaScript execution.
type ResumeRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Client) Resume() *ResumeRequest {
	return &ResumeRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Resumes JavaScript execution.
func (r *ResumeRequest) Do() error {
	return r.client.Call("Debugger.resume", r.opts, nil)
}

// Searches for given string in script content. (experimental)
type SearchInContentRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Client) SearchInContent() *SearchInContentRequest {
	return &SearchInContentRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Id of the script to search in.
func (r *SearchInContentRequest) ScriptId(v interface{}) *SearchInContentRequest {
	r.opts["scriptId"] = v
	return r
}

// String to search for.
func (r *SearchInContentRequest) Query(v string) *SearchInContentRequest {
	r.opts["query"] = v
	return r
}

// If true, search is case sensitive. (optional)
func (r *SearchInContentRequest) CaseSensitive(v bool) *SearchInContentRequest {
	r.opts["caseSensitive"] = v
	return r
}

// If true, treats string parameter as regex. (optional)
func (r *SearchInContentRequest) IsRegex(v bool) *SearchInContentRequest {
	r.opts["isRegex"] = v
	return r
}

type SearchInContentResult struct {
	// List of search matches.
	Result []*SearchMatch `json:"result"`
}

func (r *SearchInContentRequest) Do() (*SearchInContentResult, error) {
	var result SearchInContentResult
	err := r.client.Call("Debugger.searchInContent", r.opts, &result)
	return &result, err
}

// Edits JavaScript source live.
type SetScriptSourceRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Client) SetScriptSource() *SetScriptSourceRequest {
	return &SetScriptSourceRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Id of the script to edit.
func (r *SetScriptSourceRequest) ScriptId(v interface{}) *SetScriptSourceRequest {
	r.opts["scriptId"] = v
	return r
}

// New content of the script.
func (r *SetScriptSourceRequest) ScriptSource(v string) *SetScriptSourceRequest {
	r.opts["scriptSource"] = v
	return r
}

// If true the change will not actually be applied. Dry run may be used to get result description without actually modifying the code. (optional)
func (r *SetScriptSourceRequest) DryRun(v bool) *SetScriptSourceRequest {
	r.opts["dryRun"] = v
	return r
}

type SetScriptSourceResult struct {
	// New stack trace in case editing has happened while VM was stopped. (optional)
	CallFrames []*CallFrame `json:"callFrames"`

	// Whether current call stack  was modified after applying the changes. (optional)
	StackChanged bool `json:"stackChanged"`

	// Async stack trace, if any. (optional)
	AsyncStackTrace interface{} `json:"asyncStackTrace"`

	// Exception details if any. (optional)
	ExceptionDetails interface{} `json:"exceptionDetails"`
}

func (r *SetScriptSourceRequest) Do() (*SetScriptSourceResult, error) {
	var result SetScriptSourceResult
	err := r.client.Call("Debugger.setScriptSource", r.opts, &result)
	return &result, err
}

// Restarts particular call frame from the beginning.
type RestartFrameRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Client) RestartFrame() *RestartFrameRequest {
	return &RestartFrameRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Call frame identifier to evaluate on.
func (r *RestartFrameRequest) CallFrameId(v CallFrameId) *RestartFrameRequest {
	r.opts["callFrameId"] = v
	return r
}

type RestartFrameResult struct {
	// New stack trace.
	CallFrames []*CallFrame `json:"callFrames"`

	// Async stack trace, if any. (optional)
	AsyncStackTrace interface{} `json:"asyncStackTrace"`
}

func (r *RestartFrameRequest) Do() (*RestartFrameResult, error) {
	var result RestartFrameResult
	err := r.client.Call("Debugger.restartFrame", r.opts, &result)
	return &result, err
}

// Returns source for the script with given id.
type GetScriptSourceRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Client) GetScriptSource() *GetScriptSourceRequest {
	return &GetScriptSourceRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Id of the script to get source for.
func (r *GetScriptSourceRequest) ScriptId(v interface{}) *GetScriptSourceRequest {
	r.opts["scriptId"] = v
	return r
}

type GetScriptSourceResult struct {
	// Script source.
	ScriptSource string `json:"scriptSource"`
}

func (r *GetScriptSourceRequest) Do() (*GetScriptSourceResult, error) {
	var result GetScriptSourceResult
	err := r.client.Call("Debugger.getScriptSource", r.opts, &result)
	return &result, err
}

// Defines pause on exceptions state. Can be set to stop on all exceptions, uncaught exceptions or no exceptions. Initial pause on exceptions state is <code>none</code>.
type SetPauseOnExceptionsRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Client) SetPauseOnExceptions() *SetPauseOnExceptionsRequest {
	return &SetPauseOnExceptionsRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Pause on exceptions mode.
func (r *SetPauseOnExceptionsRequest) State(v string) *SetPauseOnExceptionsRequest {
	r.opts["state"] = v
	return r
}

// Defines pause on exceptions state. Can be set to stop on all exceptions, uncaught exceptions or no exceptions. Initial pause on exceptions state is <code>none</code>.
func (r *SetPauseOnExceptionsRequest) Do() error {
	return r.client.Call("Debugger.setPauseOnExceptions", r.opts, nil)
}

// Evaluates expression on a given call frame.
type EvaluateOnCallFrameRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Client) EvaluateOnCallFrame() *EvaluateOnCallFrameRequest {
	return &EvaluateOnCallFrameRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Call frame identifier to evaluate on.
func (r *EvaluateOnCallFrameRequest) CallFrameId(v CallFrameId) *EvaluateOnCallFrameRequest {
	r.opts["callFrameId"] = v
	return r
}

// Expression to evaluate.
func (r *EvaluateOnCallFrameRequest) Expression(v string) *EvaluateOnCallFrameRequest {
	r.opts["expression"] = v
	return r
}

// String object group name to put result into (allows rapid releasing resulting object handles using <code>releaseObjectGroup</code>). (optional)
func (r *EvaluateOnCallFrameRequest) ObjectGroup(v string) *EvaluateOnCallFrameRequest {
	r.opts["objectGroup"] = v
	return r
}

// Specifies whether command line API should be available to the evaluated expression, defaults to false. (optional)
func (r *EvaluateOnCallFrameRequest) IncludeCommandLineAPI(v bool) *EvaluateOnCallFrameRequest {
	r.opts["includeCommandLineAPI"] = v
	return r
}

// In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides <code>setPauseOnException</code> state. (optional)
func (r *EvaluateOnCallFrameRequest) Silent(v bool) *EvaluateOnCallFrameRequest {
	r.opts["silent"] = v
	return r
}

// Whether the result is expected to be a JSON object that should be sent by value. (optional)
func (r *EvaluateOnCallFrameRequest) ReturnByValue(v bool) *EvaluateOnCallFrameRequest {
	r.opts["returnByValue"] = v
	return r
}

// Whether preview should be generated for the result. (optional, experimental)
func (r *EvaluateOnCallFrameRequest) GeneratePreview(v bool) *EvaluateOnCallFrameRequest {
	r.opts["generatePreview"] = v
	return r
}

// Whether to throw an exception if side effect cannot be ruled out during evaluation. (optional, experimental)
func (r *EvaluateOnCallFrameRequest) ThrowOnSideEffect(v bool) *EvaluateOnCallFrameRequest {
	r.opts["throwOnSideEffect"] = v
	return r
}

type EvaluateOnCallFrameResult struct {
	// Object wrapper for the evaluation result.
	Result interface{} `json:"result"`

	// Exception details. (optional)
	ExceptionDetails interface{} `json:"exceptionDetails"`
}

func (r *EvaluateOnCallFrameRequest) Do() (*EvaluateOnCallFrameResult, error) {
	var result EvaluateOnCallFrameResult
	err := r.client.Call("Debugger.evaluateOnCallFrame", r.opts, &result)
	return &result, err
}

// Changes value of variable in a callframe. Object-based scopes are not supported and must be mutated manually.
type SetVariableValueRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Client) SetVariableValue() *SetVariableValueRequest {
	return &SetVariableValueRequest{opts: make(map[string]interface{}), client: d.Client}
}

// 0-based number of scope as was listed in scope chain. Only 'local', 'closure' and 'catch' scope types are allowed. Other scopes could be manipulated manually.
func (r *SetVariableValueRequest) ScopeNumber(v int) *SetVariableValueRequest {
	r.opts["scopeNumber"] = v
	return r
}

// Variable name.
func (r *SetVariableValueRequest) VariableName(v string) *SetVariableValueRequest {
	r.opts["variableName"] = v
	return r
}

// New variable value.
func (r *SetVariableValueRequest) NewValue(v interface{}) *SetVariableValueRequest {
	r.opts["newValue"] = v
	return r
}

// Id of callframe that holds variable.
func (r *SetVariableValueRequest) CallFrameId(v CallFrameId) *SetVariableValueRequest {
	r.opts["callFrameId"] = v
	return r
}

// Changes value of variable in a callframe. Object-based scopes are not supported and must be mutated manually.
func (r *SetVariableValueRequest) Do() error {
	return r.client.Call("Debugger.setVariableValue", r.opts, nil)
}

// Enables or disables async call stacks tracking.
type SetAsyncCallStackDepthRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Client) SetAsyncCallStackDepth() *SetAsyncCallStackDepthRequest {
	return &SetAsyncCallStackDepthRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Maximum depth of async call stacks. Setting to <code>0</code> will effectively disable collecting async call stacks (default).
func (r *SetAsyncCallStackDepthRequest) MaxDepth(v int) *SetAsyncCallStackDepthRequest {
	r.opts["maxDepth"] = v
	return r
}

// Enables or disables async call stacks tracking.
func (r *SetAsyncCallStackDepthRequest) Do() error {
	return r.client.Call("Debugger.setAsyncCallStackDepth", r.opts, nil)
}

// Replace previous blackbox patterns with passed ones. Forces backend to skip stepping/pausing in scripts with url matching one of the patterns. VM will try to leave blackboxed script by performing 'step in' several times, finally resorting to 'step out' if unsuccessful. (experimental)
type SetBlackboxPatternsRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Client) SetBlackboxPatterns() *SetBlackboxPatternsRequest {
	return &SetBlackboxPatternsRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Array of regexps that will be used to check script url for blackbox state.
func (r *SetBlackboxPatternsRequest) Patterns(v []string) *SetBlackboxPatternsRequest {
	r.opts["patterns"] = v
	return r
}

// Replace previous blackbox patterns with passed ones. Forces backend to skip stepping/pausing in scripts with url matching one of the patterns. VM will try to leave blackboxed script by performing 'step in' several times, finally resorting to 'step out' if unsuccessful. (experimental)
func (r *SetBlackboxPatternsRequest) Do() error {
	return r.client.Call("Debugger.setBlackboxPatterns", r.opts, nil)
}

// Makes backend skip steps in the script in blackboxed ranges. VM will try leave blacklisted scripts by performing 'step in' several times, finally resorting to 'step out' if unsuccessful. Positions array contains positions where blackbox state is changed. First interval isn't blackboxed. Array should be sorted. (experimental)
type SetBlackboxedRangesRequest struct {
	client *rpc.Client
	opts   map[string]interface{}
}

func (d *Client) SetBlackboxedRanges() *SetBlackboxedRangesRequest {
	return &SetBlackboxedRangesRequest{opts: make(map[string]interface{}), client: d.Client}
}

// Id of the script.
func (r *SetBlackboxedRangesRequest) ScriptId(v interface{}) *SetBlackboxedRangesRequest {
	r.opts["scriptId"] = v
	return r
}

func (r *SetBlackboxedRangesRequest) Positions(v []*ScriptPosition) *SetBlackboxedRangesRequest {
	r.opts["positions"] = v
	return r
}

// Makes backend skip steps in the script in blackboxed ranges. VM will try leave blacklisted scripts by performing 'step in' several times, finally resorting to 'step out' if unsuccessful. Positions array contains positions where blackbox state is changed. First interval isn't blackboxed. Array should be sorted. (experimental)
func (r *SetBlackboxedRangesRequest) Do() error {
	return r.client.Call("Debugger.setBlackboxedRanges", r.opts, nil)
}

func init() {
	rpc.EventTypes["Debugger.scriptParsed"] = func() interface{} { return new(ScriptParsedEvent) }
	rpc.EventTypes["Debugger.scriptFailedToParse"] = func() interface{} { return new(ScriptFailedToParseEvent) }
	rpc.EventTypes["Debugger.breakpointResolved"] = func() interface{} { return new(BreakpointResolvedEvent) }
	rpc.EventTypes["Debugger.paused"] = func() interface{} { return new(PausedEvent) }
	rpc.EventTypes["Debugger.resumed"] = func() interface{} { return new(ResumedEvent) }
}

// Fired when virtual machine parses script. This event is also fired for all known and uncollected scripts upon enabling debugger.
type ScriptParsedEvent struct {
	// Identifier of the script parsed.
	ScriptId interface{} `json:"scriptId"`

	// URL or name of the script parsed (if any).
	URL string `json:"url"`

	// Line offset of the script within the resource with given URL (for script tags).
	StartLine int `json:"startLine"`

	// Column offset of the script within the resource with given URL.
	StartColumn int `json:"startColumn"`

	// Last line of the script.
	EndLine int `json:"endLine"`

	// Length of the last line of the script.
	EndColumn int `json:"endColumn"`

	// Specifies script creation context.
	ExecutionContextId interface{} `json:"executionContextId"`

	// Content hash of the script.
	Hash string `json:"hash"`

	// Embedder-specific auxiliary data. (optional)
	ExecutionContextAuxData interface{} `json:"executionContextAuxData"`

	// True, if this script is generated as a result of the live edit operation. (optional, experimental)
	IsLiveEdit bool `json:"isLiveEdit"`

	// URL of source map associated with script (if any). (optional)
	SourceMapURL string `json:"sourceMapURL"`

	// True, if this script has sourceURL. (optional, experimental)
	HasSourceURL bool `json:"hasSourceURL"`

	// True, if this script is ES6 module. (optional, experimental)
	IsModule bool `json:"isModule"`

	// This script length. (optional, experimental)
	Length int `json:"length"`

	// JavaScript top stack frame of where the script parsed event was triggered if available. (optional, experimental)
	StackTrace interface{} `json:"stackTrace"`
}

// Fired when virtual machine fails to parse the script.
type ScriptFailedToParseEvent struct {
	// Identifier of the script parsed.
	ScriptId interface{} `json:"scriptId"`

	// URL or name of the script parsed (if any).
	URL string `json:"url"`

	// Line offset of the script within the resource with given URL (for script tags).
	StartLine int `json:"startLine"`

	// Column offset of the script within the resource with given URL.
	StartColumn int `json:"startColumn"`

	// Last line of the script.
	EndLine int `json:"endLine"`

	// Length of the last line of the script.
	EndColumn int `json:"endColumn"`

	// Specifies script creation context.
	ExecutionContextId interface{} `json:"executionContextId"`

	// Content hash of the script.
	Hash string `json:"hash"`

	// Embedder-specific auxiliary data. (optional)
	ExecutionContextAuxData interface{} `json:"executionContextAuxData"`

	// URL of source map associated with script (if any). (optional)
	SourceMapURL string `json:"sourceMapURL"`

	// True, if this script has sourceURL. (optional, experimental)
	HasSourceURL bool `json:"hasSourceURL"`

	// True, if this script is ES6 module. (optional, experimental)
	IsModule bool `json:"isModule"`

	// This script length. (optional, experimental)
	Length int `json:"length"`

	// JavaScript top stack frame of where the script parsed event was triggered if available. (optional, experimental)
	StackTrace interface{} `json:"stackTrace"`
}

// Fired when breakpoint is resolved to an actual script and location.
type BreakpointResolvedEvent struct {
	// Breakpoint unique identifier.
	BreakpointId BreakpointId `json:"breakpointId"`

	// Actual breakpoint location.
	Location *Location `json:"location"`
}

// Fired when the virtual machine stopped on breakpoint or exception or any other stop criteria.
type PausedEvent struct {
	// Call stack the virtual machine stopped on.
	CallFrames []*CallFrame `json:"callFrames"`

	// Pause reason.
	Reason string `json:"reason"`

	// Object containing break-specific auxiliary properties. (optional)
	Data interface{} `json:"data"`

	// Hit breakpoints IDs (optional)
	HitBreakpoints []string `json:"hitBreakpoints"`

	// Async stack trace, if any. (optional)
	AsyncStackTrace interface{} `json:"asyncStackTrace"`
}

// Fired when the virtual machine resumed execution.
type ResumedEvent struct {
}
