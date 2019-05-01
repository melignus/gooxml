// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@github.com/melignus.

package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/melignus/gooxml/schema/soo/sml"
)

func TestCT_BreakConstructor(t *testing.T) {
	v := sml.NewCT_Break()
	if v == nil {
		t.Errorf("sml.NewCT_Break must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Break should validate: %s", err)
	}
}

func TestCT_BreakMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Break()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Break()
	xml.Unmarshal(buf, v2)
}
