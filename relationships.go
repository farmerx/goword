package goword

import (
	"encoding/xml"
	"sync"
)

// relationships describe references from parts to other internal resources in the package or to external resources.
type relationships struct {
	sync.Mutex
	XMLName       xml.Name       `xml:"http://schemas.openxmlformats.org/package/2006/relationships Relationships"`
	Relationships []relationship `xml:"Relationship"`
}

// relationship contains relations which maps id and XML.
type relationship struct {
	ID         string `xml:"Id,attr"`
	Target     string `xml:",attr"`
	Type       string `xml:",attr"`
	TargetMode string `xml:",attr,omitempty"`
}
