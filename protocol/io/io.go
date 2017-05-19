// Input/Output operations for streams produced by DevTools. (experimental)
package io

import (
	"github.com/neelance/cdp-go/rpc"
)

// Input/Output operations for streams produced by DevTools. (experimental)
type Domain struct {
	Client *rpc.Client
}

type StreamHandle string

type ReadOpts struct {
	// Handle of the stream to read.
	Handle StreamHandle `json:"handle"`

	// Seek to the specified offset before reading (if not specificed, proceed with offset following the last read). (optional)
	Offset int `json:"offset,omitempty"`

	// Maximum number of bytes to read (left upon the agent discretion if not specified). (optional)
	Size int `json:"size,omitempty"`
}

type ReadResult struct {
	// Data that were read.
	Data string `json:"data"`

	// Set if the end-of-file condition occured while reading.
	Eof bool `json:"eof"`
}

// Read a chunk of the stream
func (d *Domain) Read(opts *ReadOpts) (*ReadResult, error) {
	var result ReadResult
	err := d.Client.Call("IO.read", opts, &result)
	return &result, err
}

type CloseOpts struct {
	// Handle of the stream to close.
	Handle StreamHandle `json:"handle"`
}

// Close the stream, discard any temporary backing storage.
func (d *Domain) Close(opts *CloseOpts) error {
	return d.Client.Call("IO.close", opts, nil)
}
