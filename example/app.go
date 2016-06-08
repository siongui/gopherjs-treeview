package main

import (
	"encoding/json"
	tv "github.com/siongui/gopherjs-treeview"
	"strings"
)

const treeviewData = `{"child": [
      { "text": "item 1",
        "child": [
          { "text": "item 1-1" },
          { "text": "item 1-2",
            "child": [
              { "text": "item 1-2-1" },
              { "text": "item 1-2-2" }
            ]},
          { "text": "item 1-3" }
        ]},
      { "text": "item 2",
        "child": [
          { "text": "item 2-1" },
          { "text": "item 2-2" }
        ]},
      { "text": "item 3" }
    ]}`

func main() {
	dec := json.NewDecoder(strings.NewReader(treeviewData))
	root := tv.Node{}
	dec.Decode(&root)
	tv.NewTreeview("treeview", root)
}
