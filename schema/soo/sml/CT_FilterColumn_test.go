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

func TestCT_FilterColumnConstructor(t *testing.T) {
	v := sml.NewCT_FilterColumn()
	if v == nil {
		t.Errorf("sml.NewCT_FilterColumn must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_FilterColumn should validate: %s", err)
	}
}

func TestCT_FilterColumnMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_FilterColumn()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_FilterColumn()
	xml.Unmarshal(buf, v2)
}
