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

type Animation struct {
	// <code>Animation</code>'s id.
	Id string `json:"id"`

	// <code>Animation</code>'s name.
	Name string `json:"name"`

	// <code>Animation</code>'s internal paused state.
	PausedState bool `json:"pausedState"`

	// <code>Animation</code>'s play state.
	PlayState string `json:"playState"`

	// <code>Animation</code>'s playback rate.
	PlaybackRate float64 `json:"playbackRate"`

	// <code>Animation</code>'s start time.
	StartTime float64 `json:"startTime"`

	// <code>Animation</code>'s current time.
	CurrentTime float64 `json:"currentTime"`

	// <code>Animation</code>'s source animation node.
	Source *AnimationEffect `json:"source"`

	// Animation type of <code>Animation</code>.
	Type string `json:"type"`

	// A unique ID for <code>Animation</code> representing the sources that triggered this CSS animation/transition. (optional)
	CssId string `json:"cssId,omitempty"`
}

// AnimationEffect instance (experimental)

type AnimationEffect struct {
	// <code>AnimationEffect</code>'s delay.
	Delay float64 `json:"delay"`

	// <code>AnimationEffect</code>'s end delay.
	EndDelay float64 `json:"endDelay"`

	// <code>AnimationEffect</code>'s iteration start.
	IterationStart float64 `json:"iterationStart"`

	// <code>AnimationEffect</code>'s iterations.
	Iterations float64 `json:"iterations"`

	// <code>AnimationEffect</code>'s iteration duration.
	Duration float64 `json:"duration"`

	// <code>AnimationEffect</code>'s playback direction.
	Direction string `json:"direction"`

	// <code>AnimationEffect</code>'s fill mode.
	Fill string `json:"fill"`

	// <code>AnimationEffect</code>'s target node.
	BackendNodeId interface{} `json:"backendNodeId"`

	// <code>AnimationEffect</code>'s keyframes. (optional)
	KeyframesRule *KeyframesRule `json:"keyframesRule,omitempty"`

	// <code>AnimationEffect</code>'s timing function.
	Easing string `json:"easing"`
}

// Keyframes Rule

type KeyframesRule struct {
	// CSS keyframed animation's name. (optional)
	Name string `json:"name,omitempty"`

	// List of animation keyframes.
	Keyframes []*KeyframeStyle `json:"keyframes"`
}

// Keyframe Style

type KeyframeStyle struct {
	// Keyframe's time offset.
	Offset string `json:"offset"`

	// <code>AnimationEffect</code>'s timing function.
	Easing string `json:"easing"`
}

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
	Animation *Animation `json:"animation"`
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
