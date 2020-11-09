package vast

type Impression struct {
	Id  string      `xml:"id,attr,omitempty"`
	Uri TrimmedData `xml:",cdata"`
}
