package main

import (
	"encoding/xml"
	"fmt"
)

var data = `
<?xml version="1.0" encoding="utf-8"?>
<D:propfind xmlns:D="DAV">
  <D:prop>
    <D:getlastmodified/>
    <D:getcontentlength/>
    <D:creationdate/>
    <D:resourcetype/>
  </D:prop>
</D:propfind>
`

type PropFindReq struct {
	XMLName xml.Name   `xml:"propfind"`
	Prop    []Prop `xml:"prop"`
}

type Prop struct {
	XMLName xml.Name `xml:"prop"`
	GetLastModified xml.Name `xml:"getlastmodified"`
	GetContentLength xml.Name `xml:"getcontentlength"`
	CreationDate xml.Name `xml:"creationdate"`
	ResourceType xml.Name `xml:"resourcetype"`
}

func main() {
	v := &PropFindReq{}
	err := xml.Unmarshal([]byte(data), v)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("%#v\n", v)
	}
}

