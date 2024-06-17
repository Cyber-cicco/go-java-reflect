package reflect

import sitter "github.com/smacker/go-tree-sitter"

type Scope struct {
    root *sitter.Node
    document *Document
}

func NewScope(node *sitter.Node, document *Document) (*Scope, error) {
    return &Scope{
        root: node,
        document: document,
    }, nil
}

func (s *Scope) ToString() string {
    return s.root.Content(s.document.content)
}


