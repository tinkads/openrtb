package vast

type AdSystem struct {
	Version string `xml:"version,attr,omitempty"`
	System  string `xml:",chardata"`
}
