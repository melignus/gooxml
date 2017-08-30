// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_ColorFilter struct {
	// Differential Format Record Id
	DxfIdAttr *uint32
	// Filter By Cell Color
	CellColorAttr *bool
}

func NewCT_ColorFilter() *CT_ColorFilter {
	ret := &CT_ColorFilter{}
	return ret
}
func (m *CT_ColorFilter) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.DxfIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dxfId"},
			Value: fmt.Sprintf("%v", *m.DxfIdAttr)})
	}
	if m.CellColorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cellColor"},
			Value: fmt.Sprintf("%v", *m.CellColorAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_ColorFilter) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "dxfId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DxfIdAttr = &pt
		}
		if attr.Name.Local == "cellColor" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CellColorAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_ColorFilter: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_ColorFilter) Validate() error {
	return m.ValidateWithPath("CT_ColorFilter")
}
func (m *CT_ColorFilter) ValidateWithPath(path string) error {
	return nil
}