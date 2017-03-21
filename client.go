package cdp

import (
	"github.com/neelance/cdp-go/protocol/accessibility"
	"github.com/neelance/cdp-go/protocol/animation"
	"github.com/neelance/cdp-go/protocol/applicationcache"
	"github.com/neelance/cdp-go/protocol/cachestorage"
	"github.com/neelance/cdp-go/protocol/css"
	"github.com/neelance/cdp-go/protocol/database"
	"github.com/neelance/cdp-go/protocol/deviceorientation"
	"github.com/neelance/cdp-go/protocol/dom"
	"github.com/neelance/cdp-go/protocol/domdebugger"
	"github.com/neelance/cdp-go/protocol/domstorage"
	"github.com/neelance/cdp-go/protocol/emulation"
	"github.com/neelance/cdp-go/protocol/indexeddb"
	"github.com/neelance/cdp-go/protocol/input"
	"github.com/neelance/cdp-go/protocol/inspector"
	"github.com/neelance/cdp-go/protocol/io"
	"github.com/neelance/cdp-go/protocol/layertree"
	"github.com/neelance/cdp-go/protocol/log"
	"github.com/neelance/cdp-go/protocol/memory"
	"github.com/neelance/cdp-go/protocol/network"
	"github.com/neelance/cdp-go/protocol/page"
	"github.com/neelance/cdp-go/protocol/rendering"
	"github.com/neelance/cdp-go/protocol/security"
	"github.com/neelance/cdp-go/protocol/serviceworker"
	"github.com/neelance/cdp-go/protocol/storage"
	"github.com/neelance/cdp-go/protocol/systeminfo"
	"github.com/neelance/cdp-go/protocol/target"
	"github.com/neelance/cdp-go/protocol/tethering"
	"github.com/neelance/cdp-go/protocol/tracing"
	"github.com/neelance/cdp-go/rpc"

	"golang.org/x/net/websocket"
)

type Client struct {
	Accessibility     accessibility.Domain
	Animation         animation.Domain
	ApplicationCache  applicationcache.Domain
	CacheStorage      cachestorage.Domain
	CSS               css.Domain
	Database          database.Domain
	DeviceOrientation deviceorientation.Domain
	DOM               dom.Domain
	DOMDebugger       domdebugger.Domain
	DOMStorage        domstorage.Domain
	Emulation         emulation.Domain
	IndexedDB         indexeddb.Domain
	Input             input.Domain
	Inspector         inspector.Domain
	IO                io.Domain
	LayerTree         layertree.Domain
	Log               log.Domain
	Memory            memory.Domain
	Network           network.Domain
	Page              page.Domain
	Rendering         rendering.Domain
	Security          security.Domain
	ServiceWorker     serviceworker.Domain
	Storage           storage.Domain
	SystemInfo        systeminfo.Domain
	Target            target.Domain
	Tethering         tethering.Domain
	Tracing           tracing.Domain
}

func Dial(url string) *Client {
	conn, err := websocket.Dial(url, "", url)
	if err != nil {
		panic(err)
	}

	cl := rpc.NewClient(conn)
	return &Client{
		Accessibility:     accessibility.Domain{Client: cl},
		Animation:         animation.Domain{Client: cl},
		ApplicationCache:  applicationcache.Domain{Client: cl},
		CacheStorage:      cachestorage.Domain{Client: cl},
		CSS:               css.Domain{Client: cl},
		Database:          database.Domain{Client: cl},
		DeviceOrientation: deviceorientation.Domain{Client: cl},
		DOM:               dom.Domain{Client: cl},
		DOMDebugger:       domdebugger.Domain{Client: cl},
		DOMStorage:        domstorage.Domain{Client: cl},
		Emulation:         emulation.Domain{Client: cl},
		IndexedDB:         indexeddb.Domain{Client: cl},
		Input:             input.Domain{Client: cl},
		Inspector:         inspector.Domain{Client: cl},
		IO:                io.Domain{Client: cl},
		LayerTree:         layertree.Domain{Client: cl},
		Log:               log.Domain{Client: cl},
		Memory:            memory.Domain{Client: cl},
		Network:           network.Domain{Client: cl},
		Page:              page.Domain{Client: cl},
		Rendering:         rendering.Domain{Client: cl},
		Security:          security.Domain{Client: cl},
		ServiceWorker:     serviceworker.Domain{Client: cl},
		Storage:           storage.Domain{Client: cl},
		SystemInfo:        systeminfo.Domain{Client: cl},
		Target:            target.Domain{Client: cl},
		Tethering:         tethering.Domain{Client: cl},
		Tracing:           tracing.Domain{Client: cl},
	}
}
