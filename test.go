package main

import MD_Builder "github.com/Sp3ctroid/MarkDowner/pkg"

func main() {
	MD := MD_Builder.NewMD()
	MD.H_n(1, "test")
	MD.Plain_Text("Plain text test")
	Table1 := MD_Builder.New_Table()
	Table1.Table_Headers([]MD_Builder.Table_Header{{Header: "LEFT", Allignment: "l"}, {Header: "b", Allignment: "center"}, {Header: "rightish", Allignment: "r"}})
	Table1.Table_Rows([][]string{{"1", "2", "3"}, {"4", "5", "6"}})
	MD.Table(Table1)

	Test_List := MD_Builder.New_List()
	Test_List.Fill_List([]MD_Builder.MD_List_Item{
		{"Begining", 1, "ordered"},
		{Item: "a", Level: 1, Ordering: "unordered"},
		{Item: "b", Level: 2, Ordering: "unordered"},
		{Item: "c", Level: 1, Ordering: "unordered"},
		{Item: "d", Level: 2, Ordering: "unordered"},
		{"e", 1, "ordered"},
		{"f", 2, "ordered"},
		{"g", 1, "ordered"},
		{"h", 2, "ordered"},
	}, true)

	Test_List2 := MD_Builder.New_List()
	Test_List2.Fill_List([]MD_Builder.MD_List_Item{
		{"Begining", 1, "ordered"},
		{Item: "a", Level: 1, Ordering: "unordered"},
		{Item: "b", Level: 2, Ordering: "unordered"},
		{Item: "c", Level: 1, Ordering: "unordered"},
		{Item: "d", Level: 2, Ordering: "unordered"},
		{"e", 1, "ordered"},
		{"f", 2, "ordered"},
		{"g", 1, "ordered"},
		{"h", 2, "ordered"},
		{"X", 2, "ordered"},
		{"A", 1, "ordered"},
		{"B", 1, "ordered"},
	}, false)

	MD.List(Test_List)
	MD.List(Test_List2)

	BQ := MD_Builder.New_BlockQuote()
	BQ.Fill_BlockQuote("Sweet Pumpkin Pie", 1)
	MD.LineBreak()
	MD.Block_Quote(BQ)
	MD.LineBreak()
	MD.Plain_Text(MD_Builder.Strike_Through("Strike through test"))
	MD.Plain_Text(MD_Builder.Bold_Italic("Bold Italic test"))
	MD.Plain_Text(MD_Builder.Bold("Bold test"))
	MD.Plain_Text(MD_Builder.Italic("Italic test"))

	Code_block := MD_Builder.New_CodeBlock()
	Code_block.Fill_CodeBlock(`package main
import "fmt"

func main() {
	fmt.Println("Hello, World!")
}`, "go")
	MD.CodeBlock(Code_block)

	MD.Plain_Text(MD_Builder.Linkify("Google Test Link", "https://google.com"))

	MD.LineBreak()
	MD.Alert("Warning", "Warning test")
	MD.LineBreak()
	MD.Alert("Important", "Important test")
	MD.LineBreak()
	MD.Alert("Note", "Note test")
	MD.LineBreak()

	MD.SetDestFolder("test.md")
	MD.Save()
}
