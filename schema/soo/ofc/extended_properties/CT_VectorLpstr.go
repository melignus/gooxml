// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@github.com/melignus.

package extended_properties

import (
	"encoding/xml"

	"github.com/melignus/gooxml"
	"github.com/melignus/gooxml/schema/soo/ofc/docPropsVTypes"
)

type CT_VectorLpstr struct {
	Vector *docPropsVTypes.Vector
}

func NewCT_VectorLpstr() *CT_VectorLpstr {
	ret := &CT_VectorLpstr{}
	ret.Vector = docPropsVTypes.NewVector()
	return ret
}

func (m *CT_VectorLpstr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	sevector := xml.StartElement{Name: xml.Name{Local: "vt:vector"}}
	e.EncodeElement(m.Vector, sevector)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_VectorLpstr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Vector = docPropsVTypes.NewVector()
lCT_VectorLpstr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "vector"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "vector"}:
				if err := d.DecodeElement(m.Vector, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_VectorLpstr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_VectorLpstr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_VectorLpstr and its children
func (m *CT_VectorLpstr) Validate() error {
	return m.ValidateWithPath("CT_VectorLpstr")
}

// ValidateWithPath validates the CT_VectorLpstr and its children, prefixing error messages with path
func (m *CT_VectorLpstr) ValidateWithPath(path string) error {
	if err := m.Vector.ValidateWithPath(path + "/Vector"); err != nil {
		return err
	}
	return nil
}
