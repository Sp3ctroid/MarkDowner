package MD_Builder

func New_CodeBlock() *MD_CodeBlock {
	return &MD_CodeBlock{}
}

func (cb *MD_CodeBlock) Fill_CodeBlock(code string, language string) {
	cb.Language = language
	cb.Code = "```" + cb.Language + "\n" + code + "\n```\n"
}
