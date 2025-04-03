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
func (md *MD) Table(table *MD_Table) {
	md.body += "| "
	var dashes string
	for _, header := range table.Headers {
		md.body += header + " | "
		dashes += "---|"
	}
	md.body += "\n"
	md.body += "|" + dashes + "\n"
	for _, row := range table.Rows {
		md.body += "| "
		for _, cell := range row {
			md.body += cell + " | "
		}
		md.body += "\n"
	}
}
