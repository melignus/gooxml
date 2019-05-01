// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@github.com/melignus.

package sml

import (
	"encoding/xml"

	"github.com/melignus/gooxml"
)

type CT_FutureMetadataBlock struct {
	// Future Feature Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_FutureMetadataBlock() *CT_FutureMetadataBlock {
	ret := &CT_FutureMetadataBlock{}
	return ret
}

func (m *CT_FutureMetadataBlock) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_FutureMetadataBlock) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_FutureMetadataBlock:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_FutureMetadataBlock %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_FutureMetadataBlock
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_FutureMetadataBlock and its children
func (m *CT_FutureMetadataBlock) Validate() error {
	return m.ValidateWithPath("CT_FutureMetadataBlock")
}

// ValidateWithPath validates the CT_FutureMetadataBlock and its children, prefixing error messages with path
func (m *CT_FutureMetadataBlock) ValidateWithPath(path string) error {
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
