package builder

import (
	"io"
	"os"
)

type HTMLBuilder struct {
	fileName string
	file     *os.File
}

func (H *HTMLBuilder) MakeTitle(title string) {
	H.fileName = title + ".html"
	var err error
	if H.file, err = os.OpenFile(H.fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644); err != nil {
		_, _ = io.WriteString(H.file, "<html><head><title>"+title+"</title></head><body>\n")
		_, _ = io.WriteString(H.file, "<h1>"+title+"</h1>\n")
	}
}
func (H *HTMLBuilder) MakeString(str string) {
	_, _ = io.WriteString(H.file, "<p>"+str+"</p>\n")
}
func (H *HTMLBuilder) MakeItems(items []string) {
	_, _ = io.WriteString(H.file, "<ul>\n")
	for i := 0; i < len(items); i++ {
		_, _ = io.WriteString(H.file, "<li>"+items[i]+"</li>\n")
	}
	_, _ = io.WriteString(H.file, "</ul>\n")
}
func (H *HTMLBuilder) Close() {
	_, _ = io.WriteString(H.file, "</body></html>\n")
	H.file.Close()
}

func (H *HTMLBuilder) GetFileName() string {
	return H.fileName
}
