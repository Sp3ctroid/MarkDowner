package MD_Builder

import (
	"fmt"
	"os"
	"strings"
)

func (md *MD) LineBreak() {
	md.body += "\n"
}

func NewMD() *MD {
	return &MD{
		dest_folder: "",
		body:        "",
		err:         nil,
	}
}

func (md *MD) SetDestFolder(dest_folder string) {
	md.dest_folder = dest_folder
}

func (md *MD) H_n(level int, text string) {
	if level == 0 {
		md.Plain_Text(text)
		return
	}
	for i := 0; i < level; i++ {
		md.body += "#"
	}
	md.body += " " + text + "\n"
}

func (md *MD) H_f(level int, format string, a ...interface{}) {
	md.H_n(level, fmt.Sprintf(format, a...))
}

func (md *MD) Plain_Text(text string) {
	md.body += text + "\n"
}

func (md *MD) Plain_Text_f(format string, a ...interface{}) {
	md.Plain_Text(fmt.Sprintf(format, a...))
}

func (md *MD) Save() {
	f, err := os.Create(md.dest_folder)
	if err != nil {
		md.err = err
		return
	}
	defer f.Close()
	f.WriteString(md.body)
}

func (md *MD) List(list *MD_List) {
	for _, item := range list.Items {
		md.body += item + "\n"
	}
}

func (md *MD) Block_Quote(bq *MD_BlockQuote) {
	for _, line := range bq.Text {
		md.body += strings.Repeat(">", bq.Level) + " " + line + "\n"
	}
}

func (md *MD) Table(table *MD_Table) {
	md.body += "| "
	var dashes string
	for _, header := range table.Headers {
		md.body += header.Header + " | "
		if header.Allignment == "left" || header.Allignment == "l" {
			dashes += ":---|"
		} else if header.Allignment == "right" || header.Allignment == "r" {
			dashes += "---:|"
		} else if header.Allignment == "center" || header.Allignment == "c" {
			dashes += ":---:|"
		} else {
			dashes += "---|"
		}

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

func (md *MD) CodeBlock(cb *MD_CodeBlock) {
	md.body += cb.Code
}

func (md *MD) Image(image *Image, image_name string) {
	md.body += "![" + image_name + "](" + image.Source + ")"
}

func (md *MD) Alert(name string, text string) {
	if name == "Warning" {
		md.body += "[!WARNING]" + text + "\n"
	} else if name == "Important" {
		md.body += "[!IMPORTANT]" + text + "\n"
	} else if name == "Note" {
		md.body += "[!NOTE]" + text + "\n"
	}
}
