package MD_Builder

import (
	"fmt"
	"strings"
)

func New_List() *MD_List {
	return &MD_List{
		Items:     []string{},
		level_idx: make([]int, 100),
	}
}

func (ol *MD_List) Fill_List(a []MD_List_Item, is_continuous bool) {
	non_continuous_idx := 0
	for i, item := range a {
		if item.Ordering == "ordered" {
			if i > 0 {
				if item.Level > a[i-1].Level {
					ol.level_idx[item.Level-1] = 0
				}
			}
			if is_continuous {
				ol.Items = append(ol.Items, strings.Repeat("\t", item.Level-1)+fmt.Sprintf("%d. ", i+1)+item.Item)
			} else {
				ol.level_idx[item.Level-1]++
				ol.Items = append(ol.Items, strings.Repeat("\t", item.Level-1)+fmt.Sprintf("%d. ", ol.level_idx[item.Level-1])+item.Item)
				non_continuous_idx++
			}

		} else if item.Ordering == "unordered" {
			if !is_continuous {
				non_continuous_idx++
			}
			ol.Items = append(ol.Items, strings.Repeat("\t", item.Level-1)+"- "+item.Item)
		}
	}
}
