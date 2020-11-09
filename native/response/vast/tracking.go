package vast

type Tracking struct {
	Event  Event       `xml:"event,attr"`            // Required.
	Offset *Offset     `xml:"offset,attr,omitempty"` // Time at which the event should be triggered. VAST3.0.
	Uri    TrimmedData `xml:",cdata"`
}
