package gographs

import (
	"strings"
)

const (
	CEIL_STRING  = "-"
	FLOOR_STRING = "-"
	WALL_STRING  = "|"
	RESET        = "\033[0m"
	RED          = "\033[31m"
)

type Node struct {
	Text       string
	Separator  string
	Active     bool
	IsLeaf     bool
	ChildNodes []*Node
}

func (n *Node) GetOrCreateNode(text string, active bool) *Node {
	for _, node := range n.ChildNodes {
		if node.Text == text {
			if active {
				node.Active = true
			}
			return node
		}
	}

	node := CreateNode(text, n.Separator, active)
	n.ChildNodes = append(n.ChildNodes, &node)
	return &node
}
func strOf(value string, counter int) string {
	if counter < 0 {
		return ""
	}
	return strings.Repeat(value, counter)
}

func (n *Node) formatString(leftSpace int) string {
	if n.Text == "" {
		return ""
	}
	leftSpaceString := strOf(" ", leftSpace-2) + "|_ "
	box := leftSpaceString + n.Text
	if len(n.ChildNodes) > 0 {
		box += "/\n"
	}
	if n.Active {
		box = RED + box + RESET
	}
	return box
}

func (n *Node) ToString(leftSpace int, onlyActive bool) string {
	nextSpace := 0
	if n.Text != "" {
		nextSpace = 4 + leftSpace
	}

	childs := []string{}
	for _, node := range n.ChildNodes {
		if !onlyActive || (onlyActive && node.Active) {
			childs = append(childs, node.ToString(nextSpace, onlyActive))
		}
	}
	box := n.formatString(leftSpace)
	if len(childs) > 0 {
		return box + strings.Join(childs[:], "\n")
	}

	return box
}

func (n *Node) Insert(text string, active bool) bool {

	nodes := strings.SplitN(text, "/", 2)
	node := n.GetOrCreateNode(nodes[0], active)
	if len(nodes) > 1 {
		node.Insert(nodes[1], active)
	} else {
		node.IsLeaf = true
	}
	return true
}

func CreateNode(text, separator string, active bool) Node {
	return Node{
		Text:       text,
		Separator:  separator,
		Active:     active,
		ChildNodes: []*Node{},
		IsLeaf:     false,
	}
}
