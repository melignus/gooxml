// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"log"
)

type EG_TextAutofit struct {
	NoAutofit   *CT_TextNoAutofit
	NormAutofit *CT_TextNormalAutofit
	SpAutoFit   *CT_TextShapeAutofit
}

func NewEG_TextAutofit() *EG_TextAutofit {
	ret := &EG_TextAutofit{}
	return ret
}
func (m *EG_TextAutofit) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.NoAutofit != nil {
		senoAutofit := xml.StartElement{Name: xml.Name{Local: "a:noAutofit"}}
		e.EncodeElement(m.NoAutofit, senoAutofit)
	}
	if m.NormAutofit != nil {
		senormAutofit := xml.StartElement{Name: xml.Name{Local: "a:normAutofit"}}
		e.EncodeElement(m.NormAutofit, senormAutofit)
	}
	if m.SpAutoFit != nil {
		sespAutoFit := xml.StartElement{Name: xml.Name{Local: "a:spAutoFit"}}
		e.EncodeElement(m.SpAutoFit, sespAutoFit)
	}
	return nil
}
func (m *EG_TextAutofit) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_TextAutofit:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "noAutofit":
				m.NoAutofit = NewCT_TextNoAutofit()
				if err := d.DecodeElement(m.NoAutofit, &el); err != nil {
					return err
				}
			case "normAutofit":
				m.NormAutofit = NewCT_TextNormalAutofit()
				if err := d.DecodeElement(m.NormAutofit, &el); err != nil {
					return err
				}
			case "spAutoFit":
				m.SpAutoFit = NewCT_TextShapeAutofit()
				if err := d.DecodeElement(m.SpAutoFit, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_TextAutofit
		case xml.CharData:
		}
	}
	return nil
}
func (m *EG_TextAutofit) Validate() error {
	return m.ValidateWithPath("EG_TextAutofit")
}
func (m *EG_TextAutofit) ValidateWithPath(path string) error {
	if m.NoAutofit != nil {
		if err := m.NoAutofit.ValidateWithPath(path + "/NoAutofit"); err != nil {
			return err
		}
	}
	if m.NormAutofit != nil {
		if err := m.NormAutofit.ValidateWithPath(path + "/NormAutofit"); err != nil {
			return err
		}
	}
	if m.SpAutoFit != nil {
		if err := m.SpAutoFit.ValidateWithPath(path + "/SpAutoFit"); err != nil {
			return err
		}
	}
	return nil
}