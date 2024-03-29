package search

import (
	"regexp"

	"github.com/gar-r/ngore/parse"
	"golang.org/x/net/html"
)

var idRegex = regexp.MustCompile(`.*id=(\d*)`)

func ParseResponse(doc *html.Node) *Result {
	return &Result{
		Torrents: parseTorrents(doc),
		Page:     parsePageInfo(doc),
	}
}

func parseTorrents(doc *html.Node) []*Torrent {
	torrents := make([]*Torrent, 0)
	nodes := parse.GetElementsByClass(doc, "box_torrent")
	for _, node := range nodes {
		t := &Torrent{}
		txt := getTxtNode(node)
		if txt != nil {
			t.Id = extractId(txt)
			t.Title = extractTitle(txt)
			t.AltTitle = extractAltTitle(txt)
		}
		t.Health = extractHealth(node)
		t.Peers = extractPeers(node)
		t.Seeds = extractSeeds(node)
		t.Size = extractSize(node)
		t.Uploaded = extractUploaded(node)
		t.Uploader = extractUploader(node)
		torrents = append(torrents, t)
	}
	return torrents
}

func extractId(n *html.Node) string {
	a := parse.GetElementByTag(n, "a")
	if a != nil {
		href := hrefAttr(a)
		matches := idRegex.FindAllStringSubmatch(href, -1)
		if len(matches) == 1 {
			return matches[0][1]
		}
	}
	return ""
}

func getTxtNode(n *html.Node) *html.Node {
	node := parse.GetElementByClass(n, "torrent_txt")
	if node == nil {
		node = parse.GetElementByClass(n, "torrent_txt2")
	}
	return node
}

func extractTitle(n *html.Node) string {
	a := parse.GetElementByTag(n, "a")
	if a != nil {
		return titleAttr(a)
	}
	return ""
}

func extractAltTitle(n *html.Node) string {
	node := parse.GetElementByClass(n, "siterank")
	if node == nil {
		return ""
	}
	span := parse.GetElementByTag(node, "span")
	if span == nil {
		return ""
	}
	return titleAttr(span)
}

func extractHealth(n *html.Node) string {
	node := parse.GetElementByClass(n, "box_d2")
	if node == nil {
		return ""
	}
	return parse.GetText(node)
}

func extractPeers(n *html.Node) string {
	node := parse.GetElementByClass(n, "box_l2")
	if node == nil {
		return ""
	}
	a := parse.GetElementByTag(node, "a")
	if a == nil {
		return ""
	}
	return parse.GetText(a)
}

func extractSeeds(n *html.Node) string {
	node := parse.GetElementByClass(n, "box_s2")
	if node == nil {
		return ""
	}
	a := parse.GetElementByTag(node, "a")
	if a == nil {
		return ""
	}
	return parse.GetText(a)
}

func extractSize(n *html.Node) string {
	node := parse.GetElementByClass(n, "box_meret2")
	if node == nil {
		return ""
	}
	return parse.GetText(node)
}

func extractUploaded(n *html.Node) string {
	node := parse.GetElementByClass(n, "box_feltoltve2")
	if node == nil {
		return ""
	}
	return parse.GetText(node)
}

func extractUploader(n *html.Node) string {
	node := parse.GetElementByClass(n, "box_feltolto2")
	if node == nil {
		return ""
	}
	spans := parse.GetElementsByClass(node, "feltolto_szin")
	if len(spans) == 0 {
		return ""
	}
	return parse.GetText(spans[0])
}

func titleAttr(element *html.Node) string {
	title, ok := parse.FindAttr(element, "title")
	if ok {
		return title
	}
	return ""
}

func hrefAttr(element *html.Node) string {
	href, ok := parse.FindAttr(element, "href")
	if ok {
		return href
	}
	return ""
}
