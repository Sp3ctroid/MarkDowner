package MD_Builder

func New_Table() *MD_Table {
	return &MD_Table{
		Headers: []Table_Header{},
		Rows:    [][]string{},
	}
}

func (table *MD_Table) Table_Headers(Headers []Table_Header) {
	table.Headers = append(table.Headers, Headers...)
}

func (table *MD_Table) Table_Rows(Rows [][]string) {
	table.Rows = append(table.Rows, Rows...)
}
