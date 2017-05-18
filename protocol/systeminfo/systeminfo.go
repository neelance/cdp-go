// The SystemInfo domain defines methods and events for querying low-level system information. (experimental)
package systeminfo

import (
	"github.com/neelance/cdp-go/rpc"
)

// The SystemInfo domain defines methods and events for querying low-level system information. (experimental)
type Domain struct {
	Client *rpc.Client
}

// Describes a single graphics processor (GPU).
type GPUDevice interface{}

// Provides information about the GPU(s) on the system.
type GPUInfo interface{}

type GetInfoResult struct {
	// Information about the GPUs on the system.
	Gpu GPUInfo `json:"gpu"`

	// A platform-dependent description of the model of the machine. On Mac OS, this is, for example, 'MacBookPro'. Will be the empty string if not supported.
	ModelName string `json:"modelName"`

	// A platform-dependent description of the version of the machine. On Mac OS, this is, for example, '10.1'. Will be the empty string if not supported.
	ModelVersion string `json:"modelVersion"`

	// The command line string used to launch the browser. Will be the empty string if not supported.
	CommandLine string `json:"commandLine"`
}

// Returns information about the system.
func (d *Domain) GetInfo() (*GetInfoResult, error) {
	var result GetInfoResult
	err := d.Client.Call("SystemInfo.getInfo", nil, &result)
	return &result, err
}
