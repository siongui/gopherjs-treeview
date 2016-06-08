package treeview

import (
	"github.com/gopherjs/gopherjs/js"
)

type Node struct {
	Text  string
	Child []*Node
}

func traverseTreeviewData(node *Node) *js.Object {
	doc := js.Global.Get("document")

	if node.Child != nil {
		div := doc.Call("createElement", "div")
		div.Get("classList").Call("add", "item")

		sign := doc.Call("createElement", "span")
		sign.Set("innerHTML", "+")

		span := doc.Call("createElement", "span")
		span.Get("classList").Call("add", "treeNode")
		span.Set("innerHTML", node.Text)

		div.Call("appendChild", sign)
		div.Call("appendChild", span)

		childrenContainer := doc.Call("createElement", "div")
		childrenContainer.Get("classList").Call("add", "childrenContainer")
		for _, child := range node.Child {
			childrenContainer.Call("appendChild", traverseTreeviewData(child))
		}
		childrenContainer.Get("style").Set("display", "none")

		span.Set("onclick", func(event *js.Object) {
			if childrenContainer.Get("style").Get("display").String() == "none" {
				childrenContainer.Get("style").Set("display", "")
				sign.Set("innerHTML", "-")
			} else {
				childrenContainer.Get("style").Set("display", "none")
				sign.Set("innerHTML", "+")
			}
		})

		all := doc.Call("createElement", "div")
		all.Call("appendChild", div)
		all.Call("appendChild", childrenContainer)

		return all
	} else {
		div := doc.Call("createElement", "div")
		div.Get("classList").Call("add", "item")
		div.Set("innerHTML", "<span class='treeNode'>"+node.Text+"</span>")
		return div
	}
}

func appendCSSToHeadElement() {
	css := `.item {
	  margin-bottom: 3px;
	  padding-bottom: 3px;
	  border-bottom: 1px solid #E0E0E0;
	}

	.item:hover {
	  background-color: #F0F8FF;
	}

	.treeNode:hover {
	  cursor: pointer;
	  color: blue;
	}

	.childrenContainer {
	  margin-left: .4em;
	  padding-left: .4em;
	  border-left: 1px dotted blue;
	}`
	s := js.Global.Get("document").Call("createElement", "style")
	s.Set("innerHTML", css)
	// insert style of treeview at the end of head element
	js.Global.Get("document").Call("getElementsByTagName", "head").Call("item", 0).Call("appendChild", s)
}

func NewTreeview(id string, root Node) {
	appendCSSToHeadElement()
	doc := js.Global.Get("document")
	treeviewContainer := doc.Call("getElementById", id)

	for _, child := range root.Child {
		node := traverseTreeviewData(child)
		treeviewContainer.Call("appendChild", node)
	}
}
