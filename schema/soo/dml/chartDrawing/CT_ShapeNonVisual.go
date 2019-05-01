// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@github.com/melignus.

package chartDrawing

import (
	"encoding/xml"

	"github.com/melignus/gooxml"
	"github.com/melignus/gooxml/schema/soo/dml"
)

type CT_ShapeNonVisual struct {
	CNvPr   *dml.CT_NonVisualDrawingProps
	CNvSpPr *dml.CT_NonVisualDrawingShapeProps
}

func NewCT_ShapeNonVisual() *CT_ShapeNonVisual {
	ret := &CT_ShapeNonVisual{}
	ret.CNvPr = dml.NewCT_NonVisualDrawingProps()
	ret.CNvSpPr = dml.NewCT_NonVisualDrawingShapeProps()
	return ret
}

func (m *CT_ShapeNonVisual) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	secNvPr := xml.StartElement{Name: xml.Name{Local: "cNvPr"}}
	e.EncodeElement(m.CNvPr, secNvPr)
	secNvSpPr := xml.StartElement{Name: xml.Name{Local: "cNvSpPr"}}
	e.EncodeElement(m.CNvSpPr, secNvSpPr)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ShapeNonVisual) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CNvPr = dml.NewCT_NonVisualDrawingProps()
	m.CNvSpPr = dml.NewCT_NonVisualDrawingShapeProps()
lCT_ShapeNonVisual:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", Local: "cNvPr"}:
				if err := d.DecodeElement(m.CNvPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", Local: "cNvSpPr"}:
				if err := d.DecodeElement(m.CNvSpPr, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_ShapeNonVisual %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ShapeNonVisual
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ShapeNonVisual and its children
func (m *CT_ShapeNonVisual) Validate() error {
	return m.ValidateWithPath("CT_ShapeNonVisual")
}

// ValidateWithPath validates the CT_ShapeNonVisual and its children, prefixing error messages with path
func (m *CT_ShapeNonVisual) ValidateWithPath(path string) error {
	if err := m.CNvPr.ValidateWithPath(path + "/CNvPr"); err != nil {
		return err
	}
	if err := m.CNvSpPr.ValidateWithPath(path + "/CNvSpPr"); err != nil {
		return err
	}
	return nil
}
