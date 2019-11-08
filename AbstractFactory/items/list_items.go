package items

type ListLink struct {
	*LinkImpl
}

func NewListLink(caption, url string) *ListLink {
	link := &ListLink{}
	newLink := NewLinkImpl(caption, url, link)
	link.LinkImpl = newLink
	return link
}

func (LL *ListLink) MakeHtml() string {
	return "  <li><a href=\"" + LL.url + "\">" + LL.caption + "</a></li>\n"
}

type ListTray struct {
	*TrayImpl
}

func NewListTray(caption string) *ListTray {
	tray := &ListTray{}
	newTray := NewTrayImpl(caption, tray)
	tray.TrayImpl = newTray
	return tray
}

func (LT *ListTray) MakeHtml() string {
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
	*PageImpl
}

func NewListPage(title, author string) *ListPage {
	page := &ListPage{}
	newPage := NewPageImpl(title, author, page)
	page.PageImpl = newPage
	return page
}

func (LP *ListPage) MakeHtml() string {
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
