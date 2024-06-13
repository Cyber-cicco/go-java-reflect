package reflect

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func GetrootClass(n *sitter.Node) *Document {
    return &Document{root: n}
}
