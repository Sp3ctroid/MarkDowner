package MD_Builder

func New_BlockQuote() *MD_BlockQuote {
	return &MD_BlockQuote{
		Text: []string{},
	}
}

func (bq *MD_BlockQuote) Fill_BlockQuote(text string, level int) {
	bq.Text = append(bq.Text, text)
	bq.Level = level

}
