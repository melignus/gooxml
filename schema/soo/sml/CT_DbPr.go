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
)

type CT_DbPr struct {
	// Connection String
	ConnectionAttr string
	// Command Text
	CommandAttr *string
	// Command Text
	ServerCommandAttr *string
	// OLE DB Command Type
	CommandTypeAttr *uint32
}

func NewCT_DbPr() *CT_DbPr {
	ret := &CT_DbPr{}
	return ret
}

func (m *CT_DbPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "connection"},
		Value: fmt.Sprintf("%v", m.ConnectionAttr)})
	if m.CommandAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "command"},
			Value: fmt.Sprintf("%v", *m.CommandAttr)})
	}
	if m.ServerCommandAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "serverCommand"},
			Value: fmt.Sprintf("%v", *m.ServerCommandAttr)})
	}
	if m.CommandTypeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "commandType"},
			Value: fmt.Sprintf("%v", *m.CommandTypeAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DbPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "connection" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ConnectionAttr = parsed
			continue
		}
		if attr.Name.Local == "command" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CommandAttr = &parsed
			continue
		}
		if attr.Name.Local == "serverCommand" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ServerCommandAttr = &parsed
			continue
		}
		if attr.Name.Local == "commandType" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CommandTypeAttr = &pt
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_DbPr: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_DbPr and its children
func (m *CT_DbPr) Validate() error {
	return m.ValidateWithPath("CT_DbPr")
}

// ValidateWithPath validates the CT_DbPr and its children, prefixing error messages with path
func (m *CT_DbPr) ValidateWithPath(path string) error {
	return nil
}
