package reflect

import sitter "github.com/smacker/go-tree-sitter"

type Literal struct {
    root *sitter.Node
    document *Document
}

func (l *Literal) GetValue() *sitter.Node {
    switch l.root.Type() {
    case "string_literal":
        return l.root
    }
    return l.root
}
