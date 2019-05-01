// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@github.com/melignus.

package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/melignus/gooxml/schema/soo/wml"
)

func TestCT_CaptionsConstructor(t *testing.T) {
	v := wml.NewCT_Captions()
	if v == nil {
		t.Errorf("wml.NewCT_Captions must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Captions should validate: %s", err)
	}
}

func TestCT_CaptionsMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Captions()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Captions()
	xml.Unmarshal(buf, v2)
}
