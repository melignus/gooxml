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

func TestCT_ProtectedRangesConstructor(t *testing.T) {
	v := sml.NewCT_ProtectedRanges()
	if v == nil {
		t.Errorf("sml.NewCT_ProtectedRanges must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ProtectedRanges should validate: %s", err)
	}
}

func TestCT_ProtectedRangesMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ProtectedRanges()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ProtectedRanges()
	xml.Unmarshal(buf, v2)
}
