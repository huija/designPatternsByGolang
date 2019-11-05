package items

type ListLink struct {
	Link
}

func NewListLink(caption, url string) ListLink {
	return ListLink{NewLink(caption, url)}
}

func (LL ListLink) MakeHtml() string {
	return "  <li><a href=\"" + LL.url + "\">" + LL.caption + "</a></li>\n"
}

type ListTray struct {
	Tray
}

func NewListTray(caption string) ListTray {
	return ListTray{NewTray(caption)}
}

func (LT ListTray) MakeHtml() string {
	result := ""
	result += "<li>\n"
	result += LT.caption + "\n"
	result += "<ul>\n"
	for _, value := range LT.tray {
		result += value.MakeHtml()
	}
	result += "</ul>\n"
	result += "</li>\n"
	return result
}

type ListPage struct {
	Page
}

func NewListPage(title, author string) ListPage {
	return ListPage{NewPage(title, author)}
}

func (LP ListPage) MakeHtml() string {
	result := ""
	result += "<html><head><title>" + LP.title + "</title></head>\n"
	result += "<body>\n"
	result += "<h1>" + LP.title + "</h1>\n"
	result += "<ul>\n"
	for _, value := range LP.content {
		result += value.MakeHtml()
	}
	result += "</ul>\n"
	result += "<hr><address>" + LP.author + "</address>"
	result += "</body></html>\n"
	return result
}
