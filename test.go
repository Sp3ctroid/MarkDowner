package main

import MD_Builder "github.com/Sp3ctroid/MarkDowner/pkg"

func main() {
	MD := MD_Builder.NewMD()
	MD.H_n(1, "test")
	MD.Plain_Text("Plain text test")
	Table1 := MD_Builder.New_Table()
	Table1.Table_Headers("a", "b", "c")
	Table1.Table_Rows([][]string{{"1", "2", "3"}, {"4", "5", "6"}})
	MD.Table(Table1)
	Unord_list := MD_Builder.New_List()
	Unord_list.Fill_List("a", "b", "c")
	MD.Unordered_List(Unord_list)
	Ord_list := MD_Builder.New_List()
	Ord_list.Fill_List(MD_Builder.Bold_Italic("a"), MD_Builder.Bold("b"), MD_Builder.Italic("c"))
	MD.Ordered_List(Ord_list)
	BQ := MD_Builder.New_BlockQuote()
	BQ.Fill_BlockQuote("Sweet Pumpkin Pie")
	MD.LineBreak()
	MD.Block_Quote(BQ)
	MD.LineBreak()
	MD.Plain_Text(MD_Builder.Strike_Through("Strike through test"))
	MD.SetDestFolder("test.md")
	MD.Save()
}
