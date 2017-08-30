// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingDrawing

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
)

type CT_WrapThrough struct {
	WrapTextAttr ST_WrapText
	DistLAttr    *uint32
	DistRAttr    *uint32
	WrapPolygon  *CT_WrapPath
}

func NewCT_WrapThrough() *CT_WrapThrough {
	ret := &CT_WrapThrough{}
	ret.WrapTextAttr = ST_WrapText(1)
	ret.WrapPolygon = NewCT_WrapPath()
	return ret
}
func (m *CT_WrapThrough) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	attr, err := m.WrapTextAttr.MarshalXMLAttr(xml.Name{Local: "wrapText"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	if m.DistLAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distL"},
			Value: fmt.Sprintf("%v", *m.DistLAttr)})
	}
	if m.DistRAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distR"},
			Value: fmt.Sprintf("%v", *m.DistRAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	sewrapPolygon := xml.StartElement{Name: xml.Name{Local: "wp:wrapPolygon"}}
	e.EncodeElement(m.WrapPolygon, sewrapPolygon)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_WrapThrough) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.WrapTextAttr = ST_WrapText(1)
	m.WrapPolygon = NewCT_WrapPath()
	for _, attr := range start.Attr {
		if attr.Name.Local == "wrapText" {
			m.WrapTextAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "distL" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DistLAttr = &pt
		}
		if attr.Name.Local == "distR" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DistRAttr = &pt
		}
	}
lCT_WrapThrough:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "wrapPolygon":
				if err := d.DecodeElement(m.WrapPolygon, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_WrapThrough
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_WrapThrough) Validate() error {
	return m.ValidateWithPath("CT_WrapThrough")
}
func (m *CT_WrapThrough) ValidateWithPath(path string) error {
	if m.WrapTextAttr == ST_WrapTextUnset {
		return fmt.Errorf("%s/WrapTextAttr is a mandatory field", path)
	}
	if err := m.WrapTextAttr.ValidateWithPath(path + "/WrapTextAttr"); err != nil {
		return err
	}
	if err := m.WrapPolygon.ValidateWithPath(path + "/WrapPolygon"); err != nil {
		return err
	}
	return nil
}