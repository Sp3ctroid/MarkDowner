package MD_Builder

import "fmt"

func New_List() *MD_List {
	return &MD_List{
		Items: []string{},
	}
}

func (ol *MD_List) Fill_List(a ...interface{}) {
	for _, item := range a {
		ol.Items = append(ol.Items, item.(string))
	}
}

func (md *MD) Ordered_List(ol *MD_List) {
	for i, item := range ol.Items {
		md.body += fmt.Sprintf("%d. ", i+1) + item + "\n"
	}
}

func (md *MD) Unordered_List(ol *MD_List) {
	for _, item := range ol.Items {
		md.body += "- " + item + "\n"
	}
}
