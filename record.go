package reflect

import (
	sitter "github.com/smacker/go-tree-sitter"
)

type Record struct {
    root *sitter.Node
    document *Document
    parent TypeElement
}

func NewRecord(node *sitter.Node, d *Document, parent TypeElement) (*Record, error) {
    root, err := d.NewRootType(node)

	return &Record{
		root:     root,
		document: d,
        parent: parent,
	}, err
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
