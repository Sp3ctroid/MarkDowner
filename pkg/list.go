package MD_Builder

import (
	"bytes"
	"fmt"
	"strings"
)

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

func New_Complex_List(level int, ordering string) *Complex_List {
	return &Complex_List{
		Items:    make([]string, 0),
		Level:    level,
		Ordering: ordering,
	}
}

func (cl *Complex_List) Fill_List(a ...interface{}) {
	var b bytes.Buffer
	for i, item := range a {
		b.Reset()
		if cl.Ordering == "ordered" {
			b.WriteString(strings.Repeat("	", cl.Level-1) + fmt.Sprintf("%d. ", i+1) + item.(string))
			cl.Items = append(cl.Items, b.String())
		} else if cl.Ordering == "unordered" {
			b.WriteString(strings.Repeat("	", cl.Level-1) + "- " + item.(string))
			cl.Items = append(cl.Items, b.String())
		}
	}

}

func new_local_complex_list() *Complex_List {
	return &Complex_List{
		Items: []string{},
	}
}

func Connect_Complex_List(cl1 *Complex_List, cl2 *Complex_List, sep_level int) *Complex_List {
	final_list := new_local_complex_list()

	for i, item := range cl1.Items {
		if i == sep_level {
			final_list.Items = append(final_list.Items, cl2.Items...)
		}
		final_list.Items = append(final_list.Items, item)
	}

	return final_list
}
