package vast

type VideoClick struct {
	Id  string      `xml:"id,attr,omitempty"`
	Uri TrimmedData `xml:",cdata"`
}
