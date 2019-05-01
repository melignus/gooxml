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

type CT_Format struct {
	// Format Action
	ActionAttr ST_FormatAction
	// Format Id
	DxfIdAttr *uint32
	// Pivot Table Location
	PivotArea *CT_PivotArea
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_Format() *CT_Format {
	ret := &CT_Format{}
	ret.PivotArea = NewCT_PivotArea()
	return ret
}

func (m *CT_Format) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ActionAttr != ST_FormatActionUnset {
		attr, err := m.ActionAttr.MarshalXMLAttr(xml.Name{Local: "action"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.DxfIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dxfId"},
			Value: fmt.Sprintf("%v", *m.DxfIdAttr)})
	}
	e.EncodeToken(start)
	sepivotArea := xml.StartElement{Name: xml.Name{Local: "ma:pivotArea"}}
	e.EncodeElement(m.PivotArea, sepivotArea)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Format) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.PivotArea = NewCT_PivotArea()
	for _, attr := range start.Attr {
		if attr.Name.Local == "action" {
			m.ActionAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "dxfId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DxfIdAttr = &pt
			continue
		}
	}
lCT_Format:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "pivotArea"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "pivotArea"}:
				if err := d.DecodeElement(m.PivotArea, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_Format %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Format
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Format and its children
func (m *CT_Format) Validate() error {
	return m.ValidateWithPath("CT_Format")
}

// ValidateWithPath validates the CT_Format and its children, prefixing error messages with path
func (m *CT_Format) ValidateWithPath(path string) error {
	if err := m.ActionAttr.ValidateWithPath(path + "/ActionAttr"); err != nil {
		return err
	}
	if err := m.PivotArea.ValidateWithPath(path + "/PivotArea"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
