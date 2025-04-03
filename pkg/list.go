package MD_Builder

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
