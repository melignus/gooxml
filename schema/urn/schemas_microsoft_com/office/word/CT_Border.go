// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@github.com/melignus.

package word

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_Border struct {
	TypeAttr   ST_BorderType
	WidthAttr  *uint32
	ShadowAttr ST_BorderShadow
}

func NewCT_Border() *CT_Border {
	ret := &CT_Border{}
	return ret
}

func (m *CT_Border) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.TypeAttr != ST_BorderTypeUnset {
		attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.WidthAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "width"},
			Value: fmt.Sprintf("%v", *m.WidthAttr)})
	}
	if m.ShadowAttr != ST_BorderShadowUnset {
		attr, err := m.ShadowAttr.MarshalXMLAttr(xml.Name{Local: "shadow"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Border) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "width" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.WidthAttr = &pt
			continue
		}
		if attr.Name.Local == "shadow" {
			m.ShadowAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Border: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Border and its children
func (m *CT_Border) Validate() error {
	return m.ValidateWithPath("CT_Border")
}

// ValidateWithPath validates the CT_Border and its children, prefixing error messages with path
func (m *CT_Border) ValidateWithPath(path string) error {
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	if err := m.ShadowAttr.ValidateWithPath(path + "/ShadowAttr"); err != nil {
		return err
	}
	return nil
}
