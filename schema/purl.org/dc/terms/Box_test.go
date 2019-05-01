// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@github.com/melignus.

package terms_test

import (
	"encoding/xml"
	"testing"

	"github.com/melignus/gooxml/schema/purl.org/dc/terms"
)

func TestBoxConstructor(t *testing.T) {
	v := terms.NewBox()
	if v == nil {
		t.Errorf("terms.NewBox must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.Box should validate: %s", err)
	}
}

func TestBoxMarshalUnmarshal(t *testing.T) {
	v := terms.NewBox()
	buf, _ := xml.Marshal(v)
	v2 := terms.NewBox()
	xml.Unmarshal(buf, v2)
}
