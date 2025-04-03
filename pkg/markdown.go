package MD_Builder

import (
	"fmt"
	"os"
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
