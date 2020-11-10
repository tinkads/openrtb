package vast

type Event string

const (
	EVENT_CREATIVE_VIEW     Event = "creativeView"
	EVENT_START             Event = "start"
	EVENT_FIRST_QUARTILE    Event = "firstQuartile"
	EVENT_MIDPOINT          Event = "midpoint"
	EVENT_THIRD_QUARTILE    Event = "thirdQuartile"
	EVENT_COMPLETE          Event = "complete"
	EVENT_MUTE              Event = "mute"
	EVENT_UNMUTE            Event = "unmute"
	EVENT_PAUSE             Event = "pause"
	EVENT_REWIND            Event = "rewind"
	EVENT_RESUME            Event = "resume"
	EVENT_FULLSCREEN        Event = "fullscreen"
	EVENT_EXIT_FULLSCREEN   Event = "exitFullscreen" //VAST3.0.
	EVENT_EXPAND            Event = "expand"
	EVENT_COLLAPSE          Event = "collapse"
	EVENT_ACCEPT_INVITATION Event = "acceptInvitation"
	EVENT_CLOSE             Event = "close"
	EVENT_SKIP              Event = "skip"     // VAST3.0.
	EVENT_PROGRESS          Event = "progress" // VAST3.0.
)
