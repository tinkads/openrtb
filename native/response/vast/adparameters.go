package vast

type AdParameters struct {
	Parameters   TrimmedData `xml:",cdata"`
	IsXmlEncoded bool        `xml:"xmlEncoded,attr,omitempty"` // VAST3.0.
}
