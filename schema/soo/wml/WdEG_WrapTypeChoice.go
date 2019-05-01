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

	"github.com/melignus/gooxml"
)

type WdEG_WrapTypeChoice struct {
	WrapNone         *WdCT_WrapNone
	WrapSquare       *WdCT_WrapSquare
	WrapTight        *WdCT_WrapTight
	WrapThrough      *WdCT_WrapThrough
	WrapTopAndBottom *WdCT_WrapTopBottom
}

func NewWdEG_WrapTypeChoice() *WdEG_WrapTypeChoice {
	ret := &WdEG_WrapTypeChoice{}
	return ret
}

func (m *WdEG_WrapTypeChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.WrapNone != nil {
		sewrapNone := xml.StartElement{Name: xml.Name{Local: "wp:wrapNone"}}
		e.EncodeElement(m.WrapNone, sewrapNone)
	}
	if m.WrapSquare != nil {
		sewrapSquare := xml.StartElement{Name: xml.Name{Local: "wp:wrapSquare"}}
		e.EncodeElement(m.WrapSquare, sewrapSquare)
	}
	if m.WrapTight != nil {
		sewrapTight := xml.StartElement{Name: xml.Name{Local: "wp:wrapTight"}}
		e.EncodeElement(m.WrapTight, sewrapTight)
	}
	if m.WrapThrough != nil {
		sewrapThrough := xml.StartElement{Name: xml.Name{Local: "wp:wrapThrough"}}
		e.EncodeElement(m.WrapThrough, sewrapThrough)
	}
	if m.WrapTopAndBottom != nil {
		sewrapTopAndBottom := xml.StartElement{Name: xml.Name{Local: "wp:wrapTopAndBottom"}}
		e.EncodeElement(m.WrapTopAndBottom, sewrapTopAndBottom)
	}
	return nil
}

func (m *WdEG_WrapTypeChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lWdEG_WrapTypeChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "wrapNone"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "wrapNone"}:
				m.WrapNone = NewWdCT_WrapNone()
				if err := d.DecodeElement(m.WrapNone, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "wrapSquare"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "wrapSquare"}:
				m.WrapSquare = NewWdCT_WrapSquare()
				if err := d.DecodeElement(m.WrapSquare, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "wrapTight"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "wrapTight"}:
				m.WrapTight = NewWdCT_WrapTight()
				if err := d.DecodeElement(m.WrapTight, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "wrapThrough"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "wrapThrough"}:
				m.WrapThrough = NewWdCT_WrapThrough()
				if err := d.DecodeElement(m.WrapThrough, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "wrapTopAndBottom"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "wrapTopAndBottom"}:
				m.WrapTopAndBottom = NewWdCT_WrapTopBottom()
				if err := d.DecodeElement(m.WrapTopAndBottom, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on WdEG_WrapTypeChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lWdEG_WrapTypeChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the WdEG_WrapTypeChoice and its children
func (m *WdEG_WrapTypeChoice) Validate() error {
	return m.ValidateWithPath("WdEG_WrapTypeChoice")
}

// ValidateWithPath validates the WdEG_WrapTypeChoice and its children, prefixing error messages with path
func (m *WdEG_WrapTypeChoice) ValidateWithPath(path string) error {
	if m.WrapNone != nil {
		if err := m.WrapNone.ValidateWithPath(path + "/WrapNone"); err != nil {
			return err
		}
	}
	if m.WrapSquare != nil {
		if err := m.WrapSquare.ValidateWithPath(path + "/WrapSquare"); err != nil {
			return err
		}
	}
	if m.WrapTight != nil {
		if err := m.WrapTight.ValidateWithPath(path + "/WrapTight"); err != nil {
			return err
		}
	}
	if m.WrapThrough != nil {
		if err := m.WrapThrough.ValidateWithPath(path + "/WrapThrough"); err != nil {
			return err
		}
	}
	if m.WrapTopAndBottom != nil {
		if err := m.WrapTopAndBottom.ValidateWithPath(path + "/WrapTopAndBottom"); err != nil {
			return err
		}
	}
	return nil
}
