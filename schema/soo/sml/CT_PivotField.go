// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@github.com/melignus.

package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/melignus/gooxml"
)

type CT_PivotField struct {
	// Field Name
	NameAttr *string
	// Axis
	AxisAttr ST_Axis
	// Data Field
	DataFieldAttr *bool
	// Custom Subtotal Caption
	SubtotalCaptionAttr *string
	// Show PivotField Header Drop Downs
	ShowDropDownsAttr *bool
	// Hidden Level
	HiddenLevelAttr *bool
	// Unique Member Property
	UniqueMemberPropertyAttr *string
	// Compact
	CompactAttr *bool
	// All Items Expanded
	AllDrilledAttr *bool
	// Number Format Id
	NumFmtIdAttr *uint32
	// Outline Items
	OutlineAttr *bool
	// Subtotals At Top
	SubtotalTopAttr *bool
	// Drag To Row
	DragToRowAttr *bool
	// Drag To Column
	DragToColAttr *bool
	// Multiple Field Filters
	MultipleItemSelectionAllowedAttr *bool
	// Drag Field to Page
	DragToPageAttr *bool
	// Field Can Drag to Data
	DragToDataAttr *bool
	// Drag Off
	DragOffAttr *bool
	// Show All Items
	ShowAllAttr *bool
	// Insert Blank Row
	InsertBlankRowAttr *bool
	// Server-based Page Field
	ServerFieldAttr *bool
	// Insert Item Page Break
	InsertPageBreakAttr *bool
	// Auto Show
	AutoShowAttr *bool
	// Top Auto Show
	TopAutoShowAttr *bool
	// Hide New Items
	HideNewItemsAttr *bool
	// Measure Filter
	MeasureFilterAttr *bool
	// Inclusive Manual Filter
	IncludeNewItemsInFilterAttr *bool
	// Items Per Page Count
	ItemPageCountAttr *uint32
	// Auto Sort Type
	SortTypeAttr ST_FieldSortType
	// Data Source Sort
	DataSourceSortAttr *bool
	// Auto Sort
	NonAutoSortDefaultAttr *bool
	// Auto Show Rank By
	RankByAttr *uint32
	// Show Default Subtotal
	DefaultSubtotalAttr *bool
	// Sum Subtotal
	SumSubtotalAttr *bool
	// CountA
	CountASubtotalAttr *bool
	// Average
	AvgSubtotalAttr *bool
	// Max Subtotal
	MaxSubtotalAttr *bool
	// Min Subtotal
	MinSubtotalAttr *bool
	// Product Subtotal
	ProductSubtotalAttr *bool
	// Count
	CountSubtotalAttr *bool
	// StdDev Subtotal
	StdDevSubtotalAttr *bool
	// StdDevP Subtotal
	StdDevPSubtotalAttr *bool
	// Variance Subtotal
	VarSubtotalAttr *bool
	// VarP Subtotal
	VarPSubtotalAttr *bool
	// Show Member Property in Cell
	ShowPropCellAttr *bool
	// Show Member Property ToolTip
	ShowPropTipAttr *bool
	// Show As Caption
	ShowPropAsCaptionAttr *bool
	// Drill State
	DefaultAttributeDrillStateAttr *bool
	// Field Items
	Items *CT_Items
	// AutoSort Scope
	AutoSortScope *CT_AutoSortScope
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_PivotField() *CT_PivotField {
	ret := &CT_PivotField{}
	return ret
}

func (m *CT_PivotField) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.NameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
			Value: fmt.Sprintf("%v", *m.NameAttr)})
	}
	if m.AxisAttr != ST_AxisUnset {
		attr, err := m.AxisAttr.MarshalXMLAttr(xml.Name{Local: "axis"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.DataFieldAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dataField"},
			Value: fmt.Sprintf("%d", b2i(*m.DataFieldAttr))})
	}
	if m.SubtotalCaptionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "subtotalCaption"},
			Value: fmt.Sprintf("%v", *m.SubtotalCaptionAttr)})
	}
	if m.ShowDropDownsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showDropDowns"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowDropDownsAttr))})
	}
	if m.HiddenLevelAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hiddenLevel"},
			Value: fmt.Sprintf("%d", b2i(*m.HiddenLevelAttr))})
	}
	if m.UniqueMemberPropertyAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "uniqueMemberProperty"},
			Value: fmt.Sprintf("%v", *m.UniqueMemberPropertyAttr)})
	}
	if m.CompactAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "compact"},
			Value: fmt.Sprintf("%d", b2i(*m.CompactAttr))})
	}
	if m.AllDrilledAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "allDrilled"},
			Value: fmt.Sprintf("%d", b2i(*m.AllDrilledAttr))})
	}
	if m.NumFmtIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "numFmtId"},
			Value: fmt.Sprintf("%v", *m.NumFmtIdAttr)})
	}
	if m.OutlineAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "outline"},
			Value: fmt.Sprintf("%d", b2i(*m.OutlineAttr))})
	}
	if m.SubtotalTopAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "subtotalTop"},
			Value: fmt.Sprintf("%d", b2i(*m.SubtotalTopAttr))})
	}
	if m.DragToRowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dragToRow"},
			Value: fmt.Sprintf("%d", b2i(*m.DragToRowAttr))})
	}
	if m.DragToColAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dragToCol"},
			Value: fmt.Sprintf("%d", b2i(*m.DragToColAttr))})
	}
	if m.MultipleItemSelectionAllowedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "multipleItemSelectionAllowed"},
			Value: fmt.Sprintf("%d", b2i(*m.MultipleItemSelectionAllowedAttr))})
	}
	if m.DragToPageAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dragToPage"},
			Value: fmt.Sprintf("%d", b2i(*m.DragToPageAttr))})
	}
	if m.DragToDataAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dragToData"},
			Value: fmt.Sprintf("%d", b2i(*m.DragToDataAttr))})
	}
	if m.DragOffAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dragOff"},
			Value: fmt.Sprintf("%d", b2i(*m.DragOffAttr))})
	}
	if m.ShowAllAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showAll"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowAllAttr))})
	}
	if m.InsertBlankRowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "insertBlankRow"},
			Value: fmt.Sprintf("%d", b2i(*m.InsertBlankRowAttr))})
	}
	if m.ServerFieldAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "serverField"},
			Value: fmt.Sprintf("%d", b2i(*m.ServerFieldAttr))})
	}
	if m.InsertPageBreakAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "insertPageBreak"},
			Value: fmt.Sprintf("%d", b2i(*m.InsertPageBreakAttr))})
	}
	if m.AutoShowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "autoShow"},
			Value: fmt.Sprintf("%d", b2i(*m.AutoShowAttr))})
	}
	if m.TopAutoShowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "topAutoShow"},
			Value: fmt.Sprintf("%d", b2i(*m.TopAutoShowAttr))})
	}
	if m.HideNewItemsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hideNewItems"},
			Value: fmt.Sprintf("%d", b2i(*m.HideNewItemsAttr))})
	}
	if m.MeasureFilterAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "measureFilter"},
			Value: fmt.Sprintf("%d", b2i(*m.MeasureFilterAttr))})
	}
	if m.IncludeNewItemsInFilterAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "includeNewItemsInFilter"},
			Value: fmt.Sprintf("%d", b2i(*m.IncludeNewItemsInFilterAttr))})
	}
	if m.ItemPageCountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "itemPageCount"},
			Value: fmt.Sprintf("%v", *m.ItemPageCountAttr)})
	}
	if m.SortTypeAttr != ST_FieldSortTypeUnset {
		attr, err := m.SortTypeAttr.MarshalXMLAttr(xml.Name{Local: "sortType"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.DataSourceSortAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dataSourceSort"},
			Value: fmt.Sprintf("%d", b2i(*m.DataSourceSortAttr))})
	}
	if m.NonAutoSortDefaultAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "nonAutoSortDefault"},
			Value: fmt.Sprintf("%d", b2i(*m.NonAutoSortDefaultAttr))})
	}
	if m.RankByAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rankBy"},
			Value: fmt.Sprintf("%v", *m.RankByAttr)})
	}
	if m.DefaultSubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "defaultSubtotal"},
			Value: fmt.Sprintf("%d", b2i(*m.DefaultSubtotalAttr))})
	}
	if m.SumSubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sumSubtotal"},
			Value: fmt.Sprintf("%d", b2i(*m.SumSubtotalAttr))})
	}
	if m.CountASubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "countASubtotal"},
			Value: fmt.Sprintf("%d", b2i(*m.CountASubtotalAttr))})
	}
	if m.AvgSubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "avgSubtotal"},
			Value: fmt.Sprintf("%d", b2i(*m.AvgSubtotalAttr))})
	}
	if m.MaxSubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "maxSubtotal"},
			Value: fmt.Sprintf("%d", b2i(*m.MaxSubtotalAttr))})
	}
	if m.MinSubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "minSubtotal"},
			Value: fmt.Sprintf("%d", b2i(*m.MinSubtotalAttr))})
	}
	if m.ProductSubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "productSubtotal"},
			Value: fmt.Sprintf("%d", b2i(*m.ProductSubtotalAttr))})
	}
	if m.CountSubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "countSubtotal"},
			Value: fmt.Sprintf("%d", b2i(*m.CountSubtotalAttr))})
	}
	if m.StdDevSubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "stdDevSubtotal"},
			Value: fmt.Sprintf("%d", b2i(*m.StdDevSubtotalAttr))})
	}
	if m.StdDevPSubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "stdDevPSubtotal"},
			Value: fmt.Sprintf("%d", b2i(*m.StdDevPSubtotalAttr))})
	}
	if m.VarSubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "varSubtotal"},
			Value: fmt.Sprintf("%d", b2i(*m.VarSubtotalAttr))})
	}
	if m.VarPSubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "varPSubtotal"},
			Value: fmt.Sprintf("%d", b2i(*m.VarPSubtotalAttr))})
	}
	if m.ShowPropCellAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showPropCell"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowPropCellAttr))})
	}
	if m.ShowPropTipAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showPropTip"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowPropTipAttr))})
	}
	if m.ShowPropAsCaptionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showPropAsCaption"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowPropAsCaptionAttr))})
	}
	if m.DefaultAttributeDrillStateAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "defaultAttributeDrillState"},
			Value: fmt.Sprintf("%d", b2i(*m.DefaultAttributeDrillStateAttr))})
	}
	e.EncodeToken(start)
	if m.Items != nil {
		seitems := xml.StartElement{Name: xml.Name{Local: "ma:items"}}
		e.EncodeElement(m.Items, seitems)
	}
	if m.AutoSortScope != nil {
		seautoSortScope := xml.StartElement{Name: xml.Name{Local: "ma:autoSortScope"}}
		e.EncodeElement(m.AutoSortScope, seautoSortScope)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PivotField) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "measureFilter" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.MeasureFilterAttr = &parsed
			continue
		}
		if attr.Name.Local == "includeNewItemsInFilter" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.IncludeNewItemsInFilterAttr = &parsed
			continue
		}
		if attr.Name.Local == "compact" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CompactAttr = &parsed
			continue
		}
		if attr.Name.Local == "itemPageCount" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ItemPageCountAttr = &pt
			continue
		}
		if attr.Name.Local == "numFmtId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.NumFmtIdAttr = &pt
			continue
		}
		if attr.Name.Local == "sortType" {
			m.SortTypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "outline" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.OutlineAttr = &parsed
			continue
		}
		if attr.Name.Local == "dataSourceSort" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DataSourceSortAttr = &parsed
			continue
		}
		if attr.Name.Local == "dragToRow" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DragToRowAttr = &parsed
			continue
		}
		if attr.Name.Local == "nonAutoSortDefault" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NonAutoSortDefaultAttr = &parsed
			continue
		}
		if attr.Name.Local == "multipleItemSelectionAllowed" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.MultipleItemSelectionAllowedAttr = &parsed
			continue
		}
		if attr.Name.Local == "rankBy" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.RankByAttr = &pt
			continue
		}
		if attr.Name.Local == "stdDevSubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.StdDevSubtotalAttr = &parsed
			continue
		}
		if attr.Name.Local == "countASubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CountASubtotalAttr = &parsed
			continue
		}
		if attr.Name.Local == "autoShow" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AutoShowAttr = &parsed
			continue
		}
		if attr.Name.Local == "showPropAsCaption" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowPropAsCaptionAttr = &parsed
			continue
		}
		if attr.Name.Local == "productSubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ProductSubtotalAttr = &parsed
			continue
		}
		if attr.Name.Local == "dragToPage" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DragToPageAttr = &parsed
			continue
		}
		if attr.Name.Local == "defaultSubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DefaultSubtotalAttr = &parsed
			continue
		}
		if attr.Name.Local == "dragToData" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DragToDataAttr = &parsed
			continue
		}
		if attr.Name.Local == "maxSubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.MaxSubtotalAttr = &parsed
			continue
		}
		if attr.Name.Local == "hideNewItems" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HideNewItemsAttr = &parsed
			continue
		}
		if attr.Name.Local == "showPropCell" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowPropCellAttr = &parsed
			continue
		}
		if attr.Name.Local == "subtotalTop" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SubtotalTopAttr = &parsed
			continue
		}
		if attr.Name.Local == "serverField" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ServerFieldAttr = &parsed
			continue
		}
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = &parsed
			continue
		}
		if attr.Name.Local == "insertPageBreak" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.InsertPageBreakAttr = &parsed
			continue
		}
		if attr.Name.Local == "sumSubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SumSubtotalAttr = &parsed
			continue
		}
		if attr.Name.Local == "dataField" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DataFieldAttr = &parsed
			continue
		}
		if attr.Name.Local == "avgSubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AvgSubtotalAttr = &parsed
			continue
		}
		if attr.Name.Local == "topAutoShow" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.TopAutoShowAttr = &parsed
			continue
		}
		if attr.Name.Local == "minSubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.MinSubtotalAttr = &parsed
			continue
		}
		if attr.Name.Local == "dragOff" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DragOffAttr = &parsed
			continue
		}
		if attr.Name.Local == "countSubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CountSubtotalAttr = &parsed
			continue
		}
		if attr.Name.Local == "uniqueMemberProperty" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.UniqueMemberPropertyAttr = &parsed
			continue
		}
		if attr.Name.Local == "varPSubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.VarPSubtotalAttr = &parsed
			continue
		}
		if attr.Name.Local == "stdDevPSubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.StdDevPSubtotalAttr = &parsed
			continue
		}
		if attr.Name.Local == "subtotalCaption" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SubtotalCaptionAttr = &parsed
			continue
		}
		if attr.Name.Local == "showPropTip" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowPropTipAttr = &parsed
			continue
		}
		if attr.Name.Local == "insertBlankRow" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.InsertBlankRowAttr = &parsed
			continue
		}
		if attr.Name.Local == "defaultAttributeDrillState" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DefaultAttributeDrillStateAttr = &parsed
			continue
		}
		if attr.Name.Local == "dragToCol" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DragToColAttr = &parsed
			continue
		}
		if attr.Name.Local == "hiddenLevel" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HiddenLevelAttr = &parsed
			continue
		}
		if attr.Name.Local == "axis" {
			m.AxisAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "showDropDowns" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowDropDownsAttr = &parsed
			continue
		}
		if attr.Name.Local == "showAll" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowAllAttr = &parsed
			continue
		}
		if attr.Name.Local == "allDrilled" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AllDrilledAttr = &parsed
			continue
		}
		if attr.Name.Local == "varSubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.VarSubtotalAttr = &parsed
			continue
		}
	}
lCT_PivotField:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "items"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "items"}:
				m.Items = NewCT_Items()
				if err := d.DecodeElement(m.Items, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "autoSortScope"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "autoSortScope"}:
				m.AutoSortScope = NewCT_AutoSortScope()
				if err := d.DecodeElement(m.AutoSortScope, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_PivotField %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PivotField
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PivotField and its children
func (m *CT_PivotField) Validate() error {
	return m.ValidateWithPath("CT_PivotField")
}

// ValidateWithPath validates the CT_PivotField and its children, prefixing error messages with path
func (m *CT_PivotField) ValidateWithPath(path string) error {
	if err := m.AxisAttr.ValidateWithPath(path + "/AxisAttr"); err != nil {
		return err
	}
	if err := m.SortTypeAttr.ValidateWithPath(path + "/SortTypeAttr"); err != nil {
		return err
	}
	if m.Items != nil {
		if err := m.Items.ValidateWithPath(path + "/Items"); err != nil {
			return err
		}
	}
	if m.AutoSortScope != nil {
		if err := m.AutoSortScope.ValidateWithPath(path + "/AutoSortScope"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
