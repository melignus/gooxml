// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@github.com/melignus.

package vml_test

import (
	"encoding/xml"
	"testing"

	"github.com/melignus/gooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestCurveConstructor(t *testing.T) {
	v := vml.NewCurve()
	if v == nil {
		t.Errorf("vml.NewCurve must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.Curve should validate: %s", err)
	}
}

func TestCurveMarshalUnmarshal(t *testing.T) {
	v := vml.NewCurve()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewCurve()
	xml.Unmarshal(buf, v2)
}
