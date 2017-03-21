// (experimental)
package animation

import (
	"encoding/json"
	"log"

	"github.com/neelance/cdp-go/rpc"
)

// (experimental)
type Domain struct {
	Client *rpc.Client
}

// Animation instance. (experimental)
type Animation interface{}

// AnimationEffect instance (experimental)
type AnimationEffect interface{}

// Keyframes Rule
type KeyframesRule interface{}

// Keyframe Style
type KeyframeStyle interface{}

// Enables animation domain notifications.
func (d *Domain) Enable() error {
	return d.Client.Call("Animation.enable", nil, nil)
}

// Disables animation domain notifications.
func (d *Domain) Disable() error {
	return d.Client.Call("Animation.disable", nil, nil)
}

type GetPlaybackRateResult struct {
	// Playback rate for animations on page.
	PlaybackRate float64 `json:"playbackRate"`
}

// Gets the playback rate of the document timeline.
func (d *Domain) GetPlaybackRate() (*GetPlaybackRateResult, error) {
	var result GetPlaybackRateResult
	err := d.Client.Call("Animation.getPlaybackRate", nil, &result)
	return &result, err
}

type SetPlaybackRateOpts struct {
	// Playback rate for animations on page
	PlaybackRate float64 `json:"playbackRate"`
}

// Sets the playback rate of the document timeline.
func (d *Domain) SetPlaybackRate(opts *SetPlaybackRateOpts) error {
	return d.Client.Call("Animation.setPlaybackRate", opts, nil)
}

type GetCurrentTimeOpts struct {
	// Id of animation.
	Id string `json:"id"`
}

type GetCurrentTimeResult struct {
	// Current time of the page.
	CurrentTime float64 `json:"currentTime"`
}

// Returns the current time of the an animation.
func (d *Domain) GetCurrentTime(opts *GetCurrentTimeOpts) (*GetCurrentTimeResult, error) {
	var result GetCurrentTimeResult
	err := d.Client.Call("Animation.getCurrentTime", opts, &result)
	return &result, err
}

type SetPausedOpts struct {
	// Animations to set the pause state of.
	Animations []string `json:"animations"`

	// Paused state to set to.
	Paused bool `json:"paused"`
}

// Sets the paused state of a set of animations.
func (d *Domain) SetPaused(opts *SetPausedOpts) error {
	return d.Client.Call("Animation.setPaused", opts, nil)
}

type SetTimingOpts struct {
	// Animation id.
	AnimationId string `json:"animationId"`

	// Duration of the animation.
	Duration float64 `json:"duration"`

	// Delay of the animation.
	Delay float64 `json:"delay"`
}

// Sets the timing of an animation node.
func (d *Domain) SetTiming(opts *SetTimingOpts) error {
	return d.Client.Call("Animation.setTiming", opts, nil)
}

type SeekAnimationsOpts struct {
	// List of animation ids to seek.
	Animations []string `json:"animations"`

	// Set the current time of each animation.
	CurrentTime float64 `json:"currentTime"`
}

// Seek a set of animations to a particular time within each animation.
func (d *Domain) SeekAnimations(opts *SeekAnimationsOpts) error {
	return d.Client.Call("Animation.seekAnimations", opts, nil)
}

type ReleaseAnimationsOpts struct {
	// List of animation ids to seek.
	Animations []string `json:"animations"`
}

// Releases a set of animations to no longer be manipulated.
func (d *Domain) ReleaseAnimations(opts *ReleaseAnimationsOpts) error {
	return d.Client.Call("Animation.releaseAnimations", opts, nil)
}

type ResolveAnimationOpts struct {
	// Animation id.
	AnimationId string `json:"animationId"`
}

type ResolveAnimationResult struct {
	// Corresponding remote object.
	RemoteObject interface{} `json:"remoteObject"`
}

// Gets the remote object of the Animation.
func (d *Domain) ResolveAnimation(opts *ResolveAnimationOpts) (*ResolveAnimationResult, error) {
	var result ResolveAnimationResult
	err := d.Client.Call("Animation.resolveAnimation", opts, &result)
	return &result, err
}

type AnimationCreatedEvent struct {
	// Id of the animation that was created.
	Id string `json:"id"`
}

// Event for each animation that has been created.
func (d *Domain) OnAnimationCreated(listener func(*AnimationCreatedEvent)) {
	d.Client.AddListener("Animation.animationCreated", func(params json.RawMessage) {
		var event AnimationCreatedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type AnimationStartedEvent struct {
	// Animation that was started.
	Animation Animation `json:"animation"`
}

// Event for animation that has been started.
func (d *Domain) OnAnimationStarted(listener func(*AnimationStartedEvent)) {
	d.Client.AddListener("Animation.animationStarted", func(params json.RawMessage) {
		var event AnimationStartedEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}

type AnimationCanceledEvent struct {
	// Id of the animation that was cancelled.
	Id string `json:"id"`
}

// Event for when an animation has been cancelled.
func (d *Domain) OnAnimationCanceled(listener func(*AnimationCanceledEvent)) {
	d.Client.AddListener("Animation.animationCanceled", func(params json.RawMessage) {
		var event AnimationCanceledEvent
		if err := json.Unmarshal(params, &event); err != nil {
			log.Print(err)
			return
		}
		listener(&event)
	})
}
