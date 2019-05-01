// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@github.com/melignus.

package dml

import (
	"encoding/xml"
	"fmt"

	"github.com/melignus/gooxml"
)

type CT_FontScheme struct {
	NameAttr  string
	MajorFont *CT_FontCollection
	MinorFont *CT_FontCollection
	ExtLst    *CT_OfficeArtExtensionList
}

func NewCT_FontScheme() *CT_FontScheme {
	ret := &CT_FontScheme{}
	ret.MajorFont = NewCT_FontCollection()
	ret.MinorFont = NewCT_FontCollection()
	return ret
}

func (m *CT_FontScheme) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	e.EncodeToken(start)
	semajorFont := xml.StartElement{Name: xml.Name{Local: "a:majorFont"}}
	e.EncodeElement(m.MajorFont, semajorFont)
	seminorFont := xml.StartElement{Name: xml.Name{Local: "a:minorFont"}}
	e.EncodeElement(m.MinorFont, seminorFont)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_FontScheme) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.MajorFont = NewCT_FontCollection()
	m.MinorFont = NewCT_FontCollection()
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
			continue
		}
	}
lCT_FontScheme:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "majorFont"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "majorFont"}:
				if err := d.DecodeElement(m.MajorFont, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "minorFont"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "minorFont"}:
				if err := d.DecodeElement(m.MinorFont, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "extLst"}:
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_FontScheme %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_FontScheme
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_FontScheme and its children
func (m *CT_FontScheme) Validate() error {
	return m.ValidateWithPath("CT_FontScheme")
}

// ValidateWithPath validates the CT_FontScheme and its children, prefixing error messages with path
func (m *CT_FontScheme) ValidateWithPath(path string) error {
	if err := m.MajorFont.ValidateWithPath(path + "/MajorFont"); err != nil {
		return err
	}
	if err := m.MinorFont.ValidateWithPath(path + "/MinorFont"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
