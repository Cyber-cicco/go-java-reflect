package reflect

import sitter "github.com/smacker/go-tree-sitter"

type Type struct {
    root *sitter.Node
    Document *Document
}

