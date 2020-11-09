package vast

type LinearWrapper struct {
	Icons       []*Icon      `xml:"Icons>Icon,omitempty"` // VAST3.0.
	Trackings   []*Tracking  `xml:"TrackingEvents>Tracking,omitempty"`
	VideoClicks *VideoClicks `xml:"VideoClicks,omitempty"`                          // VideoClicks here is different from the Linear one.
	Extensions  []*Extension `xml:"CreativeExtensions>CreativeExtension,omitempty"` // VAST3.0.
}
