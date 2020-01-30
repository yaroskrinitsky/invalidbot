package main

import (
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

type Matcher interface {
	Match(n *html.Node) bool
}

type ClassSelector struct {
	Class string
}

func (s ClassSelector) Match(n *html.Node) bool {
	if n.Type != html.ElementNode {
		return false
	}

	for _, a := range n.Attr {
		if a.Key == "class" && strings.Contains(a.Val, s.Class) {
			return true
		}
	}

	return false
}

type TagSelector struct {
	Tag string
}

func (s TagSelector) Match(n *html.Node) bool {
	return n.Type == html.ElementNode && s.Tag == n.Data
}

type CompositeSelector struct {
	selectors []Matcher
}

func (s CompositeSelector) Match(n *html.Node) bool {
	for _, sel := range s.selectors {
		if !sel.Match(n) {
			return false
		}
	}

	return true
}

func (s *CompositeSelector) Where(m Matcher) *CompositeSelector {
	if s.selectors == nil {
		s.selectors = make([]Matcher, 0)
	}

	s.selectors = append(s.selectors, m)

	return s
}

//Participant league player
type Participant struct {
	Pos         string
	Team        string
	Name        string
	TourPoints  string
	TotalPoints string
	TeamURL     string
	ProfileURL  string
}

//GetTable returns the league rankings table
func GetParticipants(url string) ([]Participant, error) {
	tableNode, err := GetTableNode(url)
	if err != nil {
		return nil, err
	}

	participants := ParseTable(tableNode)

	return participants, err
}

func ParseTable(table *html.Node) []Participant {
	tbody := Query(table, TagSelector{Tag: "tbody"})
	CleanUpTree(tbody)

	var participants []Participant
	for tr := tbody.FirstChild; tr != nil; tr = tr.NextSibling {
		p := ParseRow(tr)
		participants = append(participants, p)
	}
	return participants
}

func ParseRow(tr *html.Node) Participant {
	p := Participant{}

	cell := tr.FirstChild
	p.Pos = cell.FirstChild.Data

	cell = cell.NextSibling.NextSibling
	p.TeamURL = "https://www.sports.ru/" + GetAttributeValue(cell.FirstChild, "href")
	p.Team = cell.FirstChild.FirstChild.Data

	cell = cell.NextSibling
	p.ProfileURL = "https://www.sports.ru/" + GetAttributeValue(cell.FirstChild.FirstChild, "href")
	p.Name = cell.FirstChild.FirstChild.FirstChild.Data

	cell = cell.NextSibling
	p.TourPoints = cell.FirstChild.Data

	cell = cell.NextSibling
	p.TotalPoints = cell.FirstChild.Data

	return p
}

func GetAttributeValue(n *html.Node, attr string) string {
	for _, a := range n.Attr {
		if a.Key == attr {
			return a.Val
		}
	}

	return ""
}

func GetTableNode(url string) (*html.Node, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	nodeTree, err := html.Parse(r.Body)

	if err != nil {
		return nil, err
	}
	selector := CompositeSelector{}
	table := Query(nodeTree, selector.Where(TagSelector{Tag: "table"}).Where(ClassSelector{Class: "stat-table"}))

	return table, nil
}

func CleanUpTree(n *html.Node) {
	nodes := make([]*html.Node, 0)
	var f func(nn *html.Node)
	f = func(nn *html.Node) {
		for c := nn.FirstChild; c != nil; c = c.NextSibling {

			if c.Type == html.TextNode && (c.Data == "\r\n" || strings.TrimSpace(c.Data) == "") {
				nodes = append(nodes, c)
			}
			f(c)
		}
	}

	f(n)

	for _, node := range nodes {
		node.Parent.RemoveChild(node)
	}
}

func Query(n *html.Node, m Matcher) *html.Node {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if m.Match(c) {
			return c
		}
		if matched := Query(c, m); matched != nil {
			return matched
		}
	}

	return nil
}