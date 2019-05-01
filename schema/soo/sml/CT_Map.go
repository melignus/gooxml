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

type CT_Map struct {
	// XML Mapping ID
	IDAttr uint32
	// XML Mapping Name
	NameAttr string
	// Root Element Name
	RootElementAttr string
	// Schema Name
	SchemaIDAttr string
	// Show Validation Errors
	ShowImportExportValidationErrorsAttr bool
	// AutoFit Table on Refresh
	AutoFitAttr bool
	// Append Data to Table
	AppendAttr bool
	// Preserve AutoFilter State
	PreserveSortAFLayoutAttr bool
	// Preserve Cell Formatting
	PreserveFormatAttr bool
	// XML Mapping
	DataBinding *CT_DataBinding
}

func NewCT_Map() *CT_Map {
	ret := &CT_Map{}
	return ret
}

func (m *CT_Map) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ID"},
		Value: fmt.Sprintf("%v", m.IDAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "Name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "RootElement"},
		Value: fmt.Sprintf("%v", m.RootElementAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "SchemaID"},
		Value: fmt.Sprintf("%v", m.SchemaIDAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ShowImportExportValidationErrors"},
		Value: fmt.Sprintf("%d", b2i(m.ShowImportExportValidationErrorsAttr))})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "AutoFit"},
		Value: fmt.Sprintf("%d", b2i(m.AutoFitAttr))})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "Append"},
		Value: fmt.Sprintf("%d", b2i(m.AppendAttr))})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "PreserveSortAFLayout"},
		Value: fmt.Sprintf("%d", b2i(m.PreserveSortAFLayoutAttr))})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "PreserveFormat"},
		Value: fmt.Sprintf("%d", b2i(m.PreserveFormatAttr))})
	e.EncodeToken(start)
	if m.DataBinding != nil {
		seDataBinding := xml.StartElement{Name: xml.Name{Local: "ma:DataBinding"}}
		e.EncodeElement(m.DataBinding, seDataBinding)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Map) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "ID" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.IDAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "Name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
			continue
		}
		if attr.Name.Local == "RootElement" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RootElementAttr = parsed
			continue
		}
		if attr.Name.Local == "SchemaID" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SchemaIDAttr = parsed
			continue
		}
		if attr.Name.Local == "ShowImportExportValidationErrors" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowImportExportValidationErrorsAttr = parsed
			continue
		}
		if attr.Name.Local == "AutoFit" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AutoFitAttr = parsed
			continue
		}
		if attr.Name.Local == "Append" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AppendAttr = parsed
			continue
		}
		if attr.Name.Local == "PreserveSortAFLayout" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PreserveSortAFLayoutAttr = parsed
			continue
		}
		if attr.Name.Local == "PreserveFormat" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PreserveFormatAttr = parsed
			continue
		}
	}
lCT_Map:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "DataBinding"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "DataBinding"}:
				m.DataBinding = NewCT_DataBinding()
				if err := d.DecodeElement(m.DataBinding, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_Map %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Map
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Map and its children
func (m *CT_Map) Validate() error {
	return m.ValidateWithPath("CT_Map")
}

// ValidateWithPath validates the CT_Map and its children, prefixing error messages with path
func (m *CT_Map) ValidateWithPath(path string) error {
	if m.DataBinding != nil {
		if err := m.DataBinding.ValidateWithPath(path + "/DataBinding"); err != nil {
			return err
		}
	}
	return nil
}
