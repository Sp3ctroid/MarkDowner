package MD_Builder

type MD struct {
	dest_folder string
	body        string
	err         error
}

type MD_Table struct {
	Headers []string
	Rows    [][]string
}

type MD_List struct {
	Items []string
}

type MD_BlockQuote struct {
	Text  []string
	Level int
}

type MD_CodeBlock struct {
	Code     string
	Language string
}
