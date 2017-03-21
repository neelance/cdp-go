// (experimental)
package deviceorientation

import (
	"github.com/neelance/cdp-go/rpc"
)

// (experimental)
type Domain struct {
	Client *rpc.Client
}

type SetDeviceOrientationOverrideOpts struct {
	// Mock alpha
	Alpha float64 `json:"alpha"`

	// Mock beta
	Beta float64 `json:"beta"`

	// Mock gamma
	Gamma float64 `json:"gamma"`
}

// Overrides the Device Orientation.
func (d *Domain) SetDeviceOrientationOverride(opts *SetDeviceOrientationOverrideOpts) error {
	return d.Client.Call("DeviceOrientation.setDeviceOrientationOverride", opts, nil)
}

// Clears the overridden Device Orientation.
func (d *Domain) ClearDeviceOrientationOverride() error {
	return d.Client.Call("DeviceOrientation.clearDeviceOrientationOverride", nil, nil)
}
