// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@github.com/melignus.

package wml

import (
	"encoding/xml"
	"fmt"

	"github.com/melignus/gooxml"
	"github.com/melignus/gooxml/schema/soo/dml"
)

type WdCT_WordprocessingContentPart struct {
	BwModeAttr      dml.ST_BlackWhiteMode
	IdAttr          string
	NvContentPartPr *WdCT_WordprocessingContentPartNonVisual
	Xfrm            *dml.CT_Transform2D
	ExtLst          *dml.CT_OfficeArtExtensionList
}

func NewWdCT_WordprocessingContentPart() *WdCT_WordprocessingContentPart {
	ret := &WdCT_WordprocessingContentPart{}
	return ret
}

func (m *WdCT_WordprocessingContentPart) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.BwModeAttr != dml.ST_BlackWhiteModeUnset {
		attr, err := m.BwModeAttr.MarshalXMLAttr(xml.Name{Local: "bwMode"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	e.EncodeToken(start)
	if m.NvContentPartPr != nil {
		senvContentPartPr := xml.StartElement{Name: xml.Name{Local: "wp:nvContentPartPr"}}
		e.EncodeElement(m.NvContentPartPr, senvContentPartPr)
	}
	if m.Xfrm != nil {
		sexfrm := xml.StartElement{Name: xml.Name{Local: "wp:xfrm"}}
		e.EncodeElement(m.Xfrm, sexfrm)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "wp:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *WdCT_WordprocessingContentPart) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "id" ||
			attr.Name.Space == "http://purl.oclc.org/ooxml/officeDocument/relationships" && attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = parsed
			continue
		}
		if attr.Name.Local == "bwMode" {
			m.BwModeAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lWdCT_WordprocessingContentPart:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "nvContentPartPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "nvContentPartPr"}:
				m.NvContentPartPr = NewWdCT_WordprocessingContentPartNonVisual()
				if err := d.DecodeElement(m.NvContentPartPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "xfrm"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "xfrm"}:
				m.Xfrm = dml.NewCT_Transform2D()
				if err := d.DecodeElement(m.Xfrm, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "extLst"}:
				m.ExtLst = dml.NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on WdCT_WordprocessingContentPart %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lWdCT_WordprocessingContentPart
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the WdCT_WordprocessingContentPart and its children
func (m *WdCT_WordprocessingContentPart) Validate() error {
	return m.ValidateWithPath("WdCT_WordprocessingContentPart")
}

// ValidateWithPath validates the WdCT_WordprocessingContentPart and its children, prefixing error messages with path
func (m *WdCT_WordprocessingContentPart) ValidateWithPath(path string) error {
	if err := m.BwModeAttr.ValidateWithPath(path + "/BwModeAttr"); err != nil {
		return err
	}
	if m.NvContentPartPr != nil {
		if err := m.NvContentPartPr.ValidateWithPath(path + "/NvContentPartPr"); err != nil {
			return err
		}
	}
	if m.Xfrm != nil {
		if err := m.Xfrm.ValidateWithPath(path + "/Xfrm"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
