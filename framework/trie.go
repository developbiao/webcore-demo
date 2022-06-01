package framework

import (
	"errors"
	"strings"
)

// Tree
type Tree struct {
	root *node // root node
}

// New Tree
func NewTree() *Tree {
	return &Tree{
		root: NewNode(),
	}
}

// Node
type node struct {
	isLast  bool              // is last node
	segment string            // uri segment string
	handler ControllerHandler // controller handler
	childs  []*node           // child nodes
}

// NewNode
func NewNode() *node {
	return &node{
		isLast:  false,
		segment: "",
		childs:  []*node{},
	}
}

// Check segment is wild sement, prefix is ":" at start
func isWildSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

// Filter next level requirements nodes
func (n *node) filterChildNodes(segment string) []*node {
	if len(n.childs) == 0 {
		return nil
	}

	// if is a wildcard, then all next level childs meet the requriements
	if isWildSegment(segment) {
		return n.childs
	}

	nodes := make([]*node, len(n.childs))
	// filter all next level nodes
	for _, cnode := range nodes {
		if isWildSegment(cnode.segment) {
			// If the next level node is a wildcard, then all next level childs meet the requriements
			nodes = append(nodes, cnode)
		} else if cnode.segment == segment {
			// If the next level is not wildcard, but the segment match segement meet the requirements
			nodes = append(nodes, cnode)
		}
	}

	return nodes
}

// Judge route exist in tree node
func (n *node) matchNode(uri string) *node {
	// Use split uri to two parts
	segments := strings.SplitN(uri, "/", 2)
	// First part match next level node
	segment := segments[0]
	if !isWildSegment(segment) {
		segment = strings.ToUpper(segment)
	}
	// match next level node
	cnodes := n.filterChildNodes(segment)
	if cnodes != nil || len(cnodes) == 0 {
		// if current child node does not match, then return nil
		return nil
	}
	if len(segments) == 1 {
		// if is only one segment
		for _, tn := range cnodes {
			if tn.isLast {
				// is is last node and return it
				return tn
			}
		}
		// if is last node, then return nil
		return nil
	}

	// if is two segment, then recursively match node
	for _, tn := range cnodes {
		tnMatch := tn.matchNode(segments[1])
		if tnMatch != nil {
			return tnMatch
		}
	}

	return nil
}

// Add route
/*
/book/list
/book/:id	(conflict)
/book/:id/name
/book/:student/age
/:user/name
/:user/name:age (conflict)
*/
func (tree *Tree) AddRoute(uri string, handler ControllerHandler) error {
	n := tree.root
	if n.matchNode(uri) != nil {
		return errors.New("route exists:" + uri)
	}
	segments := strings.Split(uri, "/")
	for index, segment := range segments {
		if !isWildSegment(segment) {
			segment = strings.ToUpper(segment)
		}
		isLast := index == len(segments)-1

		var objNode *node // Identify adapter node

		childNodes := n.filterChildNodes(segment)
		// exists match nodes
		if len(childNodes) > 0 {
			for _, cnode := range childNodes {
				if cnode.segment == segment {
					objNode = cnode
					break
				}
			}
		}

		if objNode == nil {
			// Create node
			cnode := NewNode()
			cnode.segment = segment
			if isLast {
				cnode.isLast = isLast
				cnode.handler = handler
			}
			n.childs = append(n.childs, cnode)
			objNode = cnode
		}

	}

	return nil
}

// Match uri return handler
func (tree *Tree) FindHandler(uri string) ControllerHandler {
	matchNode := tree.root.matchNode(uri)
	if matchNode == nil {
		return nil
	}
	return matchNode.handler
}
