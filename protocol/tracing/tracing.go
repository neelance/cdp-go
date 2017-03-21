// (experimental)
package tracing

import (
	"encoding/json"
	"log"

	"github.com/neelance/cdp-go/rpc"
)

// (experimental)
type Domain struct {
	Client *rpc.Client
}

// Configuration for memory dump. Used only when "memory-infra" category is enabled.
type MemoryDumpConfig interface{}

type TraceConfig interface{}

type StartOpts struct {
	// Category/tag filter (optional)
	Categories string `json:"categories,omitempty"`

	// Tracing options (optional)
	Options string `json:"options,omitempty"`

	// If set, the agent will issue bufferUsage events at this interval, specified in milliseconds (optional)
	BufferUsageReportingInterval float64 `json:"bufferUsageReportingInterval,omitempty"`

	// Whether to report trace events as series of dataCollected events or to save trace to a stream (defaults to <code>ReportEvents</code>). (optional)
	TransferMode string `json:"transferMode,omitempty"`

	// (optional)
	TraceConfig TraceConfig `json:"traceConfig,omitempty"`
}

// Start trace events collection.
func (d *Domain) Start(opts *StartOpts) error {
	return d.Client.Call("Tracing.start", opts, nil)
}

// Stop trace events collection.
func (d *Domain) End() error {
	return d.Client.Call("Tracing.end", nil, nil)
}

type GetCategoriesResult struct {
	// A list of supported tracing categories.
	Categories []string `json:"categories"`
}

// Gets supported tracing categories.
func (d *Domain) GetCategories() (*GetCategoriesResult, error) {
	var result GetCategoriesResult
	err := d.Client.Call("Tracing.getCategories", nil, &result)
	return &result, err
}

type RequestMemoryDumpResult struct {
	// GUID of the resulting global memory dump.
	DumpGuid string `json:"dumpGuid"`

	// True iff the global memory dump succeeded.
	Success bool `json:"success"`
}

// Request a global memory dump.
func (d *Domain) RequestMemoryDump() (*RequestMemoryDumpResult, error) {
	var result RequestMemoryDumpResult
	err := d.Client.Call("Tracing.requestMemoryDump", nil, &result)
	return &result, err
}

type RecordClockSyncMarkerOpts struct {
	// The ID of this clock sync marker
	SyncId string `json:"syncId"`
}

// Record a clock sync marker in the trace.
func (d *Domain) RecordClockSyncMarker(opts *RecordClockSyncMarkerOpts) error {
	return d.Client.Call("Tracing.recordClockSyncMarker", opts, nil)
}

type DataCollectedEvent struct {
	Value []interface{} `json:"value"`
}

// Contains an bucket of collected trace events. When tracing is stopped collected events will be send as a sequence of dataCollected events followed by tracingComplete event.
func (d *Domain) OnDataCollected(listener func(*DataCollectedEvent)) {
	d.Client.AddListener("Tracing.dataCollected", func(params json.RawMessage) {
		var event DataCollectedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type TracingCompleteEvent struct {
	// A handle of the stream that holds resulting trace data. (optional)
	Stream interface{} `json:"stream"`
}

// Signals that tracing is stopped and there is no trace buffers pending flush, all data were delivered via dataCollected events.
func (d *Domain) OnTracingComplete(listener func(*TracingCompleteEvent)) {
	d.Client.AddListener("Tracing.tracingComplete", func(params json.RawMessage) {
		var event TracingCompleteEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type BufferUsageEvent struct {
	// A number in range [0..1] that indicates the used size of event buffer as a fraction of its total size. (optional)
	PercentFull float64 `json:"percentFull"`

	// An approximate number of events in the trace log. (optional)
	EventCount float64 `json:"eventCount"`

	// A number in range [0..1] that indicates the used size of event buffer as a fraction of its total size. (optional)
	Value float64 `json:"value"`
}

func (d *Domain) OnBufferUsage(listener func(*BufferUsageEvent)) {
	d.Client.AddListener("Tracing.bufferUsage", func(params json.RawMessage) {
		var event BufferUsageEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}
