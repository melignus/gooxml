// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@github.com/melignus.

package chart

import (
	"encoding/xml"
	"fmt"
)

type CT_DepthPercent struct {
	ValAttr *ST_DepthPercent
}

func NewCT_DepthPercent() *CT_DepthPercent {
	ret := &CT_DepthPercent{}
	return ret
}

func (m *CT_DepthPercent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ValAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
			Value: fmt.Sprintf("%v", *m.ValAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DepthPercent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := ParseUnionST_DepthPercent(attr.Value)
			if err != nil {
				return err
			}
			m.ValAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_DepthPercent: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_DepthPercent and its children
func (m *CT_DepthPercent) Validate() error {
	return m.ValidateWithPath("CT_DepthPercent")
}

// ValidateWithPath validates the CT_DepthPercent and its children, prefixing error messages with path
func (m *CT_DepthPercent) ValidateWithPath(path string) error {
	if m.ValAttr != nil {
		if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
			return err
		}
	}
	return nil
}
