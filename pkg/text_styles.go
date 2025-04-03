package MD_Builder

func Bold(text string) string {
	return "**" + text + "**"
}

func Italic(text string) string {
	return "*" + text + "*"
}

func Bold_Italic(text string) string {
	return "***" + text + "***"
}

func Strike_Through(text string) string {
	return "~~" + text + "~~"
}
