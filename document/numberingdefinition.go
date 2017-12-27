// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import "baliance.com/gooxml/schema/soo/wml"

// NumberingDefinition defines a numbering definition for a list of pragraphs.
type NumberingDefinition struct {
	x *wml.CT_Num
}

// X returns the inner wrapped XML type.
func (n NumberingDefinition) X() *wml.CT_Num {
	return n.x
}

// AbstractNumberID returns the ID that is unique within all numbering
// definitions that is used to assign the definition to a paragraph.
func (n NumberingDefinition) AbstractNumberID() int64 {
	if n.x.AbstractNumId == nil {
		return 0
	}
	return n.x.AbstractNumId.ValAttr
}

// Levels returns all of the numbering levels defined in the definition.
func (n NumberingDefinition) Levels() []NumberingLevel {
	ret := []NumberingLevel{}
	for _, nl := range n.x.LvlOverride {
		ret = append(ret, NumberingLevel{nl})
	}
	return ret
}

// AddLevel adds a new numbering level to a NumberingDefinition.
func (n NumberingDefinition) AddLevel() NumberingLevel {
	nl := wml.NewCT_NumLvl()
	nl.IlvlAttr = int64(len(n.x.LvlOverride))
	n.x.LvlOverride = append(n.x.LvlOverride, nl)
	return NumberingLevel{nl}
}