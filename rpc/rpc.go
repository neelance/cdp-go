package rpc

import (
	"encoding/json"
	"io"
	"math"
	"net/rpc"
)

type Client struct {
	*rpc.Client
	listeners map[string][]Listener
}

type Listener func(params json.RawMessage)

type clientCodec struct {
	client     *Client
	dec        *json.Decoder
	enc        *json.Encoder
	c          io.Closer
	lastResult json.RawMessage
}

func NewClient(conn io.ReadWriteCloser) *Client {
	cl := &Client{
		listeners: make(map[string][]Listener),
	}
	cl.Client = rpc.NewClientWithCodec(&clientCodec{
		client: cl,
		dec:    json.NewDecoder(conn),
		enc:    json.NewEncoder(conn),
		c:      conn,
	})
	return cl
}

func (cl *Client) AddListener(event string, listener Listener) {
	cl.listeners[event] = append(cl.listeners[event], listener)
}

func (c *clientCodec) WriteRequest(r *rpc.Request, param interface{}) error {
	return c.enc.Encode(struct {
		ID     uint64      `json:"id"`
		Method string      `json:"method"`
		Params interface{} `json:"params"`
	}{
		ID:     r.Seq,
		Method: r.ServiceMethod,
		Params: param,
	})
}

func (c *clientCodec) ReadResponseHeader(r *rpc.Response) error {
	var resp struct {
		// for responses
		ID     uint64           `json:"id"`
		Result json.RawMessage  `json:"result"`
		Error  *json.RawMessage `json:"error"`

		// for events
		Method string
		Params json.RawMessage
	}
	if err := c.dec.Decode(&resp); err != nil {
		return err
	}

	if resp.Method != "" {
		for _, l := range c.client.listeners[resp.Method] {
			l(resp.Params)
		}
		r.Seq = math.MaxUint64 // ignore
		return nil
	}

	r.Seq = resp.ID
	c.lastResult = resp.Result
	if resp.Error != nil {
		r.Error = string(*resp.Error)
	}
	return nil
}

func (c *clientCodec) ReadResponseBody(v interface{}) error {
	var err error
	if v != nil {
		err = json.Unmarshal(c.lastResult, v)
	}
	c.lastResult = nil
	return err
}

func (c *clientCodec) Close() error {
	return c.c.Close()
}
