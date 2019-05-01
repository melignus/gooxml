// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@github.com/melignus.

package pml

import (
	"encoding/xml"
	"fmt"
)

type CT_Kinsoku struct {
	// Language
	LangAttr *string
	// Invalid Kinsoku Start Characters
	InvalStCharsAttr string
	// Invalid Kinsoku End Characters
	InvalEndCharsAttr string
}

func NewCT_Kinsoku() *CT_Kinsoku {
	ret := &CT_Kinsoku{}
	return ret
}

func (m *CT_Kinsoku) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.LangAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "lang"},
			Value: fmt.Sprintf("%v", *m.LangAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "invalStChars"},
		Value: fmt.Sprintf("%v", m.InvalStCharsAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "invalEndChars"},
		Value: fmt.Sprintf("%v", m.InvalEndCharsAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Kinsoku) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "lang" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LangAttr = &parsed
			continue
		}
		if attr.Name.Local == "invalStChars" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.InvalStCharsAttr = parsed
			continue
		}
		if attr.Name.Local == "invalEndChars" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.InvalEndCharsAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Kinsoku: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Kinsoku and its children
func (m *CT_Kinsoku) Validate() error {
	return m.ValidateWithPath("CT_Kinsoku")
}

// ValidateWithPath validates the CT_Kinsoku and its children, prefixing error messages with path
func (m *CT_Kinsoku) ValidateWithPath(path string) error {
	return nil
}
