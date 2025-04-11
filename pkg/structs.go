package MD_Builder

type MD struct {
	dest_folder string
	body        string
	err         error
}

type MD_Table struct {
	Headers []Table_Header
	Rows    [][]string
}

type MD_List struct {
	Items     []string
	level_idx []int
}

type MD_List_Item struct {
	Item     string
	Level    int
	Ordering string
}

type MD_BlockQuote struct {
	Text  []string
	Level int
}

type MD_CodeBlock struct {
	Code     string
	Language string
}

type Table_Header struct {
	Header     string
	Allignment string
}

type Image struct {
	Source string
}
