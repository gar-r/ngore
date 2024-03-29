package activity

import (
	"github.com/gar-r/ngore/parse"
	"golang.org/x/net/html"
)

func ParseResponse(doc *html.Node) *Info {
	info := &Info{
		Rank:    Rank{},
		Stats:   Stats{},
		History: make([]TorrentActivity, 0),
	}
	parseTable(doc, info)
	parseHistory(doc, info)
	return info
}

func parseHistory(doc *html.Node, info *Info) {
	torrents := parse.GetElementsByClass(doc, "hnr_all")
	torrents = append(torrents, parse.GetElementsByClass(doc, "hnr_all2")...)
	for _, torrent := range torrents {
		item := TorrentActivity{
			Name:      parseName(torrent),
			Start:     parseDivText(torrent, "hnr_tstart"),
			Updated:   parseDivText(torrent, "hnr_tlastactive"),
			Status:    parseStatus(torrent, "hnr_tseed"),
			Up:        parseDivText(torrent, "hnr_tup"),
			Down:      parseDivText(torrent, "hnr_tdown"),
			Remaining: parseStatus(torrent, "hnr_ttimespent"),
			Ratio:     parseStatus(torrent, "hnr_tratio"),
		}
		info.History = append(info.History, item)
	}
}

func parseStatus(node *html.Node, class string) string {
	div := parse.GetElementByClass(node, class)
	if div == nil {
		return ""
	}
	span := div.FirstChild
	if span == nil {
		return ""
	}
	return parse.GetText(span)
}

func parseDivText(node *html.Node, class string) string {
	div := parse.GetElementByClass(node, class)
	if div == nil {
		return ""
	}
	return parse.GetText(div)
}

func parseName(node *html.Node) string {
	a := parse.GetElementByTag(node, "a")
	if a == nil {
		return ""
	}
	name, _ := parse.FindAttr(a, "title")
	return name
}

func parseTable(doc *html.Node, info *Info) {
	element := findTableElement(doc)
	if element == nil {
		return
	}
	data := parse.GetElementsByClass(element, "dd")
	if len(data) < 9 {
		return
	}
	info.Rank = Rank{
		Daily:     parse.GetText(data[0]),
		Weekly:    parse.GetText(data[1]),
		Monthly:   parse.GetText(data[2]),
		PrevMonth: parse.GetText(data[3]),
	}
	info.CanDownload = parse.GetText(data[4])
	info.Stats = Stats{
		Current:     parse.GetText(data[5]),
		Allowed:     parse.GetText(data[6]),
		PenMonths:   parse.GetText(data[7]),
		PenTorrents: parse.GetText(data[8]),
	}
}

func findTableElement(doc *html.Node) *html.Node {
	elements := parse.GetElementsByClass(doc, "fobox_tartalom")
	for _, element := range elements {
		n := parse.GetElementByClass(element, "dd")
		if n != nil {
			return element
		}
	}
	return nil
}
