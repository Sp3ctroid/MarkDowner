package MD_Builder

func New_BlockQuote() *MD_BlockQuote {
	return &MD_BlockQuote{
		Text: []string{},
	}
}

func (bq *MD_BlockQuote) Fill_BlockQuote(text string) {
	bq.Text = append(bq.Text, text)
}
func (md *MD) Block_Quote(bq *MD_BlockQuote) {
	for _, line := range bq.Text {
		md.body += "> " + line + "\n"
	}
}
