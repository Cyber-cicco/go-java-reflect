package reflect

import (
	"errors"

	sitter "github.com/smacker/go-tree-sitter"
)

type Record struct {
    root *sitter.Node
    document *Document
    parent TypeElement
}

func NewRecord(node *sitter.Node, d *Document, parent TypeElement) (*Record, error) {

    if node.Type() != "record_declaration" {
        return nil, errors.New("Unexpected node for record declaration : " + node.Type())
    }

	return &Record{
		root:     node,
		document: d,
        parent: parent,
	}, nil
}

func (r *Record) GetDocument() *Document {
    return r.document
}

func (r *Record) GetDeclaredName() string {
    return r.root.ChildByFieldName("name").Content(r.document.content)
}

func (r *Record) GetParent() TypeElement {
    return r.parent
}
