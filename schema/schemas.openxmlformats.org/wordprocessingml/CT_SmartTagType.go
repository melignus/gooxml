// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"
)

type CT_SmartTagType struct {
	// Smart Tag Namespace
	NamespaceuriAttr *string
	// Smart Tag Name
	NameAttr *string
	// Smart Tag Supplementary URL
	UrlAttr *string
}

func NewCT_SmartTagType() *CT_SmartTagType {
	ret := &CT_SmartTagType{}
	return ret
}
func (m *CT_SmartTagType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.NamespaceuriAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:namespaceuri"},
			Value: fmt.Sprintf("%v", *m.NamespaceuriAttr)})
	}
	if m.NameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:name"},
			Value: fmt.Sprintf("%v", *m.NameAttr)})
	}
	if m.UrlAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:url"},
			Value: fmt.Sprintf("%v", *m.UrlAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_SmartTagType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "namespaceuri" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NamespaceuriAttr = &parsed
		}
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = &parsed
		}
		if attr.Name.Local == "url" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.UrlAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_SmartTagType: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_SmartTagType) Validate() error {
	return m.ValidateWithPath("CT_SmartTagType")
}
func (m *CT_SmartTagType) ValidateWithPath(path string) error {
	return nil
}