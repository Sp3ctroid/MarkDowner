package MD_Builder

func New_Table() *MD_Table {
	return &MD_Table{
		Headers: []string{},
		Rows:    [][]string{},
	}
}

func (table *MD_Table) Table_Headers(a ...interface{}) {
	for _, header := range a {
		table.Headers = append(table.Headers, header.(string))
	}
}

func (table *MD_Table) Table_Rows(Rows [][]string) {
	table.Rows = append(table.Rows, Rows...)
}
