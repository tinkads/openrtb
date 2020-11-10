package vast

type Linear struct {
	SkipOffset   *Offset       `xml:"skipoffset,attr,omitempty"`
	Duration     Duration      `xml:"Duration"`               // Required.
	AdParameters *AdParameters `xml:"AdParameters,omitempty"` // Just string in VAST2.0.
	Icons        []*Icon       `xml:"Icons>Icon"`             // VAST3.0.
	Trackings    []*Tracking   `xml:"TrackingEvents>Tracking,omitempty"`
	VideoClicks  *VideoClicks  `xml:"VideoClicks,omitempty"`
	MediaFiles   []*MediaFile  `xml:"MediaFiles>MediaFile,omitempty"`
	Extensions   []*Extension  `xml:"CreativeExtensions>CreativeExtension,omitempty"` // VAST3.0.
}
