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
	"fmt"
	"strconv"

	"github.com/melignus/gooxml"
)

type CT_RowFields struct {
	// Repeated Items Count
	CountAttr *uint32
	// Row Items
	Field []*CT_Field
}

func NewCT_RowFields() *CT_RowFields {
	ret := &CT_RowFields{}
	return ret
}

func (m *CT_RowFields) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	sefield := xml.StartElement{Name: xml.Name{Local: "ma:field"}}
	for _, c := range m.Field {
		e.EncodeElement(c, sefield)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_RowFields) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
			continue
		}
	}
lCT_RowFields:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "field"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "field"}:
				tmp := NewCT_Field()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Field = append(m.Field, tmp)
			default:
				gooxml.Log("skipping unsupported element on CT_RowFields %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_RowFields
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_RowFields and its children
func (m *CT_RowFields) Validate() error {
	return m.ValidateWithPath("CT_RowFields")
}

// ValidateWithPath validates the CT_RowFields and its children, prefixing error messages with path
func (m *CT_RowFields) ValidateWithPath(path string) error {
	for i, v := range m.Field {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Field[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
